# Playbook: Marketing Mix Modeling (MMM)

## Sources
- Google / Meta — Open-Source MMM (Meridian, Robyn)
- Les Binet & Peter Field — *The Long and the Short of It*
- Byron Sharp — *How Brands Grow*
- Marketing36z OS synthesis

---

## Part 1: What Marketing Mix Modeling Is

### MMM vs. Attribution

| Dimension | Attribution (MTA) | Marketing Mix Modeling (MMM) |
|-----------|------------------|----------------------------|
| **What it measures** | Individual user journeys | Aggregate channel impact on business outcomes |
| **Data source** | Cookies, pixels, UTMs | Historical spending + revenue data |
| **Privacy dependency** | High (needs tracking) | None (uses aggregate data) |
| **Scope** | Digital channels only | All channels including offline (TV, radio, events, PR) |
| **Time horizon** | Real-time / recent | Months to years of historical data |
| **Brand impact** | Usually invisible | Can model brand awareness effects |
| **Best for** | Tactical optimization | Strategic budget allocation |
| **Minimum data** | Immediate | 2+ years of weekly/monthly data |

### When to Use MMM

| Situation | MMM Useful? | Why |
|-----------|:---:|-----|
| <$50K/mo marketing spend | Not yet | Not enough data or channel diversity |
| $50K-200K/mo, 3+ channels | Consider | Enough data to start modeling |
| $200K+/mo, 5+ channels | Yes | Statistical significance, real budget allocation impact |
| Offline marketing (events, PR, partnerships) | Yes | Only way to measure offline impact |
| Privacy/cookie concerns (post iOS 14.5) | Yes | No user-level tracking required |

---

## Part 2: The MMM Process

### Step 1: Data Collection

| Data Type | Source | Granularity | Minimum History |
|-----------|--------|:-----------:|:-:|
| **Revenue / conversions** | CRM, payment processor | Weekly or monthly | 2+ years |
| **Marketing spend by channel** | Ad platforms, finance | Weekly or monthly | 2+ years |
| **Pricing data** | Product / finance | Monthly | 2+ years |
| **External factors** | Various | Weekly or monthly | 2+ years |
| **Seasonality indicators** | Calendar | Weekly | Auto-generated |
| **Competitive data** (optional) | SimilarWeb, SEMrush | Monthly | 1+ year |

### External Factors to Include

| Factor | Why It Matters | Source |
|--------|---------------|--------|
| **Seasonality** | Demand varies by time of year | Calendar, historical data |
| **Economic conditions** | Consumer confidence affects spending | Federal Reserve, BLS data |
| **Industry events** | Conferences, regulations, crises | Industry calendar |
| **Competitor activity** | Major launches, price changes, campaigns | Competitive monitoring |
| **Weather** (if relevant) | Affects certain industries | Weather API |
| **PR / earned media** | Brand awareness spikes | Media monitoring |
| **Product changes** | New features, pricing changes | Internal log |

### Step 2: Model Building

| Component | What It Does |
|-----------|-------------|
| **Base revenue** | Revenue you'd get with zero marketing (organic demand) |
| **Channel contribution** | Additional revenue driven by each marketing channel |
| **Adstock / carryover** | Delayed effect of marketing (ads you saw last week still influence you) |
| **Saturation curves** | Diminishing returns as you spend more on a channel |
| **External factors** | Impact of non-marketing variables |

### Conceptual Model

```
Revenue = Base + Σ(Channel_i × Coefficient_i × Adstock_i × Saturation_i) + External_Factors + Error
```

### Step 3: Model Outputs

| Output | What It Tells You | Decision It Drives |
|--------|------------------|-------------------|
| **Channel ROI** | Return on investment per channel | Which channels to invest more/less in |
| **Saturation point** | Where additional spend yields diminishing returns | Maximum efficient spend per channel |
| **Optimal budget split** | Ideal allocation across channels | Budget rebalancing |
| **Base vs. incremental** | How much revenue is organic vs. marketing-driven | True marketing ROI |
| **Carryover effects** | How long marketing impact lasts after spend stops | Timing and pacing decisions |

---

## Part 3: Interpreting MMM Results

### Channel Contribution Analysis

| Channel | Monthly Spend | Attributed Revenue | ROI | Saturation % | Recommendation |
|---------|:---:|:---:|:---:|:---:|---|
| Paid Search | $ | $ | x | % | |
| Social Ads | $ | $ | x | % | |
| Content / SEO | $ | $ | x | % | |
| Email | $ | $ | x | % | |
| Events | $ | $ | x | % | |
| PR / Earned | $ | $ | x | % | |
| **Total** | **$** | **$** | **x** | | |

