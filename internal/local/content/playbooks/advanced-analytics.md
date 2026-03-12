# Playbook: Advanced Analytics

## Sources
- Avinash Kaushik — *Web Analytics 2.0*, Analytics Maturity
- Reforge — Growth Analytics
- Marketing Science literature
- Marketing36z OS synthesis

---

## Part 1: Analytics Maturity Model

| Level | Name | Capabilities | Tools |
|:-----:|------|-------------|-------|
| 1 | **Descriptive** | "What happened?" — dashboards, reports | GA4, spreadsheets |
| 2 | **Diagnostic** | "Why did it happen?" — segmentation, funnel analysis | Mixpanel, Amplitude |
| 3 | **Predictive** | "What will happen?" — forecasting, modeling | Python/R, BigQuery |
| 4 | **Prescriptive** | "What should we do?" — optimization, automation | ML pipelines, MMM |
| 5 | **Autonomous** | "System decides and acts" — automated optimization | AI-driven platforms |

---

## Part 2: Predictive Analytics for Marketing

### Customer Churn Prediction

| Input Feature | Data Source | Why It Predicts Churn |
|--------------|-----------|---------------------|
| Login frequency trend | Product analytics | Declining usage = disengagement |
| Feature adoption depth | Product analytics | Low adoption = low perceived value |
| Support ticket sentiment | Support platform | Negative sentiment = frustration |
| Days since last activity | Product analytics | Inactivity = at risk |
| Payment failures | Billing system | Involuntary churn precursor |
| NPS score | Survey | Low NPS = likely to churn |
| Engagement with emails | ESP | Not opening = disengaged |
| Contract renewal date proximity | CRM | Approaching renewal without engagement = risk |

### Lead-to-Customer Propensity

| Input Feature | What It Predicts |
|--------------|-----------------|
| Source / channel | Some sources produce higher-quality leads |
| Engagement score (pages viewed, content downloaded) | Higher engagement = higher conversion likelihood |
| Company fit score (firmographic match to ICP) | Better fit = higher conversion |
| Time to first response | Faster response = higher conversion |
| Number of stakeholders engaged | Multi-threading = serious buying intent |
| Pricing page visits | Strong purchase intent signal |

### Revenue Forecasting

| Method | Accuracy | Complexity | When to Use |
|--------|:--------:|:----------:|-----------|
| **Moving average** | Low-Medium | Low | Quick directional forecast |
| **Linear regression** | Medium | Low | Steady growth businesses |
| **Cohort-based projection** | High | Medium | SaaS with cohort data |
| **Time series (ARIMA/Prophet)** | High | High | Seasonal or complex patterns |
| **ML ensemble** | Very High | Very High | Large datasets, multiple variables |

---

## Part 3: Customer Journey Analytics

### Journey Mapping with Data

| Journey Stage | Data Points | Key Question | Action |
|-------------|------------|-------------|--------|
| **First touch** | Source, landing page, device | How are they finding us? | Optimize acquisition channels |
| **Engagement** | Pages viewed, content consumed, time on site | What interests them? | Personalize content |
| **Conversion** | Form fills, signups, purchases | What triggers the decision? | Optimize conversion paths |
| **Activation** | First value action, feature adoption | Are they getting value? | Improve onboarding |
| **Retention** | Login frequency, feature usage, engagement | Are they staying? | Reduce churn |
| **Expansion** | Upgrade triggers, add-on usage | Are they growing? | Identify expansion opportunities |
| **Advocacy** | NPS, referrals, reviews | Are they promoting us? | Activate advocacy program |

### Path Analysis

Track the most common sequences that lead to conversion:

| Rank | Path | Conv. Rate | Volume | Revenue |
|:----:|------|:---:|:---:|:---:|
| 1 | Blog → Pricing → Signup | % | | $ |
| 2 | LinkedIn → Homepage → Pricing → Signup | % | | $ |
| 3 | Google → Blog → Blog → Pricing → Signup | % | | $ |
| 4 | Referral → Homepage → Signup | % | | $ |
| 5 | Email → Blog → Pricing → Demo → Close | % | | $ |

---

## Part 4: Experimentation Analytics

### Experiment Analysis Framework

| Metric | Control | Variant | Lift | Confidence | Significance |
|--------|:-------:|:-------:|:----:|:----------:|:------------:|
| Conversion rate | % | % | % | % | Yes / No |
| Revenue per visitor | $ | $ | % | % | Yes / No |
| Average order value | $ | $ | % | % | Yes / No |

### Statistical Significance Quick Reference

