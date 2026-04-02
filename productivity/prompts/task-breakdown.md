# Task Breakdown Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Productivity |
| **Task** | Break a high-level goal into actionable tasks with estimates |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `productivity/system-instructions/project-manager.md` |
| **Output Format** | Structured Markdown task list |

---

## Prompt

Please break down the following high-level goal into actionable tasks with effort estimates and logical sequencing.

### Goal

**{{GOAL}}**

(e.g. "Launch a beta version of our API", "Write and publish a comprehensive technical tutorial", "Migrate our database from MySQL to PostgreSQL")

### Context

* **Deadline / target date:** {{DEADLINE}}
* **Team / people involved:** {{TEAM}} (e.g. "just me", "2 engineers and 1 designer", "cross-functional team of 8")
* **Available time per week:** {{TIME_AVAILABLE}} (e.g. "20 hours/week total", "50% of my work week")
* **Skills / tools available:** {{SKILLS_TOOLS}}
* **Known constraints:** {{CONSTRAINTS}} (e.g. "waiting on vendor API access", "can't start until design is approved")
* **Definition of done:** {{DONE_CRITERIA}}

### Breakdown Task

1. **Work Breakdown Structure (WBS)** — decompose the goal into phases, then tasks, then sub-tasks where needed.
2. **Effort estimates** — for each task, provide an effort estimate in hours or days (use three-point PERT: Optimistic / Most Likely / Pessimistic).
3. **Dependencies** — identify which tasks must be completed before others can start.
4. **Critical path** — identify the longest chain of dependent tasks that determines the minimum project duration.
5. **Milestones** — identify 3–5 key checkpoints that signal meaningful progress.
6. **Risk flags** — note any tasks with high uncertainty or dependency risk.

### Output Format

```
## Work Breakdown: {{GOAL}}

### Phases and Tasks
| # | Phase | Task | Effort (Optimistic/Likely/Pessimistic) | Owner | Depends On |
|---|---|---|---|---|---|

### Critical Path
Task A → Task B → Task C → Task D (estimated duration: X days)

### Milestones
| Milestone | Target Date | Criteria |
|---|---|---|

### Risk Flags
- ...
```
