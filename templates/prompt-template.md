# Prompt Template

<!--
  USAGE
  -----
  1. Copy this file into the appropriate domain's prompts/ folder.
  2. Rename using kebab-case describing the TASK, e.g. "code-review.md".
  3. Replace every {{PLACEHOLDER}} with real content or leave as a template variable
     that callers fill in at runtime.
  4. Paste the completed prompt as the user-turn message when calling Gemini.
-->

---

## Metadata

| Field | Value |
|---|---|
| **Domain** | {{DOMAIN}} |
| **Task** | {{TASK_NAME}} |
| **Gemini Model** | {{RECOMMENDED_MODEL}} (e.g. gemini-2.0-flash, gemini-2.5-pro) |
| **Recommended System Instruction** | {{SYSTEM_INSTRUCTION_FILE}} |
| **Output Format** | {{OUTPUT_FORMAT}} (e.g. Markdown, JSON, plain text) |

---

## Prompt

{{CONTEXT_SETTING_SENTENCE}}

### Input

{{INPUT_DESCRIPTION}}

```
{{INPUT_PLACEHOLDER}}
```

### Task

1. {{STEP_1}}
2. {{STEP_2}}
3. {{STEP_3}}

### Output Requirements

* {{OUTPUT_REQUIREMENT_1}}
* {{OUTPUT_REQUIREMENT_2}}
* {{OUTPUT_REQUIREMENT_3}}

### Constraints

* {{CONSTRAINT_1}}
* {{CONSTRAINT_2}}

---

## Example Usage

**Input:**

```
{{EXAMPLE_INPUT}}
```

**Expected Output:**

```
{{EXAMPLE_OUTPUT}}
```
