# Blog Post Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Writing |
| **Task** | Write or outline a blog post |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `writing/system-instructions/content-writer.md` |
| **Output Format** | Markdown |

---

## Prompt

Please {{WRITE_OR_OUTLINE}} a blog post on the following topic.

### Topic

**{{TOPIC}}**

### Content Brief

* **Target audience:** {{AUDIENCE}} (e.g. "small business owners with no marketing background", "senior software engineers")
* **Primary keyword:** {{PRIMARY_KEYWORD}} (for SEO; leave blank if not relevant)
* **Search intent:** {{SEARCH_INTENT}} (informational / navigational / commercial / transactional)
* **Tone:** {{TONE}} (e.g. professional, conversational, humorous, authoritative)
* **Target word count:** {{WORD_COUNT}}
* **Key points to cover:** {{KEY_POINTS}}
* **Key points to avoid:** {{AVOID}}
* **CTA at the end:** {{CTA}} (e.g. "Subscribe to newsletter", "Download the guide", "Leave a comment")

### Task

1. Write an attention-grabbing hook (first 1–2 sentences or paragraph).
2. Write a brief introduction that establishes why this topic matters and previews what the reader will learn.
3. Develop the body with {{N_SECTIONS}} clearly headed sections, each covering one key idea.
4. Write a conclusion that summarises the key takeaways and includes the CTA.
5. Suggest a meta description (150–160 characters) optimised for the primary keyword.
6. Suggest 3 headline variants the user can test.

### Output Requirements

* Use H2 for main sections, H3 for subsections.
* Include {{INCLUDE_EXAMPLES}} (yes / no) concrete examples, case studies, or statistics in each major section.
* Reading level: {{READING_LEVEL}} (e.g. "accessible to general audience (Grade 10 equivalent)")
* Bold the key takeaway in each section.
