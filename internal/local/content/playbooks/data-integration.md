# Data Integration Layer — GA4 + Google Search Console

*Source: SEO Machine data integration architecture — real analytics powering decisions, not guesswork (Roadmap item C7)*

> Every recommendation in Marketing36z OS should be informed by actual data. This guide establishes the integration architecture for connecting Google Analytics 4 and Google Search Console to power data-driven marketing decisions.

---

## Architecture Overview

```
┌────────────────────┐     ┌────────────────────┐
│   Google Analytics  │     │  Google Search      │
│   4 (GA4)          │     │  Console (GSC)      │
└─────────┬──────────┘     └──────────┬──────────┘
          │                           │
          │  GA4 Data API             │  Search Analytics API
          │  (Reporting)              │  (Performance)
          ▼                           ▼
┌─────────────────────────────────────────────────┐
│           Data Integration Layer                 │
│  ┌──────────┐ ┌──────────┐ ┌──────────────────┐│
│  │ Traffic  │ │ Keyword  │ │ Content           ││
│  │ Analysis │ │ Analysis │ │ Performance       ││
│  └──────────┘ └──────────┘ └──────────────────┘│
└─────────────────────┬───────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────┐
│           Marketing36z OS Agents                 │
│  ┌──────────┐ ┌──────────┐ ┌──────────────────┐│
│  │ Channel  │ │ Content  │ │ Strategy          ││
│  │ Operator │ │ Engine   │ │ Planner           ││
│  └──────────┘ └──────────┘ └──────────────────┘│
└─────────────────────────────────────────────────┘
```

---

## Part 1: GA4 Integration

### Required GA4 Data Points

| Data Category | Metrics | Used By Agent |
|--------------|---------|---------------|
| **Traffic** | Sessions, users, new users, bounce rate | Channel Operator |
| **Acquisition** | Source/medium, campaign, channel grouping | Channel Operator |
| **Engagement** | Avg. engagement time, engaged sessions, pages/session | Content Engine |
| **Conversions** | Goal completions, conversion rate, conversion value | Strategy Planner |
| **E-commerce** | Revenue, transactions, ARPU | Strategy Planner |
| **Retention** | Returning users, cohort retention | Growth & Retention |

### GA4 Data API Setup

**Step 1: Enable the API**

1. Go to Google Cloud Console → APIs & Services → Library
2. Enable "Google Analytics Data API" (v1beta)
3. Create a service account with Viewer permissions
4. Download the JSON credentials file

**Step 2: Grant Access**

1. In GA4 Admin → Property → Property Access Management
2. Add the service account email
3. Assign "Viewer" role (minimum required)

**Step 3: Configure Connection**

```yaml
# config/data-integration.yaml
ga4:
  enabled: true
  property_id: "properties/XXXXXXXXX"
  credentials_file: "config/ga4-credentials.json"

  # Data refresh schedule
  refresh_interval: "daily"  # daily | hourly | real-time
  lookback_window: 90  # days of historical data to maintain

  # Default dimensions for reports
  default_dimensions:
    - "date"
    - "sessionSource"
    - "sessionMedium"
    - "pagePath"

  # Default metrics
  default_metrics:
    - "sessions"
    - "totalUsers"
    - "newUsers"
    - "bounceRate"
    - "averageSessionDuration"
    - "conversions"
    - "totalRevenue"
```

### Key GA4 Reports for Marketing Decisions

**Report 1: Channel Performance (Weekly)**

| Metric | What It Tells You | Action Threshold |
|--------|------------------|-----------------|
| Sessions by channel | Which channels drive traffic | Channel below 5% → evaluate investment |
| Conversion rate by channel | Which channels convert | Rate below 1% → optimize landing page |
| Cost per acquisition by channel | Channel efficiency | CPA > 3× target → reduce spend |
| Engagement time by channel | Traffic quality | < 30s avg → traffic quality issue |

**Report 2: Content Performance (Weekly)**

