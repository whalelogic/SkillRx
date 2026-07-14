# Model Selection Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Data Science |
| **Task** | Select the appropriate ML algorithm for a problem |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `data-science/system-instructions/data-scientist.md` |
| **Output Format** | Structured Markdown recommendation |

---

## Prompt

Please help me select the most appropriate machine learning model(s) for the following problem.

### Problem Description

**{{PROBLEM_DESCRIPTION}}**

### Data Characteristics

* **Task type:** {{TASK_TYPE}} (binary classification / multiclass classification / regression / clustering / anomaly detection / time-series forecasting / ranking / NLP / computer vision / other)
* **Number of training samples:** {{N_SAMPLES}}
* **Number of features:** {{N_FEATURES}}
* **Feature types:** {{FEATURE_TYPES}} (numeric, categorical, text, image, time-series, mixed)
* **Target variable:** {{TARGET}} (e.g. "binary: churn yes/no", "continuous: revenue")
* **Class imbalance:** {{IMBALANCE}} (e.g. "5% positive class", "balanced", "unknown")
* **Missing data:** {{MISSING}} (e.g. "~15% missing across 3 columns", "none")

### Constraints and Requirements

* **Interpretability required:** {{INTERPRETABILITY}} (high — must explain to regulators / medium — team understanding / low — performance only)
* **Inference latency:** {{LATENCY}} (real-time < 100ms / near-real-time < 1s / batch OK)
* **Training compute:** {{COMPUTE}} (CPU only / single GPU / multi-GPU / cloud-scale)
* **Deployment environment:** {{DEPLOYMENT}} (mobile / browser / REST API / batch pipeline / embedded)
* **Primary success metric:** {{METRIC}} (e.g. AUC-ROC, F1, RMSE, precision@K)

### Task

1. Recommend 2–3 candidate algorithms with justification for each.
2. Produce a comparison table: Algorithm | Strengths | Weaknesses | Interpretability | Training Speed | Inference Speed | Typical Performance.
3. Recommend the **primary choice** with reasoning.
4. Describe the baseline model (simplest possible model to beat).
5. Describe the evaluation strategy (CV folds, train/val/test split, metrics, baseline comparison).
6. Provide a minimal Python skeleton for the recommended model using scikit-learn or the relevant framework.