| Sample Size Per Variant | Minimum Detectable Effect | Confidence Level |
|:-:|:-:|:-:|
| 100 | >30% lift | 80% |
| 500 | >10% lift | 90% |
| 1,000 | >5% lift | 95% |
| 5,000 | >2% lift | 95% |
| 10,000 | >1% lift | 99% |

### Multi-Armed Bandit vs. A/B Test

| Method | Best For | How It Works |
|--------|---------|-------------|
| **A/B test** | Definitive learning | Equal traffic split, run to significance |
| **Multi-armed bandit** | Revenue optimization during test | Dynamically routes more traffic to winning variant |

**Rule:** Use A/B tests when you need to learn. Use bandits when you need to earn while learning.

---

## Part 5: Cohort Analysis

### Cohort Analysis Template

| Signup Month | Month 0 | Month 1 | Month 2 | Month 3 | Month 6 | Month 12 |
|:----------:|:-------:|:-------:|:-------:|:-------:|:-------:|:--------:|
| Jan | 100% | % | % | % | % | % |
| Feb | 100% | % | % | % | % | |
| Mar | 100% | % | % | % | | |
| Apr | 100% | % | % | | | |

### What Cohort Analysis Reveals

| Pattern | What It Means | Action |
|---------|-------------|--------|
| Newer cohorts retain better | Product / onboarding is improving | Keep optimizing |
| Newer cohorts retain worse | Something degraded (quality, ICP drift) | Investigate changes |
| Sharp drop at Month 1 | Onboarding problem | Fix first-week experience |
| Gradual decline | Normal attrition, manageable | Focus on engagement at risk points |
| Flat retention after Month 6 | Found the "sticky" point | Get users to Month 6 faster |

---

## Part 6: Revenue Analytics

### Unit Economics Dashboard

| Metric | Current | Last Quarter | Target | Trend |
|--------|:-------:|:----------:|:------:|:-----:|
| MRR / ARR | $ | $ | $ | |
| New MRR | $ | $ | $ | |
| Expansion MRR | $ | $ | $ | |
| Churned MRR | $ | $ | $ | |
| Net New MRR | $ | $ | $ | |
| ARPU | $ | $ | $ | |
| Gross Margin | % | % | % | |
| CAC | $ | $ | $ | |
| LTV | $ | $ | $ | |
| LTV:CAC | x | x | >3x | |
| Payback Period | mo | mo | <12 | |
| Quick Ratio | x | x | >4x | |

### The SaaS Quick Ratio

```
Quick Ratio = (New MRR + Expansion MRR) / (Churned MRR + Contraction MRR)

>4x = Very healthy growth
2-4x = Good
1-2x = Leaky bucket
<1x = Shrinking
```

---

## Part 7: Analytics Stack by Company Stage

| Stage | Analytics Needs | Recommended Stack |
|-------|----------------|------------------|
| **Pre-revenue** | Basic traffic + conversion tracking | GA4 + Spreadsheet |
| **$0-1M ARR** | Funnel analysis + cohorts | GA4 + Mixpanel/Amplitude (free tier) |
| **$1M-10M ARR** | Multi-touch, experiments, forecasting | Segment + Amplitude + Looker Studio |
| **$10M+ ARR** | Full data warehouse, ML, MMM | Segment + BigQuery + Looker + custom models |
| **Enterprise** | Real-time, cross-platform, predictive | Full CDP + data warehouse + ML pipeline |

---

## Part 8: Analytics Implementation Checklist

- [ ] Web analytics configured with proper event tracking (GA4)
- [ ] Product analytics tracking key user actions
- [ ] CRM data clean and connected to marketing analytics
- [ ] Funnel stages defined with conversion tracking between each
- [ ] Cohort analysis running monthly
- [ ] Experiment infrastructure in place (feature flags or A/B tool)
- [ ] Revenue metrics dashboard live (MRR, churn, LTV, CAC)
- [ ] Attribution model implemented (multi-touch + self-reported)
- [ ] Automated reporting for key stakeholders
- [ ] Regular analytics review cadence established

---

## Integration with Marketing36z OS

| System Component | How This Playbook Connects |
|-----------------|---------------------------|
| **Attribution Guide** | Attribution is the foundation layer of marketing analytics |
| **CLV Model** | CLV calculations feed predictive analytics |
| **Pipeline Forecasting** | Advanced forecasting models improve pipeline accuracy |
| **MMM** | Marketing mix modeling is the advanced analytics version of budget optimization |
| **Experiment Framework** | Experiment analytics ensure statistical rigor |
| **Dashboard Templates** | Analytics outputs populate all dashboard levels |
