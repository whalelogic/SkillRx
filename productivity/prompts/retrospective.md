# Retrospective Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Productivity |
| **Task** | Facilitate a team retrospective and produce structured output |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `productivity/system-instructions/project-manager.md` |
| **Output Format** | Structured Markdown retrospective report |

---

## Prompt

Please facilitate a {{RETRO_FORMAT}} retrospective based on the following team input, and produce a structured output.

### Team Input

Paste raw retrospective notes, sticky notes, or responses:

```
{{RETRO_INPUT}}
```

### Retrospective Configuration

* **Format:** {{RETRO_FORMAT}} (Start-Stop-Continue / 4Ls (Liked-Learned-Lacked-Longed for) / Mad-Sad-Glad / Plus-Delta / Sailboat / 5 Whys for a specific issue)
* **Sprint / period:** {{PERIOD}}
* **Team size:** {{TEAM_SIZE}}
* **Facilitator goal:** {{GOAL}} (e.g. "improve deployment process", "reduce meetings", "general team health check")

### Facilitation Task

1. **Organise the input** — group related items into themes; remove duplicates; clarify ambiguous items.
2. **Synthesise insights** — identify the 3–5 most important patterns or themes across all input.
3. **Prioritise** — rank the themes by impact on team productivity and morale.
4. **Generate action items** — for the top 2–3 themes, propose concrete, measurable actions with suggested owners.
5. **Positive recognition** — highlight specific wins or behaviours worth celebrating.
6. **Produce the retrospective summary** — a clean document the team can review and act on.

### Output Format

```
## Retrospective Summary: {{PERIOD}}
👥 Team: {{TEAM}} | 📅 Date: {{DATE}}

### What Went Well (Keep Doing)
- ...

### What Could Be Improved (Stop / Change)
- ...

### What We Should Try (Start / Experiment)
- ...

### Key Themes
1. [Theme] — [Evidence summary]

### Action Items
| # | Action | Owner | Due | Success Criterion |
|---|---|---|---|---|

### Recognition
- ...

### Next Retrospective Focus Area
- ...
```
