# Customer Lifetime Value (CLV) Model

## Purpose
A framework for calculating, tracking, and improving Customer Lifetime Value — the total revenue a customer generates over their relationship with you. CLV is the single most important metric for making smart acquisition, retention, and expansion decisions.

---

## Part 1: CLV Calculation Methods

### Method 1: Simple CLV (Best for Starting Out)

```
CLV = Average Revenue Per User (ARPU) × Average Customer Lifetime

Example:
  ARPU: $59/month
  Average lifetime: 14 months
  CLV: $59 × 14 = $826
```

### Method 2: Revenue-Based CLV (Better Accuracy)

```
CLV = (ARPU × Gross Margin %) × (1 / Churn Rate)

Example:
  ARPU: $59/month
  Gross margin: 80%
  Monthly churn rate: 5%
  CLV: ($59 × 0.80) × (1 / 0.05) = $47.20 × 20 = $944
```

### Method 3: Cohort-Based CLV (Best Accuracy)

Track actual revenue from a cohort of customers over time:

| Month | Cohort Size | Remaining | Revenue | Cumulative Revenue | Avg. CLV |
|:-----:|:----------:|:---------:|:-------:|:------------------:|:--------:|
| 0 | 100 | 100 | $5,900 | $5,900 | $59 |
| 1 | 100 | 92 | $5,428 | $11,328 | $113 |
| 2 | 100 | 85 | $5,015 | $16,343 | $163 |
| 3 | 100 | 78 | $4,602 | $20,945 | $209 |
| 6 | 100 | 60 | $3,540 | $35,000 | $350 |
| 12 | 100 | 40 | $2,360 | $55,000 | $550 |
| 24 | 100 | 22 | $1,298 | $75,000 | $750 |

### Method 4: Predictive CLV (Enterprise-Grade)

```
Predicted CLV = Σ (Monthly Revenue × Survival Probability × Discount Factor)
               for t = 1 to T months

Where:
  Survival Probability = probability customer is still active at month t
  Discount Factor = 1 / (1 + monthly discount rate)^t
```

---

## Part 2: CLV Segmentation

### CLV by Customer Segment

| Segment | ARPU | Avg. Lifetime | CLV | % of Customers | % of Revenue |
|---------|:----:|:------------:|:---:|:--------------:|:------------:|
| **Starter tier** | $/mo | months | $ | % | % |
| **Growth tier** | $/mo | months | $ | % | % |
| **Scale tier** | $/mo | months | $ | % | % |
| **Enterprise** | $/mo | months | $ | % | % |

### CLV by Acquisition Channel

| Channel | Avg. CLV | CAC | LTV:CAC | Payback (months) | Verdict |
|---------|:--------:|:---:|:-------:|:-----------------:|---------|
| Organic search | $ | $ | x | | |
| Paid search | $ | $ | x | | |
| Social (paid) | $ | $ | x | | |
| Referral | $ | $ | x | | |
| Events | $ | $ | x | | |
| Direct / brand | $ | $ | x | | |

### CLV by Cohort (Vintage Analysis)

| Signup Cohort | 3-Month CLV | 6-Month CLV | 12-Month CLV | Projected LTV |
|:------------:|:-----------:|:-----------:|:------------:|:-------------:|
| Jan 20XX | $ | $ | $ | $ |
| Feb 20XX | $ | $ | $ | $ |
| Mar 20XX | $ | $ | $ | $ |
| ... | | | | |

---

## Part 3: The CLV/CAC Relationship

### The Unit Economics Framework

```
LTV:CAC Ratio = Customer Lifetime Value / Customer Acquisition Cost

  >3:1  = Healthy — you can grow aggressively
  2-3:1 = Acceptable — but watch efficiency
  1-2:1 = Warning — acquisition is too expensive or retention is too low
  <1:1  = Unsustainable — you lose money on every customer
```

### CAC Payback Period

```
Payback Period (months) = CAC / (ARPU × Gross Margin %)

Example:
  CAC: $200
  ARPU: $59/month
  Gross margin: 80%
  Payback: $200 / ($59 × 0.80) = 4.2 months ← Healthy
```

