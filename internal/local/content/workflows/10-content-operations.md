# Workflow: Content Operations at Scale

## Purpose
A systems-level workflow for managing content production when volume, quality, and consistency all need to work in concert. This is for teams producing 20+ pieces of content per month across multiple channels, formats, and contributors.

## When to Use This Workflow
- Content team has 3+ contributors (in-house + freelance)
- Producing 20+ pieces of content per month
- Content spans multiple channels and formats
- Need to maintain quality at volume
- Orchestrator routes here for "content scale" or "content operations" requests

---

## Phase 1: Content Operations Architecture

### The Content Supply Chain

```
Strategy → Planning → Briefing → Production → Review → Approval → Publish → Measure → Optimize
    ↑                                                                                    │
    └────────────────────────────── Learning loop ──────────────────────────────────────┘
```

### Roles in a Content Operation

| Role | Responsibility | Tools |
|------|---------------|-------|
| **Content Strategist** | Editorial calendar, content themes, performance analysis | Analytics, planning tools |
| **Editor** | Quality control, style guide enforcement, content review | CMS, editing tools |
| **Writer(s)** | Content production from briefs | Writing tools, CMS |
| **Designer** | Visual assets, templates, thumbnails | Design tools, DAM |
| **SEO Specialist** | Keyword research, on-page optimization, technical SEO | SEO tools |
| **Distribution Manager** | Publishing, social scheduling, email integration | Social tools, ESP |
| **Content Ops Manager** | Workflow management, freelancer coordination, tooling | PM tools |

---

## Phase 2: Content Planning System

### The Content Calendar Hierarchy

| Level | Timeframe | Owner | Contains |
|-------|:---------:|-------|---------|
| **Annual themes** | 12 months | Content Strategist + CMO | Quarterly themes, major campaigns, seasonal anchors |
| **Quarterly plan** | 3 months | Content Strategist | Monthly themes, content mix, resource allocation |
| **Monthly calendar** | 4 weeks | Editor | Specific pieces, briefs, deadlines, owners |
| **Weekly sprint** | 1 week | Content Ops | In-production pieces, review queue, publishing schedule |

### Content Mix Framework

| Content Category | % of Mix | Purpose | Examples |
|-----------------|:--------:|---------|---------|
| **SEO / Evergreen** | 30-40% | Organic traffic, long-term value | How-to guides, glossary, comparisons |
| **Thought leadership** | 15-25% | Authority, brand differentiation | POV pieces, original research, predictions |
| **Demand gen** | 15-25% | Lead capture, pipeline | Gated resources, webinars, case studies |
| **Social / engagement** | 15-20% | Audience growth, community | Short-form posts, threads, polls |
| **Sales enablement** | 5-10% | Sales support | Battle cards, one-pagers, case studies |

### Content Production Capacity Planning

| Content Type | Avg. Production Time | Review Cycles | Total Turnaround |
|-------------|:-------------------:|:-------------:|:----------------:|
| Blog post (1,500-2,500 words) | 4-6 hours | 1-2 | 5-8 business days |
| Long-form guide (3,000+ words) | 8-12 hours | 2-3 | 10-15 business days |
| Social media posts (batch of 10) | 2-3 hours | 1 | 3-5 business days |
| Email newsletter | 2-3 hours | 1 | 3-5 business days |
| Case study | 6-10 hours | 2-3 (inc. client approval) | 15-25 business days |
| Video script | 3-5 hours | 1-2 | 5-8 business days |
| Infographic | 4-6 hours (writer + designer) | 1-2 | 7-10 business days |

---

## Phase 3: Content Production Workflow

### The 7-Step Production Process

| Step | Activity | Owner | SLA | Output |
|:----:|---------|-------|:---:|--------|
| 1 | **Brief** — Create content brief from template | Strategist | Day 0 | Completed brief |
| 2 | **Assign** — Match brief to writer based on expertise | Editor / Ops | Day 0-1 | Assignment confirmed |
| 3 | **Draft** — Writer produces first draft | Writer | Day 1-5 | First draft |
| 4 | **Edit** — Editor reviews for quality, voice, accuracy | Editor | Day 5-7 | Edited draft |
| 5 | **Optimize** — SEO review, formatting, visual assets | SEO + Designer | Day 7-8 | Final draft |
| 6 | **Approve** — Final approval from content strategist or client | Strategist | Day 8-9 | Approved |
| 7 | **Publish** — Schedule and distribute | Distribution | Day 9-10 | Published |

### Quality Gates at Each Step

