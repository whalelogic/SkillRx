# Action Item Extraction — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Productivity |
| **Skill Name** | Action Item Extraction |
| **Compatible Roles** | Executive Assistant, Project Manager |
| **Gemini Feature Used** | Standard generation |

---

## Skill Description

Scan any text (meeting transcripts, emails, Slack threads, documents) and extract all commitments, action items, and follow-up tasks — with owners and due dates where mentioned.

## When to Activate

Activate this skill when the user:

* Pastes a meeting transcript, email thread, or conversation and asks "what are the action items?"
* Wants to convert unstructured project notes into a task list
* Asks to follow up on a document or discussion

## Behaviour

When this skill is active:

1. Read the full text before extracting — action items may be scattered or implied.
2. Extract both **explicit actions** ("John will send the report by Friday") and **implicit commitments** ("we need to follow up on the API issue" — no owner named).
3. For each action item, capture:
   - **What** — a clear, verb-led description of the task
   - **Who** — the named owner (or "Unassigned" if no owner was specified)
   - **When** — the due date or timeframe (or "Not specified")
   - **Context** — a one-phrase note about why it matters (optional but helpful)
4. Flag tasks with no owner as a risk — unowned tasks rarely get done.
5. After the table, list any open questions that need resolution before actions can proceed.

## Output Format

```
## Action Items

| # | Action | Owner | Due Date | Context |
|---|---|---|---|---|
| 1 | [Verb-led action description] | [Name / Unassigned] | [Date / Not specified] | [Brief context] |

### ⚠️ Unowned Actions (require owner assignment)
- ...

### Open Questions
- ...
```

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Action Item Extraction

When processing meeting notes, email threads, or conversations:
1. Extract all explicit commitments and implicit follow-up tasks.
2. Capture: What (verb-led), Who (owner), When (due date), Context (why it matters).
3. Flag unowned actions — they are at highest risk of being dropped.
4. Present all action items in a Markdown table.
5. List open questions that must be resolved before actions can proceed.
```

## Notes and Limitations

* Implicit commitments (e.g. "we should look into that") require judgment — flag them as lower confidence.
* Due dates inferred from context (e.g. "before the next sprint" → "in ~2 weeks") should be noted as approximate.
* In long transcripts, watch for commitments made early in the conversation that were not revisited.
