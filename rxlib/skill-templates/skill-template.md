# Skill Template

<!--
  USAGE
  -----
  1. Copy this file into the appropriate domain's skills/ folder.
  2. Rename using kebab-case describing the CAPABILITY, e.g. "chain-of-thought.md".
  3. Replace every {{PLACEHOLDER}} with real content.
  4. Append the skill block verbatim to the end of any system instruction to enable this capability.
-->

---

## Metadata

| Field | Value |
|---|---|
| **Domain** | {{DOMAIN}} |
| **Skill Name** | {{SKILL_NAME}} |
| **Compatible Roles** | {{COMPATIBLE_ROLES}} (e.g. "all", "Software Engineer", "Analyst") |
| **Gemini Feature Used** | {{GEMINI_FEATURE}} (e.g. grounding, function calling, code execution) |

---

## Skill Description

{{ONE_SENTENCE_SUMMARY_OF_WHAT_THIS_SKILL_DOES}}

## When to Activate

Activate this skill when the user:

* {{TRIGGER_CONDITION_1}}
* {{TRIGGER_CONDITION_2}}
* {{TRIGGER_CONDITION_3}}

## Behaviour

When this skill is active:

1. {{BEHAVIOUR_STEP_1}}
2. {{BEHAVIOUR_STEP_2}}
3. {{BEHAVIOUR_STEP_3}}

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: {{SKILL_NAME}}

{{SKILL_SYSTEM_INSTRUCTION_BLOCK}}
```

## Notes and Limitations

* {{NOTE_1}}
* {{NOTE_2}}
