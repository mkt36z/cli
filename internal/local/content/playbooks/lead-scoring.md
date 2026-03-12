# Playbook: Lead Scoring Model

## Sources
- Mark Roberge — *The Sales Acceleration Formula*
- HubSpot Lead Scoring Methodology
- Drift / David Cancel — Conversational Marketing, Intent Signals
- Marketing36z OS synthesis

---

## Part 1: Why Lead Scoring Matters

> **Not all leads are equal. Treat them equally and you'll waste sales time on tire-kickers while hot prospects go cold.**

Lead scoring assigns a numerical value to each lead based on their fit (who they are) and engagement (what they do), so sales can prioritize the highest-probability deals.

---

## Part 2: The Two Dimensions of Lead Scoring

### Dimension 1: Fit Score (Who They Are)

Fit scoring measures how closely a lead matches your Ideal Customer Profile.

| Attribute | Weight | High Score (5) | Medium (3) | Low (1) | Negative (-5) |
|-----------|:------:|---------------|-----------|---------|---------------|
| **Job title** | High | Decision maker (CEO, VP) | Manager, Director | IC, Intern | Student, Retired |
| **Company size** | High | Sweet spot (10-50 employees) | Adjacent (5-10 or 50-200) | Outside range | Solopreneur (if not ICP) |
| **Industry** | Medium | Target industry | Adjacent industry | Unrelated | Excluded industry |
| **Revenue/funding** | Medium | Funded or $1M+ revenue | Bootstrapped, growing | Pre-revenue | No budget signals |
| **Geography** | Low | Primary market | Secondary market | Tertiary market | Unsupported region |
| **Tech stack** | Medium | Uses complementary tools | Some overlap | No overlap | Uses competitor |

### Dimension 2: Engagement Score (What They Do)

Engagement scoring measures intent signals based on behavior.

| Action | Points | Decay | Reasoning |
|--------|:------:|:-----:|-----------|
| **Visited pricing page** | +15 | 14 days | Strongest buying intent signal |
| **Requested demo / contact sales** | +25 | 30 days | Explicit intent |
| **Downloaded gated content** | +10 | 21 days | Research-stage intent |
| **Attended webinar** | +10 | 21 days | Active engagement |
| **Opened email** | +1 | 7 days | Passive engagement |
| **Clicked email link** | +3 | 14 days | Active engagement |
| **Visited 3+ pages in one session** | +5 | 14 days | Active research |
| **Visited case study / testimonials** | +8 | 14 days | Evaluation-stage signal |
| **Returned after 7+ days** | +5 | 14 days | Re-engagement signal |
| **Social media engagement** | +2 | 7 days | Awareness-stage |
| **Unsubscribed from email** | -10 | Never | Disengagement |
| **No activity for 30 days** | -5 | Resets on activity | Going cold |
| **Bounced email** | -15 | Never | Bad data |

---

## Part 3: Lead Stages and Thresholds

### The Lead Lifecycle

| Stage | Fit Score | Engagement Score | Combined | Who Owns | Action |
|-------|:---------:|:----------------:|:--------:|----------|--------|
| **Subscriber** | Any | 0-10 | <20 | Marketing | Nurture with content |
| **Lead** | Any | 11-25 | 20-39 | Marketing | Continue engagement |
| **MQL** | ≥15 | ≥25 | 40-59 | Marketing → Sales | Notify sales, warm handoff |
| **SQL** | ≥15 | ≥30 | 60-79 | Sales | Active pursuit |
| **Opportunity** | ≥15 | ≥30 | 80+ | Sales | Pipeline stage |

### Threshold Calibration

**Start with these defaults, then calibrate quarterly:**

| Parameter | Starting Value | Calibration Question |
|-----------|:-------------:|---------------------|
| MQL threshold | 40 | Are too many or too few leads reaching sales? |
| SQL threshold | 60 | Are sales reps saying lead quality is good or bad? |
| Score decay period | 14-30 days | Are stale leads clogging the pipeline? |
| Pricing page weight | +15 | Do pricing page visitors actually convert at higher rates? |
| Demo request weight | +25 | Validate: what % of demo requests become customers? |

---

## Part 4: Building Your Scoring Model

### Step 1: Define Your ICP Fit Criteria

