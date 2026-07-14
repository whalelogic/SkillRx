# Earnings Report Summary Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Financial Markets |
| **Task** | Summarise and interpret a quarterly earnings report |
| **Recommended Model** | gemini-2.0-flash |
| **Recommended System Instruction** | `financial-markets/system-instructions/financial-researcher.md` |
| **Output Format** | Structured Markdown |

> **Disclaimer:** Educational analysis only. Not investment advice.

---

## Prompt

Please summarise and interpret the following earnings report for **{{COMPANY_NAME}} ({{TICKER}}) — {{FISCAL_PERIOD}}**.

### Earnings Release / Transcript

```
{{EARNINGS_CONTENT}}
```

(Paste the earnings press release, key financial tables, or relevant excerpts from the earnings call transcript.)

### Analyst Consensus (optional)

| Metric | Consensus Estimate | Actual |
|---|---|---|
| Revenue | {{REVENUE_EST}} | {{REVENUE_ACT}} |
| EPS (adjusted) | {{EPS_EST}} | {{EPS_ACT}} |
| Gross Margin | {{GM_EST}} | {{GM_ACT}} |
| {{OTHER_KPI}} | {{OTHER_EST}} | {{OTHER_ACT}} |

### Summary Task

1. **Headlines** — in 3 bullet points, state the most important results (beat/miss on key metrics, notable surprises).
2. **Financial performance** — revenue growth (YoY and QoQ), margin trends, EPS quality (one-time items, tax rate changes).
3. **Guidance** — what did management guide for next quarter/full year? Is it above/below prior consensus?
4. **Management commentary** — key themes from the CEO/CFO (demand environment, cost pressures, strategic initiatives).
5. **Key risks raised** — any new risk factors or deterioration in existing risks mentioned?
6. **Market reaction context** — if stock moved significantly on earnings, explain what likely drove the reaction based on the data.
7. **One-paragraph verdict** — balanced assessment of whether the results are incrementally positive, negative, or neutral for the long-term thesis.

*This is educational analysis only, not investment advice.*
