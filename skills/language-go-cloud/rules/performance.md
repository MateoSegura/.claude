# Performance Rules (GO-PRF-*)

Performance optimization should be data-driven. These rules cover common patterns that are always beneficial.

## Memory Management

- Preallocate slices when size is known
- Reuse buffers with `sync.Pool` for hot paths
- Be mindful of string/byte conversions

## Optimization Policy

1. Write clear, correct code first
2. Measure with benchmarks and profiling
3. Optimize only proven bottlenecks
4. Document why optimizations exist

---

## GO-PRF-001: Preallocate Slices When Size is Known :green_circle:

**Tier**: Recommended

**Rationale**: Appending to slices may cause reallocation. Preallocating avoids unnecessary allocations and copies.

```go
// Correct - Preallocate slice
func ProcessUsers(users []User) []Result {
    results := make([]Result, 0, len(users))  // Preallocate capacity
    for _, u := range users {
        results = append(results, process(u))
    }
    return results
}

// Correct - Direct indexing when length is known
func ProcessUsers(users []User) []Result {
    results := make([]Result, len(users))  // Preallocate with length
    for i, u := range users {
        results[i] = process(u)
    }
    return results
}

// Incorrect - Growing slice dynamically
func ProcessUsers(users []User) []Result {
    var results []Result  // Will reallocate multiple times
    for _, u := range users {
        results = append(results, process(u))
    }
    return results
}
```

---

## GO-PRF-002: Use strings.Builder for String Concatenation :green_circle:

**Tier**: Recommended

**Rationale**: String concatenation with `+` creates a new string each time. `strings.Builder` minimizes allocations.

```go
// Correct - strings.Builder for multiple concatenations
func BuildQuery(fields []string) string {
    var b strings.Builder
    b.WriteString("SELECT ")
    for i, f := range fields {
        if i > 0 {
            b.WriteString(", ")
        }
        b.WriteString(f)
    }
    b.WriteString(" FROM users")
    return b.String()
}

// Incorrect - String concatenation in loop
func BuildQuery(fields []string) string {
    query := "SELECT "
    for i, f := range fields {
        if i > 0 {
            query += ", "
        }
        query += f  // New allocation each iteration
    }
    query += " FROM users"
    return query
}
```
