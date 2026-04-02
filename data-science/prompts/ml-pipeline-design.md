# ML Pipeline Design Prompt

## Metadata

| Field | Value |
|---|---|
| **Domain** | Data Science |
| **Task** | Design an end-to-end ML pipeline |
| **Recommended Model** | gemini-2.5-pro |
| **Recommended System Instruction** | `data-science/system-instructions/ml-engineer.md` |
| **Output Format** | Architecture design + Mermaid diagram + code skeleton |

---

## Prompt

Please design an end-to-end machine learning pipeline for the following use case.

### Use Case

**{{USE_CASE_DESCRIPTION}}**

### Requirements

* **ML task:** {{ML_TASK}} (classification / regression / ranking / NLP / computer vision / anomaly detection)
* **Data source:** {{DATA_SOURCE}} (e.g. PostgreSQL, S3 CSV, real-time Kafka stream, BigQuery)
* **Training frequency:** {{TRAINING_FREQUENCY}} (one-time / daily / weekly / triggered by drift)
* **Inference mode:** {{INFERENCE_MODE}} (real-time REST API / batch / streaming / edge)
* **Scale:** {{SCALE}} (e.g. "10K predictions/day", "1M rows training data", "50ms p99 latency")
* **Infrastructure:** {{INFRASTRUCTURE}} (e.g. Google Cloud Vertex AI, AWS SageMaker, Azure ML, on-premise Kubernetes, local)
* **Team:** {{TEAM}} (1 data scientist / small team / MLOps team available)
* **Maturity:** {{MATURITY}} (MVP / production / enterprise)

### Pipeline Design Task

Design the following pipeline stages:

1. **Data ingestion** — how is raw data collected and stored?
2. **Data validation** — what checks run before training (schema validation, distribution checks)?
3. **Feature engineering** — what transformations are applied? How are features versioned and served online/offline?
4. **Model training** — algorithm, hyperparameter tuning strategy, experiment tracking.
5. **Model evaluation** — evaluation criteria, human review gate, model registry.
6. **Model deployment** — serving infrastructure, API contract, rollout strategy (canary, blue/green).
7. **Monitoring** — what signals indicate degradation? What triggers retraining?
8. **Feedback loop** — how does production data flow back to improve the model?

### Output Requirements

* A Mermaid diagram showing the complete pipeline.
* A component table: Stage | Technology | Input | Output | Owner.
* A Python code skeleton for the training stage using the recommended framework.
* A list of the top 3 risks and mitigations for this pipeline design.
