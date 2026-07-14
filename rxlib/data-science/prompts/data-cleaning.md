# Data Cleaning Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Data Science |
| **Task** | Identify and address data quality issues |
| **Recommended Model** | gemini-2.5-pro (with Code Execution) |
| **Recommended System Instruction** | `data-science/system-instructions/data-scientist.md` |
| **Output Format** | Python code + Markdown report |

---

## Prompt

Please identify and address data quality issues in the following dataset.

### Dataset

```
{{DATASET_SAMPLE_OR_SCHEMA}}
```

(Paste first 20–50 rows in CSV format, or paste the output of `df.info()`, `df.describe()`, and `df.head(20)`)

### Context

* **Dataset purpose:** {{PURPOSE}} (e.g. "training data for churn prediction model", "financial reporting")
* **Downstream task:** {{DOWNSTREAM}} (ML training / dashboard / statistical analysis / data product)
* **Volume:** {{N_ROWS}} rows × {{N_COLUMNS}} columns

### Data Cleaning Task

1. **Missing values** — identify columns with missing data; for each, recommend and implement the appropriate strategy:
   - Drop rows/columns if missingness > threshold
   - Impute with mean/median/mode for MCAR/MAR
   - Model-based imputation for MNAR
   - Flag as a separate indicator feature

2. **Duplicate rows** — detect and remove exact and near-duplicate rows.

3. **Outliers** — detect outliers using IQR and Z-score methods; decide whether to: keep, cap (winsorize), transform (log/sqrt), or remove.

4. **Data type corrections** — fix columns stored as wrong types (dates as strings, numbers as objects, booleans as integers).

5. **Inconsistent values** — standardise categorical values (e.g. "US", "USA", "United States" → "United States"), fix inconsistent capitalisation and whitespace.

6. **Invalid values** — detect and handle impossible values (negative ages, future birth dates, revenue < 0 where impossible).

7. **Encoding** — apply appropriate encoding for categorical variables intended for ML (one-hot, ordinal, target encoding).

### Output Requirements

* Complete, runnable pandas Python code that implements all cleaning steps.
* A data cleaning report table: Column | Issue | Count | Strategy | Rationale.
* Code must be idempotent — safe to run multiple times.
* Print the shape, missing value count, and dtype summary before and after cleaning.
