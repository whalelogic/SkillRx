# Accessibility Audit Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | UI/UX Development |
| **Task** | Audit a page or component against WCAG 2.2 AA |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `ui-ux-development/system-instructions/ui-developer.md` |
| **Output Format** | Structured Markdown audit report |

---

## Prompt

Please perform an accessibility audit of the following {{AUDIT_TARGET}} against WCAG 2.2 Level AA.

### Audit Target

Select one:
- **HTML/JSX code:** paste below
- **Screenshot/image:** attach the image
- **URL:** {{URL}} *(note: model cannot visit URLs — provide HTML source or screenshot)*

```
{{CODE_OR_DESCRIPTION}}
```

### Context

* **Target users / assistive technologies:** {{AT_USERS}} (e.g. screen reader users (NVDA, JAWS, VoiceOver), keyboard-only users, users with low vision)
* **Page/component type:** {{PAGE_TYPE}} (e.g. login form, data table, navigation menu, modal dialog)
* **Framework:** {{FRAMEWORK}}

### Audit Scope

| WCAG Principle | Include? |
|---|---|
| Perceivable (images, video, colour contrast) | {{YES_NO}} |
| Operable (keyboard, focus, timing, seizures) | {{YES_NO}} |
| Understandable (language, labels, error messages) | {{YES_NO}} |
| Robust (parsing, name/role/value, status messages) | {{YES_NO}} |

### Audit Task

1. Test the content against each applicable WCAG 2.2 AA success criterion.
2. For each failure, provide: the criterion violated, the element/location, a description of the failure, and a code-level fix.
3. Note any issues that are best practice but not strict WCAG failures.
4. Produce an overall accessibility score estimate (approximate % of criteria met).
5. Provide a prioritised remediation list ordered by user impact.

### Output Format

```
## Accessibility Audit Report

**Target:** [Page/Component Name]
**Standard:** WCAG 2.2 Level AA
**Audit Date:** [DATE]

### Failures

| # | WCAG Criterion | Level | Element/Location | Issue | Fix |
|---|---|---|---|---|---|

### Warnings (Best Practice)
| # | Observation | Recommendation |
|---|---|---|

### Passed Criteria (summary)
...

### Overall Estimate
Approximate pass rate: X% of applicable success criteria

### Prioritised Remediation List
1. [Highest impact fix]
```
