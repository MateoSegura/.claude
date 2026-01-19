# MCP Server Template

MCP (Model Context Protocol) servers provide external tool integrations.

## When to Use MCP

Use MCP when:
- Claude needs to access external APIs
- Integration requires authentication
- Operations involve external state
- You want to expose custom tools

## MCP Server Structure

```
my-mcp-server/
├── package.json          # Node.js dependencies
├── tsconfig.json         # TypeScript config
├── src/
│   ├── index.ts          # Entry point
│   └── tools/            # Tool implementations
│       ├── fetch.ts
│       └── mutate.ts
└── README.md
```

## Basic MCP Server Template (TypeScript)

```typescript
// src/index.ts
import { Server } from "@modelcontextprotocol/sdk/server/index.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import {
  CallToolRequestSchema,
  ListToolsRequestSchema,
} from "@modelcontextprotocol/sdk/types.js";

const server = new Server(
  {
    name: "my-mcp-server",
    version: "1.0.0",
  },
  {
    capabilities: {
      tools: {},
    },
  }
);

// Define available tools
server.setRequestHandler(ListToolsRequestSchema, async () => {
  return {
    tools: [
      {
        name: "my_tool",
        description: "Description of what this tool does",
        inputSchema: {
          type: "object",
          properties: {
            param1: {
              type: "string",
              description: "Description of param1",
            },
            param2: {
              type: "number",
              description: "Description of param2",
            },
          },
          required: ["param1"],
        },
      },
    ],
  };
});

// Handle tool calls
server.setRequestHandler(CallToolRequestSchema, async (request) => {
  const { name, arguments: args } = request.params;

  switch (name) {
    case "my_tool": {
      const { param1, param2 } = args as { param1: string; param2?: number };

      // Implement your tool logic here
      const result = await doSomething(param1, param2);

      return {
        content: [
          {
            type: "text",
            text: JSON.stringify(result, null, 2),
          },
        ],
      };
    }

    default:
      throw new Error(`Unknown tool: ${name}`);
  }
});

// Start server
async function main() {
  const transport = new StdioServerTransport();
  await server.connect(transport);
}

main().catch(console.error);
```

## package.json Template

```json
{
  "name": "my-mcp-server",
  "version": "1.0.0",
  "type": "module",
  "main": "dist/index.js",
  "scripts": {
    "build": "tsc",
    "start": "node dist/index.js"
  },
  "dependencies": {
    "@modelcontextprotocol/sdk": "^0.5.0"
  },
  "devDependencies": {
    "@types/node": "^20.0.0",
    "typescript": "^5.0.0"
  }
}
```

## Claude Desktop Configuration

Add to `~/.config/claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "my-server": {
      "command": "node",
      "args": ["/path/to/my-mcp-server/dist/index.js"],
      "env": {
        "API_KEY": "your-api-key"
      }
    }
  }
}
```

## Example MCP Servers

### GitHub API Server

```typescript
// Tools: list_repos, get_issues, create_issue, etc.
server.setRequestHandler(ListToolsRequestSchema, async () => ({
  tools: [
    {
      name: "github_list_repos",
      description: "List repositories for a user or organization",
      inputSchema: {
        type: "object",
        properties: {
          owner: { type: "string", description: "Username or org name" },
          type: { type: "string", enum: ["all", "public", "private"] },
        },
        required: ["owner"],
      },
    },
    {
      name: "github_get_issues",
      description: "Get issues for a repository",
      inputSchema: {
        type: "object",
        properties: {
          owner: { type: "string" },
          repo: { type: "string" },
          state: { type: "string", enum: ["open", "closed", "all"] },
        },
        required: ["owner", "repo"],
      },
    },
  ],
}));
```

### Database Query Server

```typescript
// Tools: query, insert, update (with careful permissions)
server.setRequestHandler(ListToolsRequestSchema, async () => ({
  tools: [
    {
      name: "db_query",
      description: "Execute a read-only SQL query",
      inputSchema: {
        type: "object",
        properties: {
          query: {
            type: "string",
            description: "SQL SELECT query (read-only)"
          },
        },
        required: ["query"],
      },
    },
  ],
}));

// Validate query is read-only
server.setRequestHandler(CallToolRequestSchema, async (request) => {
  const { query } = request.params.arguments as { query: string };

  // Security: Only allow SELECT queries
  if (!query.trim().toUpperCase().startsWith("SELECT")) {
    throw new Error("Only SELECT queries are allowed");
  }

  const result = await db.query(query);
  return { content: [{ type: "text", text: JSON.stringify(result) }] };
});
```

## Security Best Practices

1. **Validate all inputs** - Never trust tool arguments
2. **Minimize permissions** - Only expose necessary operations
3. **Use environment variables** - Never hardcode secrets
4. **Rate limit** - Prevent abuse
5. **Log operations** - Audit trail for debugging
6. **Handle errors gracefully** - Don't leak sensitive info

## Tool Design Guidelines

| Do | Don't |
|-----|-------|
| Use descriptive names | Use generic names like "do_thing" |
| Validate input types | Trust input blindly |
| Return structured data | Return unformatted strings |
| Document all parameters | Leave descriptions empty |
| Handle errors with context | Throw generic errors |
| Make read operations safe | Allow destructive operations without confirmation |

## Testing MCP Servers

```bash
# Build the server
npm run build

# Test manually with stdio
echo '{"jsonrpc":"2.0","id":1,"method":"tools/list"}' | node dist/index.js

# Test a tool call
echo '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"my_tool","arguments":{"param1":"test"}}}' | node dist/index.js
```
