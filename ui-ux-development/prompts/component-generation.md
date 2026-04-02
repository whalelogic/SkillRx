# Component Generation Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | UI/UX Development |
| **Task** | Generate a UI component with accessibility built in |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `ui-ux-development/system-instructions/ui-developer.md` |
| **Output Format** | Code (JSX/TSX, HTML, Vue SFC) |

---

## Prompt

Please generate a {{COMPONENT_TYPE}} component with full accessibility support.

### Component Specification

* **Component name:** {{COMPONENT_NAME}}
* **Framework:** {{FRAMEWORK}} (React / Vue 3 / Vanilla HTML+CSS / Web Component)
* **Language:** {{LANGUAGE}} (TypeScript / JavaScript)
* **Styling approach:** {{STYLING}} (Tailwind CSS / CSS Modules / styled-components / plain CSS / existing design system — specify)
* **Component type:** {{COMPONENT_TYPE}} (e.g. modal dialog, dropdown menu, date picker, data table, accordion, combobox, toast notification, form with validation)

### Functional Requirements

{{FUNCTIONAL_REQUIREMENTS}}

(Describe the component's behaviour, states, and interactions in detail.)

### Design Tokens / Visual Spec

```
{{DESIGN_TOKENS_OR_SPEC}}
```

(Paste design tokens, Figma specs, or a description of the desired visual appearance.)

### Accessibility Requirements

* **WCAG target:** WCAG 2.2 AA (default) / AAA
* **ARIA pattern to implement:** {{ARIA_PATTERN}} (see [ARIA Authoring Practices](https://www.w3.org/WAI/ARIA/apg/) — e.g. "dialog", "combobox", "menu button")
* **Keyboard navigation:** describe expected keyboard behaviour or leave as "follow ARIA APG pattern for {{COMPONENT_TYPE}}"

### Task

1. Implement the component following the specified framework and styling approach.
2. Include all interactive states: default, hover, focus, active, disabled, error (if applicable).
3. Implement full ARIA attributes and keyboard navigation per the ARIA APG pattern.
4. Export a usage example showing the component in context.
5. Add accessibility notes explaining the ARIA pattern implemented.
6. Add a brief Storybook story skeleton (optional — specify if needed).

### Output Requirements

* Complete, runnable component (not a fragment).
* No placeholder "TODO" comments for accessibility — implement it fully.
* TypeScript types/interfaces for all props.
* Separate the component file from the usage example.
