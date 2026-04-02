# Market Analysis Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Financial Markets |
| **Task** | Broad market conditions analysis |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `financial-markets/system-instructions/market-analyst.md` |
| **Output Format** | Structured Markdown |

> **Disclaimer:** This prompt generates educational market analysis only, not investment advice.

---

## Prompt

Please provide a structured analysis of current market conditions across the following asset classes and key themes.

### Scope

* **Asset classes to cover:** {{ASSET_CLASSES}} (e.g. "US equities, US Treasuries, Gold, USD Index" or "all major asset classes")
* **Time horizon:** {{TIME_HORIZON}} (e.g. "next 30 days", "Q3 2025", "medium-term 6–12 months")
* **Geographic focus:** {{GEOGRAPHY}} (e.g. "US", "Global", "Emerging Markets")

### Context (optional — provide current data if available)

```
{{CURRENT_DATA}}
```
(e.g. paste in recent index levels, yield levels, CPI print, Fed statement excerpts)

### Analysis Task

1. **Macro backdrop** — summarise the current economic regime (expansion, contraction, stagflation, etc.) and the dominant central bank policy stance.
2. **Asset class overview** — for each asset class in scope, assess: current trend, key drivers, and primary risk.
3. **Cross-asset correlations** — identify any notable correlation breakdowns or regime shifts.
4. **Key upcoming catalysts** — list the top 5 scheduled events (earnings, data releases, central bank meetings) that could move markets.
5. **Bull and bear scenarios** — for each scenario, describe the trigger and expected market impact.
6. **Risk factors** — rank the top 3 tail risks by probability and potential severity.

### Output Format

```
## Macro Backdrop
...

## Asset Class Summary
| Asset Class | Trend | Key Driver | Risk |
|---|---|---|---|

## Cross-Asset Correlations
...

## Key Upcoming Catalysts
| Date | Event | Expected Market Impact |
|---|---|---|

## Scenarios
### Bull Case
...
### Bear Case
...

## Top Tail Risks
| # | Risk | Probability | Severity | Trigger |
|---|---|---|---|---|

---
*This is educational analysis only, not investment advice.*
```