| Attribute | Ideal (5 pts) | Good (3 pts) | Acceptable (1 pt) | Disqualify (-5 pts) |
|-----------|-------------|-----------|---------------|-------------------|
| | | | | |
| | | | | |
| | | | | |
| | | | | |
| | | | | |

### Step 2: Map Your Engagement Signals

| Action | Points | Available in Your Tools? | Implementation Notes |
|--------|:------:|:------------------------:|---------------------|
| | | Yes / No | |
| | | Yes / No | |
| | | Yes / No | |
| | | Yes / No | |
| | | Yes / No | |

### Step 3: Set Thresholds

| Threshold | Score | Expected Volume | Sales Capacity Check |
|-----------|:-----:|:---------------:|---------------------|
| MQL | | /month | Can sales handle this volume? |
| SQL | | /month | Does quality justify pursuit? |
| Opportunity | | /month | Are these converting to customers? |

### Step 4: Validate with Historical Data

Take your last 50 customers and score them retroactively:

| Customer | Fit Score | Engagement Score | Combined | Actually Converted? | Cycle Length |
|----------|:---------:|:----------------:|:--------:|:-------------------:|:----------:|
| 1 | | | | Yes | days |
| 2 | | | | Yes | days |
| ... | | | | Yes | days |

**Key questions:**
- What's the average combined score of customers who actually bought?
- What's the average score of leads that didn't convert?
- Where should you set the MQL threshold to capture 80%+ of eventual buyers?

---

## Part 5: Lead Scoring by Business Model

### SaaS / Product-Led Growth

| Signal | Weight | Reasoning |
|--------|:------:|-----------|
| Signed up for free trial | +20 | Highest intent for PLG |
| Activated (hit key value moment) | +15 | Proven product-market fit for this user |
| Invited team members | +20 | Expansion signal |
| Hit usage limit | +10 | Upgrade trigger |
| Viewed upgrade / pricing page | +15 | Active buying consideration |

### B2B Services / Consulting

| Signal | Weight | Reasoning |
|--------|:------:|-----------|
| Downloaded case study in their industry | +12 | Looking for proof in their context |
| Attended industry-specific webinar | +10 | Active research |
| Multiple stakeholders from same company | +15 | Buying committee forming |
| Visited "how we work" / process page | +8 | Evaluating working relationship |
| Requested proposal / RFP response | +25 | Explicit buying intent |

### E-Commerce / DTC

| Signal | Weight | Reasoning |
|--------|:------:|-----------|
| Added to cart | +15 | Strong purchase intent |
| Viewed product 3+ times | +8 | Considering purchase |
| Signed up for restock notification | +10 | Want to buy, timing issue |
| Read reviews | +5 | Evaluation stage |
| Abandoned cart | +12 | Intent was there — recover it |

---

## Part 6: Lead Scoring Maintenance

### Monthly Review Checklist

- [ ] Review MQL → SQL conversion rate (target: 30-50%)
- [ ] Review SQL → Opportunity conversion rate (target: 50-70%)
- [ ] Check if sales is accepting or rejecting MQLs (>30% rejection = model needs tuning)
- [ ] Verify point values match actual conversion data
- [ ] Check decay rules — are stale leads being down-scored appropriately?
- [ ] Review new engagement signals to add (new content types, new features)

### Quarterly Calibration

| Metric | Healthy | Warning | Action |
|--------|:-------:|:-------:|--------|
| MQL → SQL acceptance rate | >70% | <50% | Raise MQL threshold or adjust fit criteria |
| SQL → Opportunity rate | >50% | <30% | Tighten engagement scoring |
| Avg. score of closed-won | Reference | Declining | Model drift — recalibrate |
| False positives (high score, never converts) | <20% | >30% | Overweighting wrong signals |
| False negatives (low score, converts) | <10% | >20% | Missing key intent signals |

---

## Integration with Marketing36z OS

| System Component | How This Playbook Connects |
|-----------------|---------------------------|
| **Marketing-Sales SLA** | Lead scoring defines MQL/SQL handoff criteria in the SLA |
| **Funnel Architecture** | Scoring maps to funnel stages and conversion gates |
| **Dashboard Templates** | Lead volume by score band populates pipeline dashboards |
| **Attribution Guide** | Attribution reveals which sources produce highest-scoring leads |
| **Experiment Framework** | Scoring threshold experiments improve model accuracy |
| **Quality Guardian** | Validates that scoring logic doesn't introduce bias |