| Payback Period | Assessment | Action |
|:--------------:|-----------|--------|
| <6 months | Excellent — grow aggressively | Increase marketing spend |
| 6-12 months | Good — healthy business | Optimize acquisition channels |
| 12-18 months | Acceptable (for enterprise) | Improve retention or reduce CAC |
| >18 months | Warning — cash flow strain | Fix retention or acquisition economics |

---

## Part 4: Improving CLV

### The 4 Levers of CLV

| Lever | Action | Impact |
|-------|--------|--------|
| **Reduce churn** | Better onboarding, product value, customer success | Extend lifetime = higher CLV |
| **Increase ARPU** | Upsells, cross-sells, price increases | More revenue per period |
| **Improve gross margin** | Reduce cost of delivery | More profit per $ of revenue |
| **Accelerate expansion** | Faster upsell, more seats, add-ons | Compound CLV growth |

### CLV Improvement Scenarios

| Scenario | Current | Improved | CLV Impact |
|----------|:-------:|:--------:|:----------:|
| Reduce churn from 5% to 3% | CLV: $944 | CLV: $1,573 | +67% |
| Increase ARPU from $59 to $79 | CLV: $944 | CLV: $1,264 | +34% |
| Both improvements | CLV: $944 | CLV: $2,107 | +123% |

---

## Part 5: CLV Dashboard

### Monthly CLV Tracking

| Metric | This Month | Last Month | 3-Month Avg | Target | Trend |
|--------|:----------:|:----------:|:-----------:|:------:|:-----:|
| **Average CLV (new customers)** | $ | $ | $ | $ | |
| **Average CLV (all customers)** | $ | $ | $ | $ | |
| **ARPU** | $ | $ | $ | $ | |
| **Monthly churn rate** | % | % | % | % | |
| **Expansion revenue** | $ | $ | $ | $ | |
| **CAC** | $ | $ | $ | $ | |
| **LTV:CAC ratio** | x | x | x | >3x | |
| **Payback period** | mo | mo | mo | <12 mo | |

### Cohort Retention Curves

Track what % of each signup cohort is still active:

| Month After Signup | Cohort A | Cohort B | Cohort C | Benchmark |
|:------------------:|:--------:|:--------:|:--------:|:---------:|
| 0 | 100% | 100% | 100% | 100% |
| 1 | % | % | % | 90% |
| 3 | % | % | % | 75% |
| 6 | % | % | % | 60% |
| 12 | % | % | % | 45% |
| 24 | % | % | % | 30% |

---

## Part 6: CLV Benchmarks by Industry

| Industry | Typical CLV Range | Key CLV Driver |
|----------|:-----------------:|---------------|
| **SaaS (SMB)** | $500-5,000 | Churn reduction |
| **SaaS (Mid-market)** | $5,000-50,000 | Expansion revenue |
| **SaaS (Enterprise)** | $50,000-500,000+ | Multi-year contracts |
| **E-Commerce** | $100-1,000 | Repeat purchase rate |
| **DTC Subscription** | $200-2,000 | Subscriber retention |
| **Professional Services** | $5,000-100,000+ | Relationship depth |
| **Agency** | $10,000-200,000+ | Retention + expansion |

---

## Part 7: Implementation Checklist

- [ ] Choose CLV calculation method appropriate to your data
- [ ] Calculate baseline CLV for all customers
- [ ] Segment CLV by tier, channel, and cohort
- [ ] Calculate LTV:CAC ratio and payback period
- [ ] Build cohort retention curves
- [ ] Identify the #1 lever for CLV improvement
- [ ] Set CLV targets for next quarter
- [ ] Establish monthly CLV tracking cadence
- [ ] Connect CLV data to acquisition decisions (stop spending on low-CLV channels)
- [ ] Share CLV insights with product, sales, and customer success

---

## Integration with Marketing36z OS

| System Component | How This Template Connects |
|-----------------|---------------------------|
| **Budget Template** | CLV/CAC ratio determines acceptable marketing spend per customer |
| **Attribution Guide** | Channel-level CLV data improves attribution-based budget decisions |
| **Lead Scoring** | High-CLV segments inform scoring weights |
| **Expansion Campaigns** | CLV analysis identifies expansion opportunity segments |
| **Pipeline Forecasting** | CLV feeds revenue forecasting models |
| **Dashboard Templates** | CLV metrics populate executive and board dashboards |
