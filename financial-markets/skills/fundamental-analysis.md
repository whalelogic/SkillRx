# Fundamental Analysis — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Financial Markets |
| **Skill Name** | Fundamental Analysis |
| **Compatible Roles** | Market Analyst, Financial Researcher |
| **Gemini Feature Used** | Standard generation; Code Execution for ratio calculations |

---

## Skill Description

Evaluate a company's financial health, competitive position, and intrinsic value using financial statement analysis, ratio analysis, and valuation models.

## When to Activate

Activate this skill when the user:

* Provides financial data (income statement, balance sheet, cash flows) and asks for interpretation
* Asks about a company's valuation relative to peers or its own history
* Wants to build a simple DCF model
* Asks about financial ratios and what they mean

## Key Ratios Quick Reference

### Profitability
| Ratio | Formula | Healthy Range (varies by sector) |
|---|---|---|
| Gross Margin | Gross Profit / Revenue | >40% for tech/software; >20% for retail |
| EBITDA Margin | EBITDA / Revenue | >20% generally attractive |
| Net Margin | Net Income / Revenue | Positive and expanding |
| ROIC | NOPAT / Invested Capital | >WACC (typically >10%) |
| ROE | Net Income / Avg. Equity | >15% sustained |

### Leverage
| Ratio | Formula | Signal |
|---|---|---|
| Net Debt / EBITDA | (Debt − Cash) / EBITDA | <3x comfortable; >5x elevated |
| Interest Coverage | EBIT / Interest Expense | >3x healthy |
| Debt / Equity | Total Debt / Total Equity | Context-dependent |

### Valuation
| Multiple | Formula | Notes |
|---|---|---|
| P/E | Price / EPS | Compare to sector avg and growth rate |
| EV/EBITDA | Enterprise Value / EBITDA | Removes capital structure effect |
| P/S | Market Cap / Revenue | Useful for unprofitable growth companies |
| P/FCF | Market Cap / Free Cash Flow | More cash-flow grounded than P/E |
| PEG | P/E / EPS Growth Rate | <1 potentially undervalued |

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Fundamental Analysis

When analysing a company's financials:
1. Calculate and contextualise the key profitability, leverage, and valuation ratios.
2. Compare ratios to sector peers and 5-year historical averages.
3. Identify quality-of-earnings concerns (non-cash items, one-time charges, working capital changes).
4. Build a bear/base/bull DCF with explicit assumptions for revenue growth, EBIT margin, WACC, and terminal growth rate.
5. Present valuation as a range, not a point estimate.
6. Always include: "This is educational analysis only, not investment advice."
```

## Notes and Limitations

* Always note that ratios must be compared to industry peers — there is no universally "good" value.
* DCF models are highly sensitive to terminal growth rate and WACC assumptions — show a sensitivity table.
* Do not fabricate financial data — only use numbers provided by the user or retrieved via grounding.
