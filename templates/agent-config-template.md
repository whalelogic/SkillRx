# Agent Configuration Template

<!--
  USAGE
  -----
  1. Copy this file and rename it to describe your agent, e.g. "code-reviewer-agent.md".
  2. Replace every {{PLACEHOLDER}} with your values.
  3. Use this document as the single source of truth when setting up the agent in
     Google AI Studio, Vertex AI Agent Builder, or via the Gemini API.
-->

---

## Agent Overview

| Field | Value |
|---|---|
| **Agent Name** | {{AGENT_NAME}} |
| **Version** | {{VERSION}} |
| **Domain** | {{DOMAIN}} |
| **Purpose** | {{ONE_LINE_PURPOSE}} |
| **Primary Users** | {{PRIMARY_USERS}} |

---

## Model Configuration

| Parameter | Value |
|---|---|
| **Model** | {{MODEL_ID}} (e.g. `gemini-2.0-flash`, `gemini-2.5-pro`) |
| **Temperature** | {{TEMPERATURE}} (0.0 – 2.0; lower = more deterministic) |
| **Top-P** | {{TOP_P}} (0.0 – 1.0) |
| **Top-K** | {{TOP_K}} |
| **Max Output Tokens** | {{MAX_OUTPUT_TOKENS}} |
| **Stop Sequences** | {{STOP_SEQUENCES}} (comma-separated, or "none") |

---

## System Instruction

> Reference the system instruction file from this repository, or paste content inline below.

**File:** `{{SYSTEM_INSTRUCTION_FILE_PATH}}`

```
{{SYSTEM_INSTRUCTION_CONTENT_OR_REFERENCE}}
```

---

## Tools and Function Calling

| Tool | Description | When to Call |
|---|---|---|
| {{TOOL_1_NAME}} | {{TOOL_1_DESCRIPTION}} | {{TOOL_1_TRIGGER}} |
| {{TOOL_2_NAME}} | {{TOOL_2_DESCRIPTION}} | {{TOOL_2_TRIGGER}} |

---

## Grounding

| Setting | Value |
|---|---|
| **Google Search Grounding** | {{ENABLED_DISABLED}} |
| **Vertex AI Search Datastore** | {{DATASTORE_ID_OR_NONE}} |
| **Dynamic Retrieval Threshold** | {{THRESHOLD_0_TO_1}} |

---

## Safety Settings

| Harm Category | Threshold |
|---|---|
| Harassment | {{THRESHOLD}} (BLOCK_NONE / BLOCK_LOW_AND_ABOVE / BLOCK_MEDIUM_AND_ABOVE / BLOCK_HIGH_AND_ABOVE) |
| Hate Speech | {{THRESHOLD}} |
| Sexually Explicit | {{THRESHOLD}} |
| Dangerous Content | {{THRESHOLD}} |

---

## Memory and Context

| Setting | Value |
|---|---|
| **Context Window Strategy** | {{STRATEGY}} (e.g. sliding window, summarisation) |
| **Session History Length** | {{N_TURNS}} turns |
| **Long-Term Memory Store** | {{MEMORY_STORE_OR_NONE}} |

---

## Skills Enabled

List the skill files from this repository that are appended to the system instruction:

* `{{SKILL_FILE_1}}`
* `{{SKILL_FILE_2}}`

---

## Example Interactions

### Turn 1

**User:** {{EXAMPLE_USER_1}}

**Agent:** {{EXAMPLE_AGENT_1}}

### Turn 2

**User:** {{EXAMPLE_USER_2}}

**Agent:** {{EXAMPLE_AGENT_2}}

---

## Evaluation Criteria

How to judge whether this agent is performing well:

* {{EVAL_CRITERION_1}}
* {{EVAL_CRITERION_2}}
* {{EVAL_CRITERION_3}}

---

## Changelog

| Date | Version | Change |
|---|---|---|
| {{DATE}} | {{VERSION}} | Initial configuration |