| Step | Quality Gate | Must Pass Before Moving Forward |
|:----:|------------|-------------------------------|
| 1 (Brief) | Brief completeness check — all fields filled | Writer assignment |
| 3 (Draft) | Writer self-review checklist | Submission to editor |
| 4 (Edit) | Editorial quality standard (voice, accuracy, structure) | SEO/design optimization |
| 5 (Optimize) | SEO checklist, visual quality, formatting | Final approval |
| 6 (Approve) | Strategic alignment + brand voice check | Publishing |
| 7 (Publish) | Link check, formatting check, tracking setup | Live |

---

## Phase 4: Managing Contributors at Scale

### Freelancer Management System

| Process | Details |
|---------|---------|
| **Onboarding** | Style guide, brand voice doc, 3 example pieces, brief template, feedback preferences |
| **Brief quality** | Every assignment includes a complete brief (use content-brief-template.md) |
| **Feedback loops** | Structured feedback after first 3 pieces, then monthly |
| **Quality scoring** | Rate each piece 1-5 on: accuracy, voice match, structure, creativity, timeliness |
| **Rates** | Clear per-piece or per-word rates, payment terms documented |
| **Capacity** | Track each freelancer's monthly capacity and specialties |

### Freelancer Scorecard

| Writer | Specialty | Avg. Quality Score | On-Time Rate | Monthly Capacity | Status |
|--------|----------|:------------------:|:------------:|:----------------:|:------:|
| | | /5 | % | pieces | Active / Bench / Exit |
| | | /5 | % | pieces | Active / Bench / Exit |
| | | /5 | % | pieces | Active / Bench / Exit |

---

## Phase 5: Content Performance Measurement

### Content Metrics by Stage

| Funnel Stage | Content Type | Primary Metric | Secondary Metric |
|-------------|-------------|---------------|-----------------|
| **Awareness** | Blog, social, video | Views, impressions | Time on page, shares |
| **Consideration** | Guides, comparisons, webinars | Downloads, signups | Email subscribers |
| **Decision** | Case studies, demos, ROI calculators | MQLs, demo requests | Pipeline influenced |
| **Retention** | Knowledge base, onboarding content | Product adoption | NPS, support deflection |

### Monthly Content Performance Report

| Content Piece | Type | Channel | Views | Engagement | Leads | Revenue Influence | Score |
|--------------|------|---------|:-----:|:----------:|:-----:|:-----------------:|:-----:|
| | | | | | | $ | /10 |
| | | | | | | $ | /10 |
| | | | | | | $ | /10 |

### Content Decay Detection

Monitor for content that's losing performance:

| Indicator | Threshold | Action |
|-----------|:---------:|--------|
| Organic traffic declining MoM | >20% drop for 2+ months | Update and republish |
| Conversion rate dropping | >30% below historical | Review CTA and offer |
| High bounce rate | >70% | Improve intro, match intent better |
| Outdated information | >12 months old | Content refresh or retire |
| Low engagement on social reshare | Below average for type | Retire from reshare rotation |

---

## Phase 6: Content Operations Tooling

### The Content Ops Stack

| Need | Budget | Mid-Range | Enterprise |
|------|--------|-----------|-----------|
| **Planning / Calendar** | Google Sheets / Notion | Asana / Monday | ContentCal / Kapost |
| **Writing / Editing** | Google Docs | Grammarly Business | Writer.com |
| **SEO** | Google Search Console | Ahrefs Lite / Surfer SEO | Ahrefs / Semrush |
| **Design** | Canva | Figma | Adobe Creative Suite |
| **CMS** | WordPress / Ghost | Webflow | Custom |
| **Social scheduling** | Buffer | Hootsuite | Sprout Social |
| **Analytics** | GA4 | GA4 + Looker Studio | Full BI stack |
| **DAM (Digital Asset Management)** | Google Drive | Brandfolder | Bynder |

---

## Phase 7: Content Operations Maturity Model

| Level | Description | Indicators |
|:-----:|-----------|-----------|
| **1: Ad Hoc** | Reactive, no process | No calendar, no briefs, inconsistent quality |
| **2: Defined** | Basic process exists | Calendar, briefs, review process in place |
| **3: Managed** | Process is measured | Performance tracking, freelancer scoring, SLAs met |
| **4: Optimized** | Data-driven decisions | Content investment tied to ROI, A/B testing, decay monitoring |
| **5: Scaled** | Self-improving system | Automated workflows, predictive performance, AI-assisted production |

---

## Integration with Marketing36z OS

| System Component | How This Workflow Connects |
|-----------------|---------------------------|
| **Content Engine** | Production workflows execute Content Engine specifications |
| **Content Brief Template** | Standardized briefs feed into production workflow |
| **Content Repurposing** | Published content enters repurposing pipeline |
| **Quality Guardian** | Quality gates at each production step |
| **Channel Operator** | Distribution follows channel-specific requirements |
| **OKR Framework** | Content OKRs tracked through content performance metrics |
