# Risk Assessment Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Financial Markets |
| **Task** | Risk assessment of a position or strategy |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `financial-markets/system-instructions/quantitative-analyst.md` |
| **Output Format** | Structured Markdown with tables |

> **Disclaimer:** Educational analysis only. Not investment advice.

---

## Prompt

Please assess the risk profile of the following {{POSITION_OR_STRATEGY}}.

### Position / Strategy Details

```
{{POSITION_DETAILS}}
```

Include:
* Security/instrument name and ticker
* Long or short
* Position size (shares, contracts, or % of portfolio)
* Entry price / current price
* For options: strike, expiration, type (call/put), quantity

### Context

* **Portfolio context:** {{PORTFOLIO_CONTEXT}} (e.g. "This is a standalone position" or "This is a hedge against a long equity portfolio")
* **Investment horizon:** {{HORIZON}}
* **Risk capital available:** {{RISK_CAPITAL}} (maximum tolerable loss in % or absolute)

### Risk Assessment Task

1. **Market risk** — estimate the position's sensitivity to price moves (delta), volatility changes (vega, if applicable), and time decay (theta, if applicable). Provide scenarios: -20%, -10%, -5%, +5%, +10%, +20% underlying move.
2. **Liquidity risk** — comment on bid/ask spread, average daily volume, and ability to exit in a stress scenario.
3. **Concentration risk** — what % of capital is at risk if the position goes to zero?
4. **Correlation risk** — how is this position correlated to the broader portfolio or major indices?
5. **Tail risk** — what is the maximum realistic loss (consider historical drawdowns and stress scenarios)?
6. **Risk/reward summary** — present as a simple table: Scenario | Probability (rough) | P&L Impact.
7. **Risk management suggestions** — position sizing guidance, stop-loss levels, hedging options.

*This is educational analysis only, not investment advice. Consult a qualified financial professional.*
