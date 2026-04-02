# Editing and Proofreading Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Writing |
| **Task** | Edit and proofread content for clarity, style, and correctness |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `writing/system-instructions/content-writer.md` |
| **Output Format** | Edited document + change summary |

---

## Prompt

Please edit and proofread the following text. The goal is {{EDIT_GOAL}}.

### Edit Type

Select one or more:
- [ ] **Proofreading** — spelling, grammar, punctuation, and typographical errors only
- [ ] **Line editing** — sentence-level clarity, concision, flow, and word choice
- [ ] **Developmental editing** — structure, argument, coherence, and completeness
- [ ] **Style guide alignment** — align with {{STYLE_GUIDE}} (AP / Chicago / Microsoft / Google / custom)
- [ ] **Tone adjustment** — adjust the tone to be more {{TARGET_TONE}} (formal / conversational / empathetic / confident)

### Original Text

```
{{TEXT}}
```

### Context

* **Document type:** {{DOCUMENT_TYPE}} (e.g. blog post, technical doc, email, academic paper, marketing copy)
* **Target audience:** {{AUDIENCE}}
* **Preserve author voice:** {{PRESERVE_VOICE}} (yes — make minimal changes / no — rewrite freely for clarity)

### Task

1. Produce the edited version of the text.
2. Provide a summary of changes made, grouped by category (grammar, clarity, structure, tone).
3. Flag any passages where the meaning was unclear and explain your interpretation.
4. If developmental editing is requested, note any structural issues (missing introduction, weak conclusion, unsupported claims) separately from line-level edits.

### Output Requirements

* Deliver the full edited text first, then the change summary.
* In the change summary, use a table: `Original | Edited | Reason`
* Do not change meaning unless explicitly instructed and flagged.
* If asked to preserve voice, prefer rephrasing over rewriting.