| Metric | What It Tells You | Action Threshold |
|--------|------------------|-----------------|
| Pageviews by page | Popular content | Top 20% content → repurpose |
| Avg. engagement time by page | Content quality | < 1 min on long-form → rewrite |
| Bounce rate by page | Content relevance | > 80% → headline/intro mismatch |
| Conversions by landing page | Revenue-driving content | High traffic + low conversion → add CTAs |

**Report 3: Funnel Analysis (Monthly)**

Track the full journey: Visit → Signup → Activation → Paid → Retained

| Funnel Step | Metric | Healthy Benchmark |
|------------|--------|:----------------:|
| Visit → Signup | Signup rate | 2-5% (B2B SaaS) |
| Signup → Activated | Activation rate | 40-60% |
| Activated → Paid | Conversion rate | 15-25% |
| Paid → Month 2 | Retention rate | 85-95% |

---

## Part 2: Google Search Console Integration

### Required GSC Data Points

| Data Category | Metrics | Used By Agent |
|--------------|---------|---------------|
| **Performance** | Clicks, impressions, CTR, avg. position | Channel Operator |
| **Keywords** | Query performance by page | Channel Operator, Content Engine |
| **Pages** | Per-page search performance | Content Engine |
| **Sitemaps** | Indexing status | Channel Operator |
| **Coverage** | Index errors, excluded pages | Channel Operator |

### Search Analytics API Setup

**Step 1: Enable the API**

1. Google Cloud Console → Enable "Search Console API"
2. Use the same service account from GA4 setup
3. In GSC → Settings → Users and permissions → Add service account as Owner

**Step 2: Configure Connection**

```yaml
# config/data-integration.yaml (continued)
gsc:
  enabled: true
  site_url: "https://yourdomain.com"  # or "sc-domain:yourdomain.com" for domain property
  credentials_file: "config/ga4-credentials.json"  # same service account

  refresh_interval: "daily"
  lookback_window: 90  # GSC stores 16 months max

  # Query filters
  default_dimensions:
    - "query"
    - "page"
    - "date"
    - "country"
    - "device"

  # Minimum thresholds (filter noise)
  min_impressions: 10
  min_clicks: 1
```

### Key GSC Reports for Marketing Decisions

**Report 1: Keyword Opportunities (Weekly)**

Pull all queries where your average position is 4-20 (page 1 potential):

| Filter | Criteria | Action |
|--------|---------|--------|
| Position 4-10 | Close to page 1 top spots | Optimize existing content — add depth, update, improve |
| Position 11-20 | Page 2 — within striking distance | Strengthen content, build internal links, add backlinks |
| High impressions + low CTR | Visible but not clicked | Rewrite title tag and meta description |
| High CTR + low impressions | Niche but engaged audience | Find related keywords to expand reach |

**Report 2: Content Gap Identification (Monthly)**

Cross-reference GSC queries with published content:

1. Export all queries driving impressions
2. Group by topic cluster
3. Identify clusters with impressions but no dedicated content → content gap
4. Prioritize using the 8-Factor Opportunity Score

**Report 3: Cannibalization Detection (Monthly)**

Use GSC data to run the Keyword Cannibalization Detection process (see Channel Operator):

1. Export queries where 2+ pages appear for the same keyword
2. Check if position alternates between pages (ranking flux)
3. Flag queries where neither page cracks top 3 despite combined impressions

**Report 4: Technical SEO Health (Weekly)**

| Check | Source | Alert Threshold |
|-------|--------|:---------------:|
| Index coverage errors | Coverage report | Any new errors |
| Mobile usability issues | Mobile Usability | Any issues |
| Core Web Vitals | CWV report | Any "poor" URLs |
| Sitemap status | Sitemaps | Submitted ≠ indexed > 20% gap |

---

## Part 3: Combined Analytics Dashboard

### The Marketing Decision Dashboard

Combine GA4 + GSC data into a unified view for weekly marketing reviews:

