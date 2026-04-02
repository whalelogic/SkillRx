# Exploratory Data Analysis Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Data Science |
| **Task** | Systematic exploratory data analysis |
| **Recommended Model** | gemini-2.5-pro (with Code Execution enabled) |
| **Recommended System Instruction** | `data-science/system-instructions/data-scientist.md` |
| **Output Format** | Python code + Markdown analysis |

---

## Prompt

Please perform a systematic exploratory data analysis (EDA) on the following dataset.

### Dataset

Paste a sample of the data (first 20–50 rows in CSV format or describe the schema):

```
{{DATASET_SAMPLE}}
```

Or describe the dataset:
```
{{DATASET_DESCRIPTION}}
```

### Context

* **Business question / analytical goal:** {{BUSINESS_QUESTION}}
* **Target variable (if any):** {{TARGET_VARIABLE}}
* **Number of rows:** {{N_ROWS}} (approximate)
* **Data source:** {{DATA_SOURCE}} (e.g. PostgreSQL, CSV, API, BigQuery)

### EDA Task

1. **Data inventory** — shape, column names, dtypes, memory usage.
2. **Missingness analysis** — missing value counts and percentages per column; pattern analysis (MCAR, MAR, MNAR assessment).
3. **Univariate analysis** — for each column: distribution summary (describe()), histograms for numeric, value counts for categorical; flag skewness > 2 and kurtosis > 7.
4. **Bivariate analysis** — correlation matrix for numeric features; relationship between key features and the target variable.
5. **Outlier detection** — IQR method and Z-score method; flag extreme outliers.
6. **Data quality issues** — inconsistent formats, suspicious values, duplicate rows, cardinality issues.
7. **Key insights** — summarise the top 5 findings that should inform modelling or business decisions.
8. **Recommended next steps** — data cleaning steps, feature engineering ideas, modelling approach.

### Output Requirements

* Produce complete, runnable Python code using pandas and matplotlib/seaborn.
* Follow each code block with a markdown interpretation of what the output means.
* Flag any data quality issues prominently.
