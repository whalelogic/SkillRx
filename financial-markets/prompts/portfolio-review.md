# Portfolio Review Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Financial Markets |
| **Task** | Portfolio concentration and diversification review |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `financial-markets/system-instructions/market-analyst.md` |
| **Output Format** | Structured Markdown with tables |

> **Disclaimer:** Educational analysis only. Not investment advice.

---

## Prompt

Please review the following investment portfolio for risk concentration, diversification quality, and alignment with the stated investment objectives.

### Portfolio Holdings

```
{{PORTFOLIO_DATA}}
```

Format: one holding per line — `Ticker | Name | Asset Class | Sector | Weight (%) | Cost Basis (optional)`

Example:
```
AAPL | Apple Inc. | Equity | Technology | 12.5% | $145.00
TLT  | iShares 20+ Year Treasury Bond ETF | Fixed Income | Government Bonds | 8.0% | $102.00
GLD  | SPDR Gold Shares | Commodity | Gold | 5.0% | $185.00
```

### Investor Profile

* **Investment objective:** {{OBJECTIVE}} (e.g. capital growth, income, capital preservation, balanced)
* **Time horizon:** {{TIME_HORIZON}} (e.g. 5 years, 20 years, retirement in 10 years)
* **Risk tolerance:** {{RISK_TOLERANCE}} (low / moderate / high)
* **Portfolio size:** {{PORTFOLIO_SIZE}} (optional — used for fee/impact context)
* **Tax jurisdiction:** {{TAX_JURISDICTION}} (optional)

### Review Task

1. **Concentration analysis** — identify positions >10% and factor/sector concentrations.
2. **Diversification quality** — assess asset class, geographic, sector, and factor diversification.
3. **Risk profile** — estimate the portfolio's expected volatility and drawdown profile relative to a 60/40 benchmark.
4. **Alignment check** — does the portfolio's actual risk profile match the stated objective and risk tolerance?
5. **Observations** — highlight any overlapping holdings (e.g. owning both an S&P 500 ETF and large-cap tech stocks) and fee drag.
6. **Suggested adjustments** — frame as rebalancing considerations (not buy/sell recommendations), respecting the investor's constraints.

*This is educational analysis only, not investment advice.*
