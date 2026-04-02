# Social Media Content Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Writing |
| **Task** | Create platform-optimised social media posts |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `writing/system-instructions/copywriter.md` |
| **Output Format** | Markdown |

---

## Prompt

Please create social media content for the following topic or source material, optimised for each specified platform.

### Source Content / Topic

```
{{SOURCE_CONTENT_OR_TOPIC}}
```

(Paste a blog post, press release, product update, or just describe the topic.)

### Campaign Brief

* **Brand / author:** {{BRAND_OR_AUTHOR}}
* **Brand voice:** {{BRAND_VOICE}} (e.g. "professional and authoritative", "playful and irreverent", "empathetic and supportive")
* **Campaign goal:** {{GOAL}} (e.g. drive traffic, increase brand awareness, generate leads, spark conversation)
* **Target audience:** {{AUDIENCE}}
* **Link to include:** {{URL}} (or "none")

### Platforms and Post Requirements

| Platform | Max Characters | Hashtags | CTA |
|---|---|---|---|
| LinkedIn | 3,000 | 3–5 professional | {{CTA_LINKEDIN}} |
| X (Twitter) | 280 | 1–3 | {{CTA_X}} |
| Instagram | 2,200 | 5–10 + niche | {{CTA_INSTAGRAM}} |
| Facebook | 500 recommended | 1–3 | {{CTA_FACEBOOK}} |

Select the platforms you need: {{SELECTED_PLATFORMS}}

### Task

For each selected platform:
1. Write a platform-native post that fits the character limit and content style of that platform.
2. Include appropriate hashtags and emojis (if appropriate for the brand voice).
3. Include a clear CTA.
4. Suggest the best time to post (general best-practice guidance).
5. Provide one A/B variant for the highest-priority platform.

### Output Format

For each platform:
```
## [Platform Name]

[Post copy]

[Hashtags]

📅 Best time to post: ...

---
### Variant B:
[Alternative post copy]
```
