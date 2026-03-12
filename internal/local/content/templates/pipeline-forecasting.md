# Pipeline Forecasting Model

## Purpose
A structured approach to predicting future revenue from your marketing and sales pipeline. Forecasting connects marketing activity to revenue outcomes and enables data-driven budget and resource decisions.

---

## Part 1: The Forecasting Framework

### Pipeline Stage Definitions

| Stage | Definition | Probability | Avg. Days in Stage |
|-------|-----------|:---:|:---:|
| **Lead** | Contact captured, not yet qualified | 5% | |
| **MQL** | Meets marketing qualification criteria | 10% | |
| **SQL** | Sales accepted, active pursuit | 20% | |
| **Discovery** | First meeting completed | 30% | |
| **Proposal** | Proposal/quote sent | 50% | |
| **Negotiation** | Terms being discussed | 75% | |
| **Verbal Commit** | Agreement in principle, awaiting signature | 90% | |
| **Closed Won** | Contract signed | 100% | |
| **Closed Lost** | Deal did not close | 0% | |

### The Weighted Pipeline Formula

```
Weighted Pipeline = Σ (Deal Value × Stage Probability)

Example:
  Lead:      10 deals × $5K × 5%  = $2,500
  SQL:        5 deals × $8K × 20% = $8,000
  Proposal:   3 deals × $10K × 50% = $15,000
  Negotiation: 2 deals × $12K × 75% = $18,000

  Total Weighted Pipeline = $43,500
```

---

## Part 2: Forecasting Models

### Model 1: Historical Conversion (Best for Predictable Businesses)

```
Forecasted Revenue = Current Pipeline × Historical Win Rate

Pipeline this month: $500K
Historical win rate: 25%
Forecasted revenue: $125K
```

### Model 2: Stage-Weighted (Best for B2B with Defined Sales Process)

| Stage | Deal Count | Total Value | Probability | Weighted Value |
|-------|:----------:|:-----------:|:---:|:-:|
| MQL | | $ | 10% | $ |
| SQL | | $ | 20% | $ |
| Discovery | | $ | 30% | $ |
| Proposal | | $ | 50% | $ |
| Negotiation | | $ | 75% | $ |
| Verbal Commit | | $ | 90% | $ |
| **Total** | | **$** | | **$** |

### Model 3: Cohort-Based (Best for SaaS/Subscription)

| Cohort | Leads In | Conv. Rate | Expected Customers | ARPU | Expected MRR |
|--------|:--------:|:----------:|:------------------:|:----:|:------------:|
| Month N (current) | | % | | $ | $ |
| Month N+1 | | % | | $ | $ |
| Month N+2 | | % | | $ | $ |
| **Total** | | | | | **$** |

### Model 4: Top-Down (Best for Board/Investor Reporting)

```
Revenue Target: $___K
Avg. Deal Size: $___K
Deals Needed: Target / Avg. Deal Size = ___
Win Rate: ___%
Opportunities Needed: Deals / Win Rate = ___
SQL → Opp Rate: ___%
SQLs Needed: Opportunities / Conversion = ___
MQL → SQL Rate: ___%
MQLs Needed: SQLs / Conversion = ___
Lead → MQL Rate: ___%
Leads Needed: MQLs / Conversion = ___

Marketing must generate ___ leads/month to hit target.
```

---

## Part 3: The Monthly Forecast Report

### Pipeline Snapshot

| Metric | This Month | Last Month | MoM Δ | vs. Target |
|--------|:----------:|:----------:|:-----:|:----------:|
| **Total pipeline value** | $ | $ | % | % |
| **Weighted pipeline** | $ | $ | % | % |
| **Average deal size** | $ | $ | % | — |
| **Number of open opportunities** | | | % | — |
| **Pipeline coverage ratio** | x | x | — | >3x target |
| **Average sales cycle** | days | days | % | — |
| **Win rate** | % | % | — | — |

