# Model Evaluation Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Data Science |
| **Task** | Evaluate and interpret ML model performance |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `data-science/system-instructions/data-scientist.md` |
| **Output Format** | Structured Markdown evaluation report |

---

## Prompt

Please evaluate and interpret the following machine learning model results.

### Model Information

* **Task type:** {{TASK_TYPE}} (binary classification / multiclass / regression / etc.)
* **Algorithm:** {{ALGORITHM}}
* **Target variable:** {{TARGET}}
* **Dataset split:** {{SPLIT}} (e.g. "70/15/15 train/val/test", "5-fold CV")

### Results

Paste the evaluation output:

```
{{EVALUATION_RESULTS}}
```

Include as many of the following as available:
* Classification: confusion matrix, classification report (precision/recall/F1 per class), ROC AUC, PR AUC
* Regression: RMSE, MAE, MAPE, R², residual plot description
* Cross-validation scores (mean ± std)
* Training vs. validation curve data
* Feature importance values
* Calibration curve data (classification)

### Business Context

* **Business metric:** {{BUSINESS_METRIC}} (e.g. "each false negative costs $500; each false positive costs $20")
* **Acceptable error rate:** {{ACCEPTABLE_ERROR}}
* **Baseline to beat:** {{BASELINE}} (e.g. "random guess = 50% accuracy", "current rule-based system = 65% precision")

### Evaluation Task

1. **Metric interpretation** — explain what each reported metric means in plain language.
2. **Bias-variance diagnosis** — is the model underfitting, overfitting, or well-fitted? What is the evidence?
3. **Class-level analysis** (classification) — which classes perform best/worst and why?
4. **Business impact assessment** — translate model performance into business outcomes using the cost matrix provided.
5. **Comparison to baseline** — is the model worth deploying over the current solution?
6. **Weaknesses and failure modes** — where should the model not be trusted?
7. **Improvement recommendations** — what are the highest-leverage next steps to improve performance?
