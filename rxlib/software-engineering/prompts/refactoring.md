# Refactoring Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Software Engineering |
| **Task** | Code refactoring |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `software-engineering/system-instructions/senior-software-engineer.md` |
| **Output Format** | Markdown with code blocks |

---

## Prompt

Please refactor the following `{{LANGUAGE}}` code. The goal is to improve {{REFACTORING_GOAL}} without changing the observable behaviour.

### Refactoring Goal

Select one or more:
- [ ] **Readability** — clearer names, smaller functions, better structure
- [ ] **Maintainability** — reduce coupling, increase cohesion, apply SOLID principles
- [ ] **Performance** — reduce time/space complexity, eliminate redundant work
- [ ] **Testability** — extract dependencies, reduce side effects, add seams
- [ ] **Idiomatic style** — align with `{{LANGUAGE}}` community conventions and best practices

### Constraints

* **Preserve public API:** {{PRESERVE_API}} (yes / no — if yes, do not rename public methods or change signatures)
* **Target framework/version:** {{FRAMEWORK_VERSION}}
* **Additional constraint:** {{ADDITIONAL_CONSTRAINT}} (e.g. "must remain compatible with Python 3.9")

### Original Code

```{{LANGUAGE}}
{{CODE}}
```

### Task

1. List the code smells or issues you identified before refactoring.
2. Provide the refactored code.
3. Summarise the changes made and explain how each addresses the identified issues.
4. If the refactoring changes any behaviour (even intentionally), clearly document it.

### Output Requirements

* "Issues Identified" section with a numbered list.
* Refactored code in a fenced block.
* "Summary of Changes" section mapping each issue to its resolution.
* "Behaviour Changes" section (can state "None" if there are none).
