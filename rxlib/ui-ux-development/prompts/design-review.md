# Design Review Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | UI/UX Development |
| **Task** | Review a UI design for usability, accessibility, and visual hierarchy |
| **Recommended Model** | gemini-2.5-pro (text) or gemini-2.0-flash (image input) |
| **Recommended System Instruction** | `ui-ux-development/system-instructions/ux-designer.md` |
| **Output Format** | Structured Markdown |

---

## Prompt

Please review the following UI design for usability, accessibility, and visual quality.

### Design Input

Select one:
- **Image:** [Attach screenshot, mockup, or Figma export]
- **Description:** Describe the interface below

```
{{DESIGN_DESCRIPTION_OR_LEAVE_BLANK_IF_IMAGE_PROVIDED}}
```

### Context

* **Product type:** {{PRODUCT_TYPE}} (e.g. web app, mobile app, dashboard, e-commerce)
* **Target users:** {{TARGET_USERS}} (e.g. "enterprise finance professionals", "first-time smartphone users")
* **Task being reviewed:** {{TASK}} (e.g. "user completing a checkout flow", "user configuring account settings")
* **Review focus:** {{FOCUS}} (all / usability only / accessibility only / visual design only)
* **Design stage:** {{STAGE}} (wireframe / low-fidelity / high-fidelity / live product)

### Review Task

1. **Usability** — apply Nielsen's 10 Usability Heuristics to identify violations. Score each finding by severity (1–4).
2. **Accessibility** — flag WCAG 2.2 AA violations or concerns (colour contrast, missing labels, focus order, target size).
3. **Visual hierarchy** — assess whether the layout directs the user's eye to the most important elements first.
4. **Copy/microcopy** — evaluate the clarity and helpfulness of any visible text.
5. **Positive observations** — note 2–3 things that work well.
6. **Priority recommendations** — list the top 3 changes that would have the greatest impact on the user experience.

### Output Format

```
## Design Review: [Page/Component Name]

### Summary
...

### Usability Findings
| # | Heuristic | Location | Severity | Description | Recommendation |
|---|---|---|---|---|---|

### Accessibility Concerns
| # | WCAG Criterion | Element | Issue | Fix |
|---|---|---|---|---|

### Visual Hierarchy Assessment
...

### Microcopy Observations
...

### Positive Observations
- ...

### Top 3 Priority Recommendations
1. ...
```
