# Weekly Planning Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Productivity |
| **Task** | Plan a productive week aligned with priorities |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `productivity/system-instructions/productivity-coach.md` |
| **Output Format** | Weekly plan (Markdown) |

---

## Prompt

Please help me plan my upcoming week for maximum focus and productivity.

### My Situation

* **Role / job type:** {{ROLE}} (e.g. "software engineer in a startup", "freelance designer", "team lead with 6 direct reports")
* **Work schedule:** {{SCHEDULE}} (e.g. "Mon–Fri 9–5", "flexible async", "4-day week")
* **Deep work capacity:** {{DEEP_WORK_HOURS}} hours per day when uninterrupted work is possible

### Goals for This Week

**Top priorities (what must get done this week):**
1. {{PRIORITY_1}}
2. {{PRIORITY_2}}
3. {{PRIORITY_3}}

**Secondary tasks (would be good to complete):**
{{SECONDARY_TASKS}}

### Existing Commitments

List fixed meetings, calls, or commitments already in the calendar:

```
{{EXISTING_COMMITMENTS}}
```

(e.g. "Mon 10am – team standup (30 min)", "Wed 2pm – client call (1 hr)")

### Context and Constraints

* **Current energy level / state:** {{ENERGY}} (e.g. "well-rested and motivated", "recovering from a busy sprint", "dealing with a personal situation")
* **Tasks to avoid this week:** {{AVOID}}
* **Preferred working style:** {{WORKING_STYLE}} (e.g. "batch similar tasks", "morning for creative, afternoon for admin", "Pomodoro 25/5")

### Planning Task

1. Assess the feasibility of completing all top priorities given available time.
2. Allocate deep work blocks for the top priorities, protecting them from meetings.
3. Schedule secondary tasks in shallow-work slots (low-energy periods, after meetings).
4. Identify any meetings or commitments that could be declined, delegated, or shortened.
5. Build in buffer (at least 20% unscheduled time) for unexpected work.
6. Define the "Friday win" — what single outcome would make this week a success?

### Output Format

```
## Weekly Plan: Week of {{DATE}}

### Theme
[One sentence describing the week's focus]

### Top 3 Priorities (MIT)
1. ...

### Suggested Daily Schedule
| Day | Morning Block | Afternoon Block | Notes |
|---|---|---|---|

### Time Budget
| Category | Hours Allocated | Tasks |
|---|---|---|

### Commitments to Decline/Defer
- ...

### Friday Win
[The one result that defines a successful week]

### End-of-Week Review Checklist
- [ ] ...
```
