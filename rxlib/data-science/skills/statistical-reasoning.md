# Statistical Reasoning — Skill

## Metadata

| Field | Value |
|---|---|
| **Domain** | Data Science |
| **Skill Name** | Statistical Reasoning |
| **Compatible Roles** | Data Scientist, Data Analyst, Research Assistant |
| **Gemini Feature Used** | Standard generation; Code Execution for calculations |

---

## Skill Description

Apply rigorous statistical thinking to analytical questions — selecting appropriate tests, checking assumptions, interpreting results correctly, and communicating uncertainty honestly.

## Core Statistical Concepts Quick Reference

### Choosing a Statistical Test

| Question | Data Type | Test |
|---|---|---|
| Compare 2 group means | Continuous, normal | Independent t-test |
| Compare 2 group means | Continuous, non-normal | Mann-Whitney U |
| Compare 3+ group means | Continuous, normal | One-way ANOVA |
| Compare 3+ group means | Continuous, non-normal | Kruskal-Wallis |
| Association between 2 categorical | Categorical | Chi-square / Fisher's exact |
| Correlation between 2 continuous | Continuous | Pearson (normal) / Spearman (non-normal) |
| Before/after same group | Continuous, normal | Paired t-test |
| Before/after same group | Continuous, non-normal | Wilcoxon signed-rank |
| A/B test for proportions | Binary | Z-test for proportions |
| Time to event | Censored continuous | Kaplan-Meier, log-rank, Cox PH |

### Key Statistical Concepts

| Concept | Correct Interpretation |
|---|---|
| p-value | Probability of observing data at least this extreme *if the null hypothesis is true*. Not the probability the null is true. |
| Confidence interval | If we repeated the experiment many times, X% of CIs would contain the true value. |
| Effect size | Practical significance. Always report alongside p-value. |
| Power | Probability of detecting a real effect of a given size. |
| Multiple comparisons | Each additional test inflates Type I error — apply Bonferroni or FDR correction. |

## Behaviour

When this skill is active:

1. For any analytical question involving data, identify the relevant statistical approach before computing.
2. Check and state the test assumptions (normality, independence, homoscedasticity).
3. Report results as: test statistic, p-value, effect size, confidence interval, and plain-language conclusion.
4. Flag when sample size is too small for reliable inference.
5. Explicitly distinguish between statistical significance and practical significance.

## Integration Snippet

Append the following block to the system instruction of any agent that should have this skill:

```
## Skill: Statistical Reasoning

For all analytical tasks involving data:
1. Identify the appropriate statistical test and justify the choice.
2. State the assumptions and check whether they are met.
3. Report: test statistic, p-value, effect size (Cohen's d / Cramér's V / r), and 95% CI.
4. Write the conclusion in plain language: "There is [strong/weak/no] evidence that..."
5. Distinguish between statistical significance (p-value) and practical significance (effect size).
6. Apply multiple comparison corrections when running more than one test.
```

## Notes and Limitations

* Statistical tests require adequate sample size — always perform a power analysis before data collection.
* Violations of test assumptions can invalidate results — prefer non-parametric alternatives when assumptions are uncertain.
* Statistical analysis can establish association but not causation in observational data — state this clearly.
