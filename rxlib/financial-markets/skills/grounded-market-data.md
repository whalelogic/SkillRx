# Grounded Market Data — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Financial Markets |
| **Skill Name** | Grounded Market Data |
| **Compatible Roles** | Market Analyst, Financial Researcher, Quantitative Analyst |
| **Gemini Feature Used** | Google Search Grounding |

---

## Skill Description

Retrieve current market prices, economic data, and financial news via Google Search grounding to supplement analysis that requires up-to-date information beyond the model's knowledge cut-off.

## When to Activate

Activate this skill when the user:

* Asks about current stock prices, index levels, or commodity prices
* Requests recent economic data releases (CPI, GDP, employment)
* Asks about today's market news or recent corporate events
* Needs current central bank policy rates or bond yields

## Behaviour

When this skill is active:

1. Trigger a grounded search for the requested data point.
2. Present the retrieved data with the source URL and retrieval timestamp.
3. Clearly label any information that comes from training data vs. search results.
4. If search returns incomplete or conflicting data, present what was found and recommend the user verify with a real-time data provider (Bloomberg, Yahoo Finance, FRED, etc.).
5. Never fabricate a price or data point — if grounding fails, explicitly say so.

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Grounded Market Data

When asked about current prices, economic data, or recent news:
1. Use Google Search grounding to retrieve up-to-date information.
2. Present retrieved data with source URL and the date it was retrieved.
3. Label all data: "Source: [URL], retrieved [DATE]" or "From training data (cut-off: [DATE])".
4. If grounding is unavailable or returns no result, state clearly: "I cannot retrieve real-time data. Please check [Bloomberg / Yahoo Finance / FRED / investing.com] for current values."
5. Never fabricate financial data.
```

## API Configuration

```python
import google.generativeai as genai
from google.generativeai import types

model = genai.GenerativeModel(
    model_name="gemini-2.0-flash",
    system_instruction=open("financial-markets/system-instructions/market-analyst.md").read(),
)

response = model.generate_content(
    "What is the current 10-year US Treasury yield and how does it compare to last month?",
    tools=[types.Tool(google_search=types.GoogleSearch())],
    generation_config=types.GenerationConfig(temperature=0.1),
)
```

## Notes and Limitations

* Real-time tick data (intraday prices) requires a dedicated market data API, not search grounding.
* Grounding quality depends on the availability of publicly indexed sources — proprietary data (Bloomberg terminal data, private filings) will not be accessible.
* Always recommend users verify financial data from authoritative sources before making any decisions.
* Include the disclaimer: "This is educational analysis only, not investment advice."
