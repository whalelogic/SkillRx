# Code Execution for Data Analysis — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Data Science |
| **Skill Name** | Code Execution for Data Analysis |
| **Compatible Roles** | Data Scientist, Data Analyst |
| **Gemini Feature Used** | Code Execution tool |

---

## Skill Description

Run Python data analysis code (pandas, NumPy, matplotlib, scipy, scikit-learn) in Gemini's sandboxed code execution environment to validate results, compute statistics, and demonstrate findings with real output.

## When to Activate

Activate this skill when the user:

* Provides a dataset and asks for analysis or statistics
* Asks for a calculation that requires running code to be accurate
* Wants to validate that a data processing snippet produces the correct output
* Asks for a visualisation (chart code can be executed to confirm it runs)

## Behaviour

When this skill is active:

1. Write a minimal, self-contained Python snippet to answer the analytical question.
2. Execute it using the Code Execution tool.
3. Include the actual output in the response (not simulated).
4. If execution raises an error, debug and re-run before presenting the answer.
5. Annotate the output with a plain-language interpretation.

## Available Libraries in Gemini Sandbox

| Library | Use |
|---|---|
| `pandas` | DataFrame manipulation, groupby, pivoting |
| `numpy` | Numerical computation, array operations |
| `matplotlib` / `seaborn` | Static visualisations |
| `scipy` | Statistical tests, optimisation |
| `scikit-learn` | ML models, preprocessing, metrics |
| `statistics` | Basic stats (mean, median, stdev) |

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Code Execution for Data Analysis

When answering data analysis questions that require computation:
1. Write a minimal self-contained Python snippet using pandas/numpy/scipy.
2. Execute it using the Code Execution tool.
3. Include the real execution output — never simulate output.
4. If the code raises an error, fix and re-run.
5. Follow the output with a plain-language interpretation.
6. Label executed code: # Executed — output shown below
```

## API Configuration

```python
import google.generativeai as genai

model = genai.GenerativeModel(
    model_name="gemini-2.5-pro",
    tools=[{"code_execution": {}}],
    system_instruction=open("data-science/system-instructions/data-scientist.md").read(),
)
```

## Notes and Limitations

* The sandbox has no persistent state — data must be defined within each code block.
* External network calls and file system access are not available.
* For large datasets (>10MB), paste a representative sample instead of the full dataset.
* Matplotlib `plt.show()` generates inline images in some interfaces — if not supported, save to buffer and show dimensions.
