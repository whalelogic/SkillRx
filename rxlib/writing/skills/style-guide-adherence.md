# Style Guide Adherence — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Writing |
| **Skill Name** | Style Guide Adherence |
| **Compatible Roles** | Content Writer, Technical Writer, Copywriter |
| **Gemini Feature Used** | Standard generation |

---

## Skill Description

Apply the rules of a specified style guide consistently throughout any writing task, including preferred punctuation, capitalisation, number formatting, heading style, and terminology conventions.

## Supported Style Guides

### AP Stylebook (Associated Press)
Key rules:
- Numbers: spell out one through nine; use numerals for 10+
- No Oxford comma
- Capitalise job titles only when directly preceding the name: "President Biden" but "Joe Biden, president"
- Dates: Jan. 15, 2024 (abbreviated months with 6+ letters)
- Use % symbol with numerals: 5%

### Chicago Manual of Style (17th edition)
Key rules:
- Oxford comma required
- Numbers: spell out one through one hundred in text; use numerals in scientific/technical contexts
- Headline-style capitalisation for titles: Capitalise First, Last, and All Major Words
- Em dashes without spaces — like this

### Google Developer Documentation Style Guide
Key rules:
- Use second person ("you") always; avoid "we" and "the user"
- Present tense for procedures
- Imperative mood for instructions: "Click **Save**."
- Oxford comma required
- Capitalise only proper nouns and the first word of sentences in UI elements
- Avoid "please" in technical instructions

### Microsoft Writing Style Guide
Key rules:
- Friendly and conversational, but clear and direct
- Use "select" instead of "click" for accessibility
- Avoid jargon: use "use" not "utilise"
- Oxford comma required
- Sentence-style capitalisation for headings

## When to Activate

Activate this skill when the user:

* Specifies a style guide in their request
* Is writing content for a publication with known style requirements
* Wants consistency across a multi-part document

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Style Guide Adherence

Apply the {{STYLE_GUIDE_NAME}} throughout all writing output:
1. Follow the guide's rules for punctuation (especially Oxford comma and em dash usage).
2. Apply the guide's number style, date format, and capitalisation rules.
3. Use the guide's preferred terminology (e.g. "select" vs. "click" in Microsoft style).
4. After producing content, note any rule applications that required a judgement call.
5. If the user's input violates the style guide, correct it silently and note the correction.
```

## Notes and Limitations

* Style guides are complex — this skill covers the most commonly applied rules, not the full guide.
* Some rules conflict between guides; always clarify which guide takes precedence.
* Custom brand style guides should be provided by the user in full or as a summary of key rules.
