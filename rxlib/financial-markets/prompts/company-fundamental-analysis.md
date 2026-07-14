# Company Fundamental Analysis Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Financial Markets |
| **Task** | Fundamental analysis of a company |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `financial-markets/system-instructions/financial-researcher.md` |
| **Output Format** | Investment memo (Markdown) |

> **Disclaimer:** Educational analysis only. Not investment advice.

---

## Prompt

Please perform a fundamental analysis of **{{COMPANY_NAME}} ({{TICKER}})**.

### Financial Data

Paste the relevant financial data below (from SEC filings, Bloomberg, or Yahoo Finance):

```
{{FINANCIAL_DATA}}
```

Suggested data to include:
* Last 3 years of annual income statements (Revenue, Gross Profit, EBITDA, Net Income, EPS)
* Last 3 years of cash flow statements (Operating CF, Capex, Free Cash Flow)
* Latest balance sheet (Total Assets, Total Debt, Cash, Equity)
* Current stock price and market cap
* Analyst consensus estimates for next 2 years (optional)

### Company Context

* **Sector / Industry:** {{SECTOR}}
* **Primary business:** {{BUSINESS_DESCRIPTION}}
* **Main competitors:** {{COMPETITORS}}
* **Recent news / events:** {{RECENT_NEWS}}

### Analysis Task

1. **Business quality assessment** — evaluate the competitive moat, revenue quality, and management track record.
2. **Financial health** — analyse profitability trends, balance sheet strength, and free cash flow generation.
3. **Valuation** — calculate and contextualise key multiples vs. sector peers and 5-year historical averages; include a basic DCF with explicit assumptions.
4. **Bear case** — identify the top 3 risks that could impair the investment thesis.
5. **Catalysts** — list upcoming events (earnings, product launches, regulatory decisions) that could re-rate the stock.
6. **Summary** — balanced assessment of the company's outlook without a buy/sell recommendation.

### Output Format

Structure as a research note:
- Company Overview
- Investment Thesis (bull case)
- Financial Summary table
- Valuation (comps + DCF)
- Bear Case / Key Risks
- Catalysts
- Conclusion

End with: *"This is educational analysis only, not investment advice."*
