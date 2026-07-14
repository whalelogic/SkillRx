# Technical Analysis — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Financial Markets |
| **Skill Name** | Technical Analysis |
| **Compatible Roles** | Market Analyst, Quantitative Analyst |
| **Gemini Feature Used** | Multimodal (vision) — chart images; Code Execution — indicator calculation |

---

## Skill Description

Interpret price charts, identify trend structure, support/resistance levels, and common technical patterns; calculate and interpret momentum indicators.

## When to Activate

Activate this skill when the user:

* Shares a chart image and asks for analysis
* Asks about support/resistance levels, trend lines, or chart patterns
* Requests interpretation of technical indicators (RSI, MACD, Bollinger Bands, etc.)
* Asks about price action-based entry/exit signals

## Behaviour

When this skill is active:

1. **Trend identification:** determine the primary and secondary trend (higher highs/higher lows for uptrend; lower highs/lower lows for downtrend).
2. **Key levels:** identify significant support and resistance zones using swing pivots, round numbers, and volume nodes.
3. **Pattern recognition:** identify classic chart patterns (head and shoulders, double top/bottom, flags, wedges, triangles, cup-and-handle).
4. **Indicators:** interpret requested indicators in the context of price — do not rely on indicators in isolation.
5. **Caveat:** always pair technical analysis with a note that it is probabilistic, not deterministic.

## Common Indicators Quick Reference

| Indicator | Interpretation |
|---|---|
| RSI (14) | >70 overbought, <30 oversold; look for divergence |
| MACD | Signal line crossovers, histogram direction changes |
| Bollinger Bands | Price touching upper/lower band + reversal candle = mean reversion signal |
| 200-day SMA | Above = bullish bias; below = bearish bias |
| Volume | Trend move confirmed by expanding volume; reversal suspected on diverging volume |
| ADX | >25 = strong trend; <20 = weak/ranging market |

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Technical Analysis

When asked to analyse a chart or price action:
1. Identify the primary trend using higher highs/lower lows structure.
2. Mark key support and resistance levels.
3. Identify any chart patterns present and describe the implied measured move.
4. Interpret any indicators mentioned in context of price (not in isolation).
5. State the technical bias clearly (bullish/bearish/neutral) with the specific level that would invalidate it.
6. Always note that technical analysis is probabilistic — include: "Technical signals are not guarantees of future price movement."
```

## Notes and Limitations

* Technical analysis from chart images requires Gemini's multimodal (vision) capability — use `gemini-2.0-flash` or `gemini-2.5-pro`.
* The model cannot see live charts — users must provide image files or describe the chart.
* Combine with fundamental analysis for higher-conviction views.
