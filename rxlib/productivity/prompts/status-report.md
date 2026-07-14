# Status Report Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Productivity |
| **Task** | Generate a stakeholder status report from raw project notes |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `productivity/system-instructions/project-manager.md` |
| **Output Format** | Structured Markdown status report |

---

## Prompt

Please generate a professional stakeholder status report from the following raw project notes.

### Raw Notes / Updates

```
{{RAW_NOTES}}
```

(Paste unstructured notes, Slack messages, bullet points, or a brain dump of what happened this period.)

### Report Configuration

* **Project name:** {{PROJECT_NAME}}
* **Reporting period:** {{PERIOD}} (e.g. "Week ending 2025-03-07", "Sprint 24", "Q1 2025")
* **Overall status:** {{STATUS}} (🟢 On Track / 🟡 At Risk / 🔴 Off Track / ⚪ On Hold — or leave blank to have the model assess from notes)
* **Audience:** {{AUDIENCE}} (e.g. "executive sponsor and C-suite", "project steering committee", "engineering team leads")
* **Report format:** {{FORMAT}} (concise executive summary / detailed technical report / board-level slide narrative)

### Report Task

1. **Executive summary** — 2–3 sentence status narrative capturing progress, status, and the key issue if at risk.
2. **Accomplishments this period** — bullet list of completed milestones and notable deliverables.
3. **Planned vs. actual** — are we on schedule, on budget, and within scope? Note any variances.
4. **Issues and risks** — table with severity, description, impact, and mitigation/owner.
5. **Next period plan** — what will be accomplished in the next period?
6. **Decisions required** — flag any decisions the audience needs to make for the project to continue.
7. **Metrics** — if any metrics were mentioned in the notes, present them in a table.

### Output Requirements

* Translate raw, informal language into professional stakeholder-appropriate prose.
* Group related items — do not just convert bullets into bullets.
* If the status is "at risk" or "off track", lead with the issue and proposed path forward.
