# Design System Planning Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | UI/UX Development |
| **Task** | Plan or extend a design system |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `ui-ux-development/system-instructions/ui-developer.md` |
| **Output Format** | Structured Markdown plan |

---

## Prompt

Please help me {{PLAN_OR_EXTEND}} a design system for {{PRODUCT_OR_BRAND}}.

### Context

* **Organisation size:** {{ORG_SIZE}} (startup / scale-up / enterprise)
* **Number of products/surfaces:** {{PRODUCTS}} (e.g. "1 web app", "3 web apps + 2 mobile apps")
* **Team structure:** {{TEAM}} (e.g. "1 designer, 2 frontend engineers", "5 designers, 20+ engineers, dedicated DS team")
* **Current state:** {{CURRENT_STATE}} (greenfield / existing component library / Figma library only / mature DS)
* **Tech stack:** {{TECH_STACK}} (e.g. React + TypeScript + Tailwind, Vue 3 + CSS Modules)
* **Design tool:** {{DESIGN_TOOL}} (Figma / Sketch / Adobe XD)

### Goals

{{GOALS}}

(e.g. "Reduce inconsistency across products", "Speed up component development", "Improve accessibility compliance", "Enable designers and engineers to work in sync")

### Planning Task

1. **Foundations** — recommend the token architecture: spacing scale, type scale, colour system (semantic colour tokens), border radius, shadow levels, motion/animation tokens.
2. **Component inventory** — list the 20 highest-priority components to build first, grouped by category (layout, navigation, forms, feedback, data display).
3. **Component API principles** — what naming and props conventions should the team adopt?
4. **Documentation** — what documentation is needed (Storybook, Figma annotations, usage guidelines, do/don't examples)?
5. **Contribution model** — how should teams propose, review, and contribute new components?
6. **Versioning and release** — recommend a versioning strategy and deprecation policy.
7. **Adoption roadmap** — provide a phased 3-month adoption plan.

### Output Format

Produce a structured Design System Planning Document covering all sections above. Include:
- A token naming convention example for colours (semantic + primitive layers)
- A component priority matrix table
- A 3-month roadmap table: Month | Milestone | Owner | Success Metric
