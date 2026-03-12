# Playbook: Revenue Operations (RevOps) Integration

## Sources
- Winning by Design — Revenue Architecture
- Clari / Gong — RevOps Best Practices
- SaaStr — Revenue Operations frameworks
- Marketing36z OS synthesis

---

## Part 1: What RevOps Is

### RevOps = Alignment of Marketing + Sales + Customer Success

```
Traditional (Siloed):
Marketing → (throw leads over wall) → Sales → (throw customers over wall) → CS

RevOps (Aligned):
Marketing ←→ Sales ←→ Customer Success
         ↕           ↕           ↕
      ────── Revenue Operations ──────
      (Shared data, processes, goals)
```

### The RevOps Mandate

| Function | What RevOps Aligns |
|----------|-------------------|
| **Process** | Standardized handoffs, SLAs, deal stages across all teams |
| **Data** | Single source of truth for leads, pipeline, revenue |
| **Technology** | Unified tech stack (CRM, analytics, automation) |
| **Metrics** | Shared revenue metrics, not departmental vanity metrics |
| **Incentives** | Compensation aligned to revenue outcomes, not activity |

---

## Part 2: The Revenue Funnel

### Unified Revenue Stages

| Stage | Owner | Definition | Exit Criteria | SLA |
|-------|-------|-----------|--------------|-----|
| **Visitor** | Marketing | Anonymous website traffic | Identifies themselves | — |
| **Lead** | Marketing | Known contact (email captured) | Meets MQL criteria | — |
| **MQL** | Marketing | Marketing Qualified Lead | Accepted by sales | <4 hours to handoff |
| **SAL** | Sales | Sales Accepted Lead | Sales confirms fit | <24 hours acceptance |
| **SQL** | Sales | Sales Qualified Lead | Discovery call completed | <48 hours first contact |
| **Opportunity** | Sales | Active deal in pipeline | Budget, timeline confirmed | Track days in stage |
| **Customer** | CS / Sales | Closed won, payment received | Onboarding starts | <24 hours CS handoff |
| **Advocate** | CS / Marketing | Active promoter | NPS 9+, referral activity | Ongoing nurture |

### The Bowtie Model

```
         ACQUISITION                          EXPANSION

   Awareness                              Adoption
        ↓                                     ↓
   Consideration           ──▷            Expansion
        ↓              CLOSE                  ↓
   Decision           ◁──                Advocacy
        ↓                                     ↓
   Purchase                               Referral
```

The bowtie model recognizes that revenue doesn't end at "Close Won" — it expands through the customer lifecycle.

---

## Part 3: Marketing ↔ Sales Alignment

### The Lead Handoff Process

| Step | Action | Owner | SLA |
|:----:|--------|-------|:---:|
| 1 | Lead reaches MQL threshold | Marketing Ops | — |
| 2 | MQL notification sent to assigned rep | Marketing Ops | Real-time |
| 3 | Sales rep reviews and accepts/rejects | Sales rep | <4 hours |
| 4 | If accepted → first outreach | Sales rep | <24 hours |
| 5 | If rejected → feedback to marketing with reason | Sales rep | <24 hours |
| 6 | Rejected leads return to nurture or disqualified | Marketing | — |

### Lead Disposition Reasons (Sales → Marketing Feedback)

| Disposition | Definition | Marketing Action |
|------------|-----------|-----------------|
| **Accepted** | Meets SQL criteria, pursuing | Track conversion |
| **Rejected — Wrong fit** | Not ICP (wrong size, industry, etc.) | Refine scoring / targeting |
| **Rejected — Not ready** | Right fit but not in buying mode | Return to nurture sequence |
| **Rejected — Bad data** | Incorrect contact info | Clean data, review capture |
| **Rejected — Duplicate** | Already in pipeline or customer | Deduplicate |

### Shared Metrics (Marketing + Sales)

| Metric | Why It's Shared | Owner |
|--------|----------------|-------|
| **Pipeline generated** | Marketing creates, sales progresses | Joint |
| **MQL → SQL conversion rate** | Lead quality indicator | Marketing + Sales |
| **Sales cycle length** | Efficiency of the whole funnel | Joint |
| **Win rate** | Sales effectiveness + lead quality | Joint |
| **Revenue per lead** | End-to-end funnel efficiency | RevOps |
| **Customer acquisition cost (CAC)** | Full funnel economics | RevOps |

---

## Part 4: Marketing ↔ Customer Success Alignment

### The Customer Handoff (Sales → CS)

| Step | Action | Owner | SLA |
|:----:|--------|-------|:---:|
| 1 | Deal closed, customer record updated | Sales | Same day |
| 2 | Internal handoff meeting (sales context to CS) | Sales + CS | <48 hours |
| 3 | Customer welcome + onboarding scheduled | CS | <48 hours |
| 4 | Marketing enters customer into success content track | Marketing | <1 week |

### Handoff Information Package

