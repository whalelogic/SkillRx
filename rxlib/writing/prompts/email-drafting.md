# Email Drafting Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Writing |
| **Task** | Draft a professional email |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `writing/system-instructions/content-writer.md` |
| **Output Format** | Markdown (plain text for pasting into email client) |

---

## Prompt

Please draft a professional email for the following situation.

### Email Brief

* **From:** {{SENDER_NAME}}, {{SENDER_ROLE}}
* **To:** {{RECIPIENT_NAME}}, {{RECIPIENT_ROLE}}
* **Relationship context:** {{RELATIONSHIP}} (e.g. "first contact", "existing client", "direct report", "senior stakeholder")
* **Purpose / goal:** {{PURPOSE}} (e.g. "request a meeting", "follow up on a proposal", "deliver feedback", "announce a decision")
* **Tone:** {{TONE}} (formal / semi-formal / direct / empathetic / urgent)
* **Key message:** {{KEY_MESSAGE}} (the one thing the recipient must take away)

### Background and Key Points to Include

{{BACKGROUND_AND_POINTS}}

### Constraints

* **Maximum length:** {{MAX_LENGTH}} (e.g. 150 words, 5 sentences, no limit)
* **Subject line variants:** provide {{N_SUBJECT_VARIANTS}} subject line options
* **CTA (call to action):** {{CTA}} (e.g. "Schedule a 30-minute call", "Confirm receipt", "Approve the attached document")

### Task

1. Draft the email with a clear structure: opening → context → key message → CTA → closing.
2. Match the tone to the relationship and purpose.
3. Write {{N_SUBJECT_VARIANTS}} subject line options (aim for specificity and clarity — avoid vague subjects like "Following up").
4. Optionally, provide a shorter "TL;DR version" if the email exceeds 150 words.

### Output Format

```
Subject: [Subject line option 1]
Alt Subject: [Subject line option 2]

---

[Email body]

---

TL;DR (if applicable): [1–2 sentence summary of the email's purpose and CTA]
```
