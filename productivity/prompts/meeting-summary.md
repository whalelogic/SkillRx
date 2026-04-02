# Meeting Summary Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Productivity |
| **Task** | Summarise a meeting transcript into decisions, actions, and context |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `productivity/system-instructions/executive-assistant.md` |
| **Output Format** | Structured Markdown meeting summary |

---

## Prompt

Please summarise the following meeting transcript into a structured meeting summary.

### Transcript

```
{{MEETING_TRANSCRIPT}}
```

(Paste the raw transcript, notes, or recording auto-transcript. If using Google Meet/Zoom auto-transcript, paste it directly.)

### Meeting Context

* **Meeting title:** {{TITLE}}
* **Date:** {{DATE}}
* **Attendees:** {{ATTENDEES}} (list names and roles if known)
* **Meeting purpose:** {{PURPOSE}} (e.g. weekly standup, product review, client call, design critique, incident review)

### Summary Task

1. **Executive summary** — 2–3 sentences capturing the purpose and outcome of the meeting.
2. **Key decisions made** — numbered list of decisions with brief rationale.
3. **Action items** — table with: Action | Owner | Due Date | Priority.
4. **Key discussion points** — 5–7 bullet points of the main topics discussed (not verbatim, synthesised).
5. **Open questions / parking lot** — items that were raised but not resolved.
6. **Next meeting** — date, purpose, and prep needed (if mentioned in the transcript).

### Output Format

```
## Meeting Summary: {{TITLE}}
📅 {{DATE}} | 👥 {{ATTENDEES}}

### Executive Summary
...

### Decisions
1. ...

### Action Items
| # | Action | Owner | Due Date | Priority |
|---|---|---|---|---|

### Key Discussion Points
- ...

### Open Questions
- ...

### Next Steps / Next Meeting
- ...
```