### Pipeline Coverage Analysis

```
Pipeline Coverage = Weighted Pipeline / Revenue Target

Example:
  Weighted pipeline: $375K
  Quarterly target: $125K
  Coverage: 3.0x ← Healthy (target: 3x minimum)
```

| Coverage | Interpretation | Action |
|:--------:|---------------|--------|
| <2x | Pipeline gap — at risk of missing target | Accelerate demand gen |
| 2-3x | Marginal — tight but achievable | Increase close rates and acceleration |
| 3-4x | Healthy — good pipeline balance | Maintain current efforts |
| >4x | Strong — may have pipeline quality issues | Focus on conversion, not volume |

### Revenue Forecast by Month

| Month | Commit (90%+) | Best Case (50%+) | Upside (20%+) | Target |
|-------|:---:|:---:|:---:|:---:|
| Current month | $ | $ | $ | $ |
| Next month | $ | $ | $ | $ |
| Month +2 | $ | $ | $ | $ |
| **Quarter Total** | **$** | **$** | **$** | **$** |

---

## Part 4: Forecast Accuracy Tracking

### Monthly Forecast vs. Actual

| Month | Forecast | Actual | Variance | Accuracy |
|-------|:--------:|:------:|:--------:|:--------:|
| Jan | $ | $ | $ | % |
| Feb | $ | $ | $ | % |
| Mar | $ | $ | $ | % |
| ... | | | | |

**Forecast accuracy:** `1 - (|Forecast - Actual| / Actual) × 100`

| Accuracy | Interpretation | Action |
|:--------:|---------------|--------|
| >90% | Excellent — model is well-calibrated | Maintain |
| 80-90% | Good — minor adjustments needed | Review stage probabilities |
| 70-80% | Fair — model needs tuning | Recalibrate conversion rates quarterly |
| <70% | Poor — model unreliable | Rebuild with fresh data |

---

## Part 5: Improving Forecast Accuracy

### Common Forecast Errors

| Error | Cause | Fix |
|-------|-------|-----|
| **Perpetual optimism** | Deals stay in pipeline without progressing | Enforce stage aging rules (move or close) |
| **Sandbagging** | Sales under-reports to beat forecast | Separate commit from stretch forecasts |
| **Stale pipeline** | Deals sitting for months without activity | Auto-close deals with no activity in 30+ days |
| **Wrong stage assignment** | Deals in "Proposal" that haven't actually sent one | Clear, verifiable criteria per stage |
| **Missing pipeline** | Deals not entered until late | CRM hygiene SLA: enter within 24 hours |

### Stage Verification Criteria

| Stage | Verification | Evidence |
|-------|-------------|---------|
| **SQL** | Sales has spoken to the contact | Call log or email reply |
| **Discovery** | Discovery call completed, needs identified | Meeting notes |
| **Proposal** | Proposal document sent | Sent email with attachment |
| **Negotiation** | Prospect responded to proposal | Reply email or meeting |
| **Verbal Commit** | Decision maker said yes | Written confirmation |

---

## Part 6: Forecasting Cadence

| Activity | Frequency | Owner |
|----------|:---------:|-------|
| Update pipeline data | Daily | Sales reps |
| Weekly forecast review | Weekly | Sales + Marketing |
| Monthly forecast report | Monthly | Marketing Ops / Rev Ops |
| Quarterly model recalibration | Quarterly | Marketing Ops |
| Annual model rebuild | Annually | Marketing Ops + Finance |

---

## Integration with Marketing36z OS

| System Component | How This Template Connects |
|-----------------|---------------------------|
| **Dashboard Templates** | Pipeline data populates executive and board dashboards |
| **Budget Template** | Forecast informs marketing budget allocation |
| **OKR Framework** | Revenue OKRs tracked through forecast actuals |
| **Lead Scoring** | Lead quality directly impacts conversion rates in forecast |
| **Attribution Guide** | Source-level data improves forecast precision by channel |
