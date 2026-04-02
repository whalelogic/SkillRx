# Tone Adjustment — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Writing |
| **Skill Name** | Tone Adjustment |
| **Compatible Roles** | Content Writer, Copywriter, Technical Writer, Creative Writer |
| **Gemini Feature Used** | Standard generation |

---

## Skill Description

Rewrite or adjust any piece of content to match a specified tone, preserving the meaning and information while changing the register, formality level, and emotional quality.

## Tone Spectrum Reference

| Tone | Characteristics | Typical Use |
|---|---|---|
| **Formal / Academic** | Passive voice acceptable, no contractions, precise vocabulary, complex sentences | Academic papers, legal documents, official reports |
| **Professional** | Active voice, no slang, clear and direct, first/second person | Business communication, technical docs, corporate blogs |
| **Semi-formal** | Conversational but polished, light use of contractions, relatable examples | Company blogs, newsletters, LinkedIn |
| **Conversational** | Contractions, colloquialisms, short sentences, casual examples | Consumer blogs, how-to guides, social media captions |
| **Empathetic** | Warm, validating language, acknowledges the reader's situation, "you" focused | Support content, mental health, non-profit, HR |
| **Authoritative** | Confident, no hedging, declarative statements, data-backed | Expert commentary, whitepapers, analyst reports |
| **Playful / Witty** | Wordplay, humour, unexpected metaphors, cultural references | Brand copy, entertainment, consumer apps |
| **Urgent / Persuasive** | Action verbs, scarcity/benefit framing, short impactful sentences | Ad copy, fundraising, crisis communications |

## When to Activate

Activate this skill when the user:

* Asks to make content "more formal / casual / friendly / direct"
* Provides a brand voice guide and asks the writing to match it
* Wants to repurpose content for a different audience or channel

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Tone Adjustment

When asked to change the tone of content:
1. Identify the current tone and the target tone from the user's instruction.
2. Rewrite the content adjusting: sentence length, vocabulary complexity, use of contractions, active vs. passive voice, use of pronouns ("you"/"we"), and emotional register.
3. Preserve all factual content and meaning — tone changes only.
4. Produce the rewritten content followed by a one-paragraph note explaining the key adjustments made.
```

## Notes and Limitations

* Tone adjustment must preserve meaning — flag any case where the requested tone change would distort the message.
* Extreme formality changes (e.g. academic to playful) may require structural rewriting as well as word-level changes.