| Info | Source | Why CS Needs It |
|------|--------|----------------|
| Original pain points | Sales notes | Tailor onboarding to their specific needs |
| Decision criteria | CRM | Understand what they value most |
| Champion and stakeholders | CRM | Know who to engage |
| Expectations discussed | Sales notes | Deliver on promises made |
| Contract terms | CRM | Know what they're paying for |
| Timeline commitments | CRM | Meet deadlines |

### Customer Marketing Collaboration

| Activity | Marketing Role | CS Role |
|----------|--------------|---------|
| **Onboarding content** | Create educational sequences | Personalize and deliver |
| **Product adoption** | Feature announcement campaigns | 1:1 coaching |
| **Expansion** | Upgrade campaign | Identify opportunity, facilitate conversation |
| **Advocacy** | Case study, review, referral programs | Identify candidates, facilitate introductions |
| **Retention** | Re-engagement campaigns | Proactive health monitoring |

---

## Part 5: RevOps Technology Stack

### The Integrated Tech Stack

| Layer | Purpose | Examples |
|-------|---------|---------|
| **CRM** (foundation) | Single source of truth for all customer data | Salesforce, HubSpot |
| **Marketing automation** | Email, lead scoring, nurture sequences | HubSpot, Marketo, ActiveCampaign |
| **Sales engagement** | Outreach sequences, call tracking | Outreach, SalesLoft, Apollo |
| **Revenue intelligence** | Call recording, deal insights | Gong, Chorus |
| **CS platform** | Health scoring, renewal management | Gainsight, ChurnZero, Vitally |
| **Data warehouse** | Unified data for analytics | Snowflake, BigQuery |
| **BI / Reporting** | Dashboards and reports | Looker, Tableau, Looker Studio |
| **CDP** | Unified customer profiles | Segment, mParticle |

### Data Flow Architecture

```
Marketing Tools ──▷ ┌──────────────┐ ──▷ BI Dashboard
                    │              │
Sales Tools    ──▷  │   CRM        │ ──▷ Revenue Reports
                    │   (SSOT)     │
CS Tools       ──▷  │              │ ──▷ Forecasting
                    └──────────────┘
                           ▲
                    Data Warehouse
                    (historical +
                     cross-system)
```

---

## Part 6: RevOps Metrics Dashboard

### The Revenue Waterfall

| Metric | This Month | Last Month | MoM Δ | Target |
|--------|:----------:|:----------:|:-----:|:------:|
| Website visitors | | | % | |
| Leads generated | | | % | |
| MQLs | | | % | |
| SALs (accepted) | | | % | |
| SQLs | | | % | |
| Opportunities created | | | % | |
| Deals closed | | | % | |
| New customers | | | % | |
| Expansion revenue | $ | $ | % | $ |
| Churned revenue | $ | $ | % | $ |
| **Net new ARR** | **$** | **$** | **%** | **$** |

### Conversion Rates Between Stages

| Transition | Rate | Benchmark | Status |
|-----------|:----:|:---------:|:------:|
| Visitor → Lead | % | 2-5% | 🟢🟡🔴 |
| Lead → MQL | % | 10-30% | 🟢🟡🔴 |
| MQL → SAL | % | 60-80% | 🟢🟡🔴 |
| SAL → SQL | % | 50-70% | 🟢🟡🔴 |
| SQL → Opportunity | % | 30-50% | 🟢🟡🔴 |
| Opportunity → Close | % | 20-35% | 🟢🟡🔴 |

---

## Part 7: RevOps Implementation Roadmap

### Phase 1: Foundation (Month 1-2)
- [ ] Define unified revenue stages (agree across marketing, sales, CS)
- [ ] Document lead handoff SLAs
- [ ] Clean CRM data (deduplication, stage definitions)
- [ ] Build basic revenue waterfall dashboard

### Phase 2: Alignment (Month 3-4)
- [ ] Implement lead scoring model (marketing + sales input)
- [ ] Set up sales disposition feedback loop
- [ ] Create customer handoff process (sales → CS)
- [ ] Launch shared weekly metrics review

### Phase 3: Optimization (Month 5-6)
- [ ] Connect marketing automation to CRM with proper attribution
- [ ] Build cohort-based pipeline analysis
- [ ] Implement customer health scoring
- [ ] Align compensation to shared revenue metrics

### Phase 4: Intelligence (Month 7+)
- [ ] Revenue forecasting with ML / predictive models
- [ ] Full-funnel attribution (marketing → sale → expansion)
- [ ] Automated lead routing based on scoring + fit
- [ ] Revenue intelligence (conversation analytics, deal insights)

---

## Integration with Marketing36z OS

| System Component | How This Playbook Connects |
|-----------------|---------------------------|
| **Marketing-Sales SLA** | RevOps operationalizes the SLA |
| **Lead Scoring** | Lead scoring model feeds RevOps handoff criteria |
| **Pipeline Forecasting** | RevOps data powers accurate forecasting |
| **Attribution Guide** | Attribution spans the full revenue funnel |
| **Expansion Campaigns** | CS signals feed marketing expansion programs |
| **Dashboard Templates** | RevOps dashboard unifies marketing, sales, and CS metrics |
