# User Flow Design Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | UI/UX Development |
| **Task** | Design or evaluate a user flow |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `ui-ux-development/system-instructions/ux-designer.md` |
| **Output Format** | Mermaid diagram + Markdown analysis |

---

## Prompt

Please {{DESIGN_OR_EVALUATE}} the user flow for the following task.

### Task to Complete

**{{USER_TASK}}**

(e.g. "A new user signs up, verifies their email, and completes onboarding", "A returning user resets their forgotten password")

### Product Context

* **Product type:** {{PRODUCT_TYPE}} (web app / mobile app / multi-platform)
* **User type:** {{USER_TYPE}} (e.g. first-time user, power user, admin)
* **Entry point:** {{ENTRY_POINT}} (e.g. marketing landing page, deep link, push notification)
* **Success state:** {{SUCCESS_STATE}} (what does "task complete" look like?)

### Existing Flow (for evaluation tasks)

```
{{EXISTING_FLOW_DESCRIPTION_OR_STEPS}}
```

### Design / Evaluation Task

1. Map the complete flow from entry point to success state, including: all screens/states, decision points, error states, and happy-path vs. recovery paths.
2. Present the flow as a Mermaid flowchart.
3. Analyse the flow for:
   - **Steps to completion** — is the path as short as it could be?
   - **Cognitive load** — are there decision points that could be eliminated or simplified?
   - **Error prevention** — where might users make mistakes, and how is recovery handled?
   - **Drop-off risk** — which steps have the highest friction or likelihood of abandonment?
4. Recommend improvements with justification.
5. Define the key metrics to measure the flow's effectiveness (task completion rate, time-on-task, error rate, drop-off by step).

### Output Format

```
## User Flow: [Task Name]

### Flow Diagram
```mermaid
flowchart TD
  ...
```

### Flow Analysis
| Step | Screen/State | User Action | Decision? | Friction Risk | Notes |
|---|---|---|---|---|---|

### Improvement Recommendations
1. ...

### Success Metrics
| Metric | Definition | Target |
|---|---|---|
```
