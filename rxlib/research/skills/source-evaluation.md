# Source Evaluation — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Research |
| **Skill Name** | Source Evaluation |
| **Compatible Roles** | Research Assistant, Academic Researcher, Investigative Analyst |
| **Gemini Feature Used** | Standard generation; Google Search Grounding for author/publisher lookup |

---

## Skill Description

Apply the SIFT (Stop, Investigate, Find better coverage, Trace claims) and CRAAP (Currency, Relevance, Authority, Accuracy, Purpose) frameworks to systematically assess the quality and credibility of any source before citing it.

## When to Activate

Activate this skill when the user:

* Provides a source and asks whether it is credible
* Is building a bibliography and wants to vet sources
* Encounters a claim that seems dubious and wants to trace it to its origin
* Is distinguishing between primary, secondary, and tertiary sources

## Behaviour

When this skill is active:

1. **Stop** — pause before accepting the source at face value; note any initial red flags.
2. **Investigate the source** — look up the author credentials, publisher reputation, and any known biases.
3. **Find better coverage** — check if the claim is corroborated by other independent, high-quality sources.
4. **Trace claims** — follow citations back to original research to check if secondary sources accurately represent primary findings.
5. Score each CRAAP dimension (Currency, Relevance, Authority, Accuracy, Purpose) as High / Medium / Low.
6. Give an overall verdict: Recommended / Use with caveats / Not recommended.

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Source Evaluation (SIFT + CRAAP)

Before citing or relying on any source:
1. Apply SIFT: Stop, Investigate the source (author, publisher, date), Find better coverage, Trace claims to originals.
2. Score CRAAP: Currency (is it recent enough?), Relevance (does it address the question?), Authority (qualified author/publisher?), Accuracy (evidence-based?), Purpose (any bias or motive?).
3. State your verdict: Recommended / Use with caveats / Not recommended.
4. Never cite a source you cannot verify — flag it as "unverified" and recommend the user check directly.
```

## Notes and Limitations

* The model cannot access paywalled academic databases — it can assess metadata and abstracts but not full-text papers.
* Author and publisher credibility checks benefit from Google Search grounding.
* Even high-quality sources can contain errors — critical reading is always necessary.