### Budget Optimization Scenario

| Scenario | Channel A | Channel B | Channel C | Total Spend | Projected Revenue | Δ vs. Current |
|----------|:---------:|:---------:|:---------:|:-----------:|:-----------------:|:-------------:|
| **Current** | $ | $ | $ | $ | $ | — |
| **Optimized** | $ | $ | $ | $ | $ | +$__ (+__%) |
| **Aggressive** | $ | $ | $ | $ | $ | +$__ (+__%) |
| **Conservative** | $ | $ | $ | $ | $ | +$__ (+__%) |

---

## Part 4: Practical MMM Without Data Science

### The Simplified MMM Approach (Spreadsheet-Based)

For companies without a data science team, a simplified approach:

| Step | Method |
|------|--------|
| 1. Collect monthly spend and revenue by channel (24+ months) | Spreadsheet |
| 2. Calculate correlation between each channel's spend and revenue | =CORREL() in Excel |
| 3. Run basic regression (revenue as dependent, channels as independent) | Excel Data Analysis or Google Sheets LINEST |
| 4. Identify top-contributing channels | Coefficient strength |
| 5. Test budget shifts incrementally (10-20% reallocation) | A/B geographic or time-based test |

### DIY Channel Contribution Estimation

| Method | How | Accuracy |
|--------|-----|:---:|
| **Spend-revenue correlation** | Plot monthly spend vs. revenue per channel | Low-Medium |
| **Holdout test** | Turn off a channel for 4 weeks, measure revenue impact | Medium-High |
| **Geographic test** | Different spend levels in different regions | High |
| **Time-based test** | Increase/decrease spend in alternating months | Medium |
| **Incrementality test** | Exposed vs. holdout audience | High |

---

## Part 5: Long-Term vs. Short-Term Effects

### The Binet & Field Framework

| Dimension | Brand Building | Performance / Activation |
|-----------|:---:|:---:|
| **Objective** | Build mental availability | Trigger immediate purchase |
| **Timeframe** | Months to years | Days to weeks |
| **Channels** | Content, PR, brand ads, events | Paid search, retargeting, email, promotions |
| **Measurement** | Brand awareness, SOV, consideration | Leads, conversions, ROAS |
| **MMM visibility** | Difficult (long lag) | Easy (short lag) |
| **Optimal budget split** | 60% of budget | 40% of budget |

### Brand vs. Performance Budget Split by Business Stage

| Stage | Brand | Performance | Rationale |
|-------|:-----:|:----------:|-----------|
| **Pre-revenue** | 20% | 80% | Need immediate traction |
| **$0-1M ARR** | 30% | 70% | Building brand while acquiring |
| **$1M-10M ARR** | 40% | 60% | Brand starts compounding |
| **$10M+ ARR** | 60% | 40% | Brand is the moat |

---

## Part 6: MMM Implementation Checklist

### Getting Started

- [ ] Collect 24+ months of channel-level spend data
- [ ] Collect matching revenue/conversion data at the same granularity
- [ ] Document external factors (seasonality, promotions, product changes)
- [ ] Choose approach: simplified (spreadsheet) or advanced (Robyn/Meridian)
- [ ] Run initial model and validate against known results
- [ ] Present findings to leadership
- [ ] Implement 10-20% budget reallocation as a test
- [ ] Re-run model quarterly with updated data

### Open-Source MMM Tools

| Tool | By | Language | Strengths |
|------|-----|---------|----------|
| **Robyn** | Meta | R | Automated optimization, wide adoption |
| **Meridian** | Google | Python | Bayesian approach, causal inference |
| **LightweightMMM** | Google | Python | Simpler implementation |
| **PyMC-Marketing** | PyMC | Python | Flexible Bayesian modeling |

---

## Integration with Marketing36z OS

| System Component | How This Playbook Connects |
|-----------------|---------------------------|
| **Attribution Guide** | MMM complements attribution — macro view + micro view together |
| **Budget Template** | MMM outputs drive budget allocation decisions |
| **Dashboard Templates** | Channel ROI from MMM populates executive dashboards |
| **Experiment Framework** | Holdout and geo tests validate MMM predictions |
| **Strategy Planner** | MMM informs long-term channel strategy |