```
┌─────────────────────────────────────────────────────────────┐
│ MARKETING DECISION DASHBOARD — Week of [Date]               │
├──────────────────────┬──────────────────────────────────────┤
│ TRAFFIC (GA4)        │ SEARCH (GSC)                         │
│ Sessions: X,XXX      │ Total clicks: X,XXX                  │
│ vs. last week: +X%   │ vs. last week: +X%                   │
│ New users: X,XXX     │ Avg. position: X.X                   │
│ Bounce rate: XX%     │ Avg. CTR: X.X%                       │
├──────────────────────┼──────────────────────────────────────┤
│ CONVERSIONS (GA4)    │ OPPORTUNITIES (GSC)                   │
│ Signups: XXX         │ Keywords pos. 4-10: XX                │
│ Conv. rate: X.X%     │ Keywords pos. 11-20: XX               │
│ Revenue: $X,XXX      │ High imp/low CTR pages: XX            │
│ ARPU: $XX            │ New cannibalization flags: X           │
├──────────────────────┴──────────────────────────────────────┤
│ TOP ACTIONS THIS WEEK                                        │
│ 1. [Auto-generated from data: e.g., "Optimize title for     │
│    'marketing automation' — pos 6, CTR 1.2% (below avg)"]   │
│ 2. [...]                                                     │
│ 3. [...]                                                     │
└─────────────────────────────────────────────────────────────┘
```

### Automated Alerts

Configure alerts for conditions that require immediate action:

| Alert | Condition | Severity | Action |
|-------|----------|:--------:|--------|
| Traffic drop | Sessions -20% WoW | High | Investigate — algorithm update? Tracking broken? |
| Conversion drop | Conv rate -15% WoW | Critical | Check funnel — form broken? Pricing page changed? |
| Ranking loss | Key page drops 5+ positions | High | Check content freshness, backlink changes, competitor moves |
| Index coverage | New errors > 5 | Medium | Fix technical issues ASAP |
| CWV failure | Core Web Vital fails "good" | Medium | Performance optimization sprint |

---

## Part 4: Data-Driven Agent Enhancements

### How Each Agent Uses the Data

| Agent | Data Source | Enhancement |
|-------|-----------|-------------|
| **Channel Operator** | GA4 acquisition + GSC performance | Recommend channel allocation based on actual ROI, not assumptions |
| **Content Engine** | GA4 engagement + GSC queries | Prioritize topics based on search demand, optimize underperforming content |
| **Strategy Planner** | GA4 conversions + revenue | Set realistic targets based on historical trends, identify growth levers |
| **Quality Guardian** | GA4 engagement time + bounce rate | Use engagement data to validate content quality scores |
| **Orchestrator** | All combined | Generate the 90-day plan using actual performance data as the baseline |

### Data Freshness Requirements

| Decision Type | Data Freshness Needed | Refresh Rate |
|--------------|:--------------------:|:------------:|
| Content prioritization | 7-day rolling | Daily |
| Channel budget allocation | 30-day rolling | Weekly |
| Quarterly strategy review | 90-day rolling | Monthly |
| A/B test analysis | Real-time or daily | Per experiment |
| Alert monitoring | 24-hour rolling | Daily |

---

## Implementation Checklist

| Step | Task | Priority |
|:----:|------|:--------:|
| 1 | Set up GA4 property with proper event tracking | Critical |
| 2 | Verify GSC ownership and submit sitemap | Critical |
| 3 | Create service account and enable APIs | Critical |
| 4 | Configure `data-integration.yaml` with credentials | Critical |
| 5 | Set up weekly automated report generation | High |
| 6 | Configure automated alerts | High |
| 7 | Build the Marketing Decision Dashboard | High |
| 8 | Connect data outputs to agent decision logic | Medium |
| 9 | Create monthly data review process | Medium |
| 10 | Set up quarterly deep-dive analysis template | Low |
