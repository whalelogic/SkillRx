# SEO Optimisation — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Writing |
| **Skill Name** | SEO Optimisation |
| **Compatible Roles** | Content Writer, Copywriter |
| **Gemini Feature Used** | Standard generation; Google Search Grounding for SERP research |

---

## Skill Description

Optimise written content for search engine discoverability by applying on-page SEO best practices — keyword integration, heading structure, meta descriptions, internal linking strategy, and content depth — without sacrificing readability.

## When to Activate

Activate this skill when the user:

* Asks for SEO-optimised content
* Wants to improve an existing page's search ranking
* Is planning a content calendar and needs keyword strategy input
* Asks for a meta description or title tag

## Core On-Page SEO Checklist

| Element | Best Practice |
|---|---|
| **Title tag** | Primary keyword near the start; 50–60 characters; unique per page |
| **Meta description** | Include primary keyword; 150–160 characters; clear value proposition + CTA |
| **H1** | Contains primary keyword; only one H1 per page; matches search intent |
| **H2/H3 headings** | Include secondary keywords naturally; structure the page for featured snippets |
| **Keyword density** | Primary keyword: 1–2%; avoid keyword stuffing; use semantic variations (LSI keywords) |
| **URL slug** | Short, descriptive, hyphen-separated, primary keyword included |
| **Internal links** | Link to 2–5 relevant pages on the same site; use descriptive anchor text |
| **Content length** | Match or exceed top-ranking pages for the target keyword |
| **Featured snippet optimisation** | Answer the target question directly in the first 40–50 words of the relevant section |
| **Image alt text** | Descriptive, includes keyword where natural |
| **Page speed** | Flag any heavy embeds or unoptimised images that could hurt Core Web Vitals |

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: SEO Optimisation

When writing or reviewing content for search:
1. Identify the primary keyword and 3–5 secondary/LSI keywords.
2. Place the primary keyword in: H1, first 100 words, at least one H2, meta description, URL slug.
3. Use secondary keywords naturally in H2/H3 headings and body text.
4. Write a meta description of 150–160 characters with the primary keyword and a clear CTA.
5. Suggest 2–3 internal link opportunities.
6. Structure at least one section to target a featured snippet (direct question + concise answer + list or table).
7. After writing, provide an SEO checklist with ✅ / ❌ for each element.
```

## Notes and Limitations

* SEO is driven by Google's algorithms, which evolve — recommendations reflect current best practices but may change.
* Keyword research tools (Ahrefs, SEMrush, Google Search Console) provide volume and difficulty data that this skill cannot replicate.
* Quality, relevance, and E-E-A-T (Experience, Expertise, Authoritativeness, Trustworthiness) matter more than keyword density.
