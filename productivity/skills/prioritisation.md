# Prioritisation — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Productivity |
| **Skill Name** | Prioritisation |
| **Compatible Roles** | Productivity Coach, Project Manager, Executive Assistant |
| **Gemini Feature Used** | Standard generation |

---

## Skill Description

Apply structured prioritisation frameworks to any task list, backlog, or list of options to determine what to work on next and what to defer, delegate, or drop.

## Prioritisation Frameworks Quick Reference

### Eisenhower Matrix (Urgent / Important)

| | Urgent | Not Urgent |
|---|---|---|
| **Important** | Do Now (Quadrant I) | Schedule (Quadrant II) |
| **Not Important** | Delegate (Quadrant III) | Eliminate (Quadrant IV) |

**Best for:** personal task management, time audit, distinguishing reactive from proactive work.

### MoSCoW

* **Must Have** — non-negotiable; failure to include breaks the deliverable
* **Should Have** — important but not critical for launch
* **Could Have** — nice-to-have; include if time/resources allow
* **Won't Have (this time)** — explicitly out of scope for this cycle

**Best for:** feature/scope prioritisation, release planning, stakeholder alignment.

### RICE Score

`RICE = (Reach × Impact × Confidence) / Effort`

* **Reach** — number of people/customers affected per period
* **Impact** — how much it moves the key metric (0.25 / 0.5 / 1 / 2 / 3)
* **Confidence** — how sure are we? (100% = high / 80% = medium / 50% = low)
* **Effort** — person-months of work

**Best for:** product backlog prioritisation, comparing features with different audiences.

### Value vs. Effort Matrix

Plot items on a 2×2 grid:
* **High Value / Low Effort** — Quick wins: do these first
* **High Value / High Effort** — Major projects: plan and resource properly
* **Low Value / Low Effort** — Fill-ins: do if time allows
* **Low Value / High Effort** — Time sinks: eliminate

## When to Activate

Activate this skill when the user:

* Has a list of tasks and is unsure what to do first
* Needs to say no to some items and yes to others
* Is planning a sprint, release, or week and needs to cut scope
* Wants to communicate prioritisation decisions to stakeholders

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Prioritisation

When the user provides a list of tasks, features, or decisions to prioritise:
1. Identify the appropriate framework based on context (Eisenhower for personal tasks, RICE for product features, MoSCoW for scope decisions).
2. Apply the framework to score or categorise each item.
3. Present the prioritised list with brief justification for the top 3 and bottom 3.
4. Flag items where more information is needed to prioritise reliably.
5. End with a recommended "start here" action — the single most important next step.
```

## Notes and Limitations

* Prioritisation requires context — ask about goals, constraints, and success metrics before applying any framework.
* Numbers in RICE scoring are estimates; the value is in relative comparison, not absolute accuracy.
* Prioritisation frameworks help structure thinking but cannot replace strategic judgment.
