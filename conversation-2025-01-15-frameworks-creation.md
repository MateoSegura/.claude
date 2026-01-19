# Conversation Summary: Frameworks Folder Creation

**Date:** 2025-01-15
**Branch:** product/financial-advisor

---

## What Was Done

Created the `/docs/frameworks/` folder - the architecture & design phase that bridges product validation and legal protection.

### Files Created (13 total)

**Templates (5):**
- `ARCHITECTURE_BRIEF.MD` - Quick architecture capture
- `COMPONENT_INVENTORY.MD` - List all components with IDs
- `CONTRACT_DEFINITIONS.MD` - Define all 6 contract types
- `DEPENDENCY_MATRIX.MD` - Map what blocks what
- `SKILLS_BOM.MD` - Skills bill of materials

**Guides (3):**
- `CONTRACT_TAXONOMY.MD` - 6 contract types explained (Domain, Persistence, API, Message, Service, Infrastructure)
- `PRODUCT_TYPE_FRAMEWORKS.MD` - Custom framework for each product path (A-H)
- `AI_COORDINATION.MD` - How to use Claude Code in parallel development

**Checklists (5):**
- `01_ARCHITECTURE_VISION.MD` - Phase 0 prep
- `02_CONTRACT_DEFINITION.MD` - Phase 0 execution (freeze contracts)
- `03_FOUNDATION_SETUP.MD` - Phase 1 (CI/CD, shared libs)
- `04_IMPLEMENTATION.MD` - Phase 2 (parallel workstreams)
- `05_INTEGRATION.MD` - Phase 3-4 (integration & validation)

### Files Updated

- `docs/README.md` - Added frameworks to main documentation map
- `docs/product/README.MD` - Updated integration diagram
- `docs/product/checklists/05_GO_NO_GO_DECISION.MD` - Points to frameworks after GO
- `docs/product/guides/VALIDATION_PLAYBOOK.MD` - **STILL NEEDS UPDATE** (identified in analysis)

---

## Key Concepts Implemented

1. **Contract-First Parallel Execution** - Define interfaces before implementation
2. **6 Contract Types** - Domain, Persistence, API, Message, Service, Infrastructure
3. **5 Development Phases** - 0 (Contracts) → 1 (Foundation) → 2 (Implementation) → 3 (Integration) → 4 (Validation)
4. **Product Type Mapping** - Each path (A-H) has specific contract requirements
5. **AI Coordination Protocol** - How to run parallel Claude Code sessions safely

---

## New Workflow

```
/product/ (validate) → GO → /frameworks/ (architect) → CONTRACTS → /legal/ (protect)
                                    ↓
                              BUILD (parallel)
```

---

## Outstanding Issues Identified (Not Yet Fixed)

### Issue 1: VALIDATION_PLAYBOOK.MD outdated
Still points directly to legal after GO decision.

### Issue 2: Legal Schedule Attachments ↔ Framework Contracts Gap
Legal schedules expect specific attachments (A-2: Architecture Doc, A-3: API Doc, etc.)
Framework contracts are in different format. No explicit mapping.

### Issue 3: Hardware Paths (E-H) Missing Specific Contract Templates
SCHEDULE_C has very specific requirements (BOM structure, PCB specs, EDA tools)
Framework contracts are generic.

### Issue 4: Business Model → Architecture Disconnected
BUSINESS_MODEL_CANVAS determines revenue model, but doesn't inform which contracts are needed.

### Issue 5: Royalty Selection Timing Problem
Product Phase 4 selects royalty template BEFORE frameworks defines architecture.
But royalty depends on architecture decisions.

### Issue 6: Firmware Paths Missing Specific Contracts
HAL interfaces, memory maps, protocol specs not captured in generic contract taxonomy.

### Issue 7: Cloud Infrastructure Gap
SCHEDULE_D exists but frameworks doesn't have cloud-specific templates.

### Issue 8: No Explicit Mapping Table
No document showing product output → framework input → legal attachment flow.

---

## Questions for User (Pending Answers)

1. Should frameworks produce documents that become legal attachments directly?
2. Should I create hardware-specific contract templates (for Paths E-H)?
3. Should architecture explicitly check business model to determine required contracts?
4. Should royalty selection move to frameworks, or stay in product?
5. Add firmware-specific contract types, or treat as service contract subtypes?
6. Create cloud-specific architecture templates, or keep generic?
7. Create an explicit document showing product → framework → legal flow?

---

## Context: Claude's Methodology That Informed This

The user provided Claude's output describing:
- Contract-First Parallel Execution methodology
- Phase 0-4 development structure
- Contract taxonomy (6 types)
- Bill of Materials framework
- Dependency mapping approach
- AI coordination for parallel development

This was synthesized with the existing product types (A-H) defined in `/docs/product/` and `/docs/legal/`.

---

## To Resume This Work

1. Answer the outstanding questions above
2. Fix VALIDATION_PLAYBOOK.MD
3. Potentially create hardware-specific contract templates
4. Create explicit mapping between frameworks outputs → legal attachments
5. Consider moving royalty selection timing
