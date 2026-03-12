# Playbook: Attribution & Modern Measurement

## Sources
- Avinash Kaushik — *Web Analytics 2.0*, Multi-Channel Attribution
- Rand Fishkin — *Lost and Founder*, Attribution Reality for Startups
- Chris Walker — Dark Social and Self-Reported Attribution
- Marketing36z OS synthesis

---

## Part 1: The Attribution Problem

### Why Attribution Is Hard (and Why You Still Need It)

> **"Attribution is a model, not a fact."** — Avinash Kaushik

Every customer touches multiple channels before buying. Attribution models try to assign credit for the conversion. No model is perfect — but having a model is infinitely better than guessing.

84% of B2B content sharing happens through "dark" channels invisible to analytics — DMs, Slack messages, text messages, email forwards, and verbal recommendations. Traditional attribution captures only the tip of the iceberg. Top 1% marketers build for both the visible and invisible funnel.

### The Attribution Spectrum

| Approach | Sophistication | Best For | Limitation |
|----------|:---:|----------|-----------|
| **No attribution** | 0 | Nobody | You're flying blind |
| **Last-touch** | 1 | Founders just starting | Ignores everything before the final click |
| **First-touch** | 2 | Understanding demand generation | Ignores nurture and closing channels |
| **Linear** | 3 | Teams with multi-touch funnels | Treats all touches equally (they're not) |
| **Time-decay** | 4 | Longer sales cycles | Undervalues awareness-stage content |
| **Position-based (U-shaped)** | 4 | Most B2B companies | Arbitrary weight allocation |
| **Data-driven / algorithmic** | 5 | Companies with significant data volume | Requires 600+ conversions/month |

---

## Part 2: Understanding Dark Social

### What Is Dark Social?
Dark social refers to content sharing and influence that happens through private, untrackable channels:

| Channel | How Sharing Happens | Trackable? |
|---------|-------------------|:----------:|
| Slack DMs/Groups | Link or screenshot shared in work channels | No |
| Text/iMessage | "Check this out" texts between colleagues | No |
| Email forwards | Forwarding newsletters, articles, reports | Partially |
| Private communities | Sharing in gated Slack/Discord groups | No |
| Word of mouth | "Have you heard of [brand]?" conversations | No |
| Podcast mentions | Host recommends a product | Partially |
| Meeting references | "I saw this framework from [brand]..." | No |

### The Dark Funnel
The buyer's journey that happens before they ever touch your website:

```
DARK FUNNEL (invisible)                    VISIBLE FUNNEL (tracked)
━━━━━━━━━━━━━━━━━━━━━━━━                  ━━━━━━━━━━━━━━━━━━━━━━
Colleague mentions you    ──┐
Sees your LinkedIn post   ──┤
Hears podcast mention     ──┤
Reads forwarded article   ──┼──→  Google searches your name  ──→  Visits website
Gets Slack recommendation ──┤     (Brand search = dark       ──→  Requests demo
Sees screenshot in group  ──┘      funnel indicator)          ──→  Becomes customer
```

**Key insight:** By the time someone types your brand name into Google, the marketing has already happened in the dark funnel. The "last touch" gets credit, but the dark funnel did the work.

---

## Part 3: Choosing Your Attribution Model

### The Decision Matrix

| Your Situation | Recommended Model | Why |
|---------------|------------------|-----|
| <$10K/mo spend, 1-2 channels | Last-touch + self-reported | Simple, actionable, catches dark social |
| $10K-50K/mo, 3-5 channels | Position-based (U-shaped) | Credits discovery AND closing |
| $50K-200K/mo, 5+ channels | Time-decay or data-driven | Enough data to model the journey |
| $200K+/mo, enterprise sales | Multi-touch + MMM hybrid | Statistical rigor at scale |

### The U-Shaped Model (Best Default for Growth-Stage)

```
First Touch: 40% credit   →   Middle Touches: 20% split   →   Last Touch: 40% credit

Example: $1,000 deal
- Blog post (first touch): $400 credit
- LinkedIn ad (middle): $66 credit
- Email click (middle): $66 credit
- Webinar (middle): $68 credit
- Demo request page (last touch): $400 credit
```

### Which Model to Use (By Company Stage)

| Company Stage | Recommended Model | Why |
|--------------|------------------|-----|
| Pre-revenue to $500K ARR | First-touch + Last-touch (run both) | Simple, directional, low data volume |
| $500K-$5M ARR | U-Shaped or Time-Decay | Captures journey without complexity |
| $5M-$20M ARR | W-Shaped + Self-Reported | Matches B2B buying stages |
| $20M+ ARR | Data-Driven + Incrementality | Enough data for ML models |

---

## Part 4: Self-Reported Attribution (HDYHAU)

### Why Self-Reported Matters

Most B2B attribution misses the reality: buyers discover you through podcasts, word-of-mouth, Slack communities, and DMs that no tracking pixel captures.

**Rule:** Run BOTH software attribution AND self-reported attribution. Where they disagree, self-reported is usually more accurate for the *discovery* moment.

### Implementation Guide

#### Step 1: Add "How Did You Hear About Us?" (HDYHAU)

Add an **open-text field** (not dropdown) to:
- Demo request forms
- Signup flows
- Checkout pages
- Sales qualification calls

**Why open-text, not dropdown:** Dropdowns constrain responses to categories you've already imagined. Open-text reveals channels you didn't know existed (e.g., "my friend Sarah told me," "saw your thread on Twitter last week," "your CEO's podcast with Lenny").

#### Step 2: Categorize Responses Monthly

| Category | Example Responses | Dark Funnel? |
|----------|------------------|:------------:|
| **Word of mouth** | "friend told me," "colleague recommended" | Yes |
| **Podcast** | "heard on Lenny's Podcast," "your CEO on [show]" | Yes |
| **LinkedIn** | "saw your post on LinkedIn," "LinkedIn ad" | Partially |
| **Twitter/X** | "your thread about [topic]" | Partially |
| **Google Search** | "googled it," "searched for [keyword]" | No |
| **Community** | "someone in [Slack/Discord] mentioned you" | Yes |
| **Event** | "met at [conference]," "your webinar" | Partially |
| **Newsletter** | "read your newsletter," "saw in [publication]" | Partially |
| **Content** | "read your blog post about [topic]" | No |
| **Referral program** | "[person] referred me" | Yes |

#### Step 3: Compare Self-Reported vs Tool-Based

| Source | Tool-Based Attribution | Self-Reported | Gap |
|--------|:---------------------:|:-------------:|:---:|
| Google Search | 35% | 15% | Tool over-credits |
| Direct | 25% | 5% | Tool over-credits |
| Podcast | 0% | 20% | Tool misses entirely |
| Word of mouth | 0% | 25% | Tool misses entirely |
| LinkedIn (organic) | 5% | 15% | Tool under-credits |
| Community | 0% | 10% | Tool misses entirely |
| Paid ads | 30% | 8% | Tool over-credits |

**Typical finding:** Self-reported reveals that podcast, word-of-mouth, and community are 3-5x more influential than tool-based attribution shows.

#### Step 4: Blend Into Budget Decisions

**Recommended weighting:** 50% tool-based + 50% self-reported

If self-reported shows podcast drives 20% of customers but tool-based shows 0%, the truth is somewhere in between. Allocate budget accordingly.

---

## Part 5: Multi-Touch Attribution Models

### Model Comparison

| Model | How Credit Is Assigned | Best For | Limitations |
|-------|----------------------|----------|-------------|
| **First-Touch** | 100% to first interaction | Understanding awareness channels | Ignores nurture and conversion |
| **Last-Touch** | 100% to final interaction | Understanding conversion triggers | Ignores discovery and nurture |
| **Linear** | Equal credit to all touchpoints | Balanced view of full journey | Over-credits low-value touches |
| **Time-Decay** | More credit to recent touches | Long sales cycles (60+ days) | Under-credits top-of-funnel |
| **U-Shaped** | 40% first, 40% last, 20% middle | B2B with clear discovery + conversion | Arbitrary split |
| **W-Shaped** | 30% first, 30% lead creation, 30% opportunity creation, 10% rest | B2B with defined stages | Requires stage tracking in CRM |
| **Data-Driven** | ML model determines credit | Large datasets, mature orgs | Requires 1000+ conversions |

---

## Part 6: The Attribution Tech Stack

### By Company Stage

| Stage | Tool | Cost | What It Does |
|-------|------|:----:|-------------|
| **Pre-revenue** | UTM parameters + Google Analytics | Free | Basic campaign tracking |
| **$0-50K MRR** | GA4 + self-reported field | Free | Multi-touch view + dark social |
| **$50K-200K MRR** | HubSpot / Segment + GA4 | $$$  | Full journey mapping, CRM integration |
| **$200K+ MRR** | Dreamdata / Factors.ai / HockeyStack | $$$$ | B2B-specific multi-touch attribution |
| **Enterprise** | Custom data warehouse + Looker/Tableau | $$$$$ | Full marketing mix modeling |

---

## Part 7: UTM Parameter Standards

Every link you share should have UTM parameters:

| Parameter | Format | Example |
|-----------|--------|---------|
| `utm_source` | Platform name | `linkedin`, `google`, `newsletter` |
| `utm_medium` | Channel type | `organic`, `cpc`, `email`, `social` |
| `utm_campaign` | Campaign name | `q1-launch`, `pmf-webinar-series` |
| `utm_content` | Creative variant | `headline-a`, `image-b`, `cta-demo` |
| `utm_term` | Keyword (paid search only) | `marketing+automation` |

**Naming convention rules:**
1. Always lowercase
2. Use hyphens, never spaces or underscores
3. Be specific but consistent
4. Document your naming convention and share with the team

---

## Part 8: Building Your Attribution Dashboard

### The Executive View (Board / CEO)

| Metric | Definition | Target |
|--------|-----------|--------|
| **Blended CAC** | Total marketing spend / new customers | Decreasing trend |
| **CAC by channel** | Channel spend / channel-attributed customers | Below LTV/3 |
| **CAC payback period** | Months to recover CAC | <12 months |
| **Marketing-sourced pipeline** | % of pipeline from marketing | >50% |
| **Marketing-influenced revenue** | Revenue where marketing touched the journey | >70% |

### The Marketing Manager View

| Metric | Definition | Frequency |
|--------|-----------|-----------|
| **Channel-level CAC** | Spend per channel / attributed conversions | Weekly |
| **Conversion by touch count** | Avg. touches before conversion | Monthly |
| **Time to conversion** | Days from first touch to close | Monthly |
| **Top converting paths** | Most common channel sequences | Monthly |
| **Content attribution** | Which content pieces appear in winning journeys | Monthly |
| **Self-reported vs. tracked source** | Delta between what analytics says and what customers say | Monthly |

### The Operator View (Channel Owners)

| Metric | Definition | Frequency |
|--------|-----------|-----------|
| **First-touch by source** | Which channels generate awareness | Weekly |
| **Assisted conversions** | Touches that didn't close but contributed | Weekly |
| **Channel velocity** | Time from channel touch to next action | Weekly |
| **Creative-level performance** | Which specific ads/posts drive attribution | Weekly |
| **UTM coverage rate** | % of traffic with proper UTM tags | Weekly |

---

## Part 9: Attribution for Different Business Models

### SaaS (Product-Led Growth)

| Journey Stage | Key Attribution Point | How to Track |
|--------------|----------------------|-------------|
| Awareness | First website visit | UTM source + self-reported |
| Signup | Free trial / freemium registration | Conversion event + source |
| Activation | First value moment in-product | Product analytics (Mixpanel, Amplitude) |
| Conversion | Free → paid upgrade | CRM + product event |
| Expansion | Upgrade / add seats | CRM + in-app trigger |

### SaaS (Sales-Led)

| Journey Stage | Key Attribution Point | How to Track |
|--------------|----------------------|-------------|
| Awareness | First content touch | UTM + content downloads |
| Engagement | Multiple content interactions | Lead scoring + page views |
| MQL | Meets lead score threshold | CRM automation |
| SQL | Sales accepts the lead | CRM disposition |
| Opportunity | Demo / proposal sent | CRM pipeline stage |
| Close | Contract signed | CRM + revenue attribution |

### E-Commerce / DTC

| Journey Stage | Key Attribution Point | How to Track |
|--------------|----------------------|-------------|
| Discovery | First site visit or ad impression | GA4 + Meta/Google Ads |
| Consideration | Product page view, cart add | Enhanced e-commerce events |
| Purchase | Checkout complete | Conversion tracking + UTM |
| Repeat purchase | Second order | Cohort analysis, email attribution |
| Referral | Customer-driven new purchase | Referral tracking code |

---

## Part 10: Content Designed for Dark Social

### High-Share Content Formats

| Format | Why It Gets Shared Privately | Design Tips |
|--------|----------------------------|------------|
| **Original data/stats** | Makes the sharer look informed | Clean single-stat graphics, branded |
| **Framework visuals** | Makes the sharer look smart | 2x2 matrices, flowcharts, checklists |
| **Contrarian takes** | Sparks group discussion | Strong opinion + evidence |
| **"I wish I knew this" posts** | Helps others avoid mistakes | Personal, specific, actionable |
| **Industry benchmarks** | Useful reference for teams | Comparison tables with percentages |
| **Templates/checklists** | Immediately useful | Download/screenshot-ready |

### Dark Social Content Design Rules

1. **Screenshot-ready design:** Your best content should look good as a screenshot (people screenshot and share more than they link)
2. **Self-contained value:** The insight should be clear WITHOUT clicking through
3. **Branded but not salesy:** Include your brand subtly so it travels with the share
4. **"Forward this to..." CTAs:** Explicitly prompt sharing: "Know a founder who needs this? Forward it to them"
5. **One insight per piece:** Make it easy to describe in a DM: "Check this out, it's about [X]"
6. **Quotable statements:** Include lines worth quoting in conversations

### The "Would They Screenshot This?" Test

Before publishing, ask: "Would someone screenshot this and send it to their team's Slack channel?"

If not, add:
- A surprising data point
- A useful framework or matrix
- A contrarian take with evidence
- A template they can use immediately

---

## Part 11: Incrementality Testing

### The Gold Standard for Measuring True Impact

Attribution tells you who touched the deal. Incrementality tells you whether the deal would have happened WITHOUT your marketing.

### Test Types

#### Geo Holdout Test
1. Select 2-4 comparable geographic regions (similar size, industry mix, growth rate)
2. Continue marketing in "control" regions
3. Turn off specific marketing in "test" regions for 4-6 weeks
4. Compare pipeline, revenue, and brand search between test and control
5. Calculate true incremental lift

**When to use:** Any channel with $5K+/month spend

#### On/Off Test
1. Pause a specific channel entirely for 2-4 weeks
2. Measure impact on total pipeline (not just that channel's attributed pipeline)
3. If pipeline drops: the channel has true incremental impact
4. If pipeline stays flat: the channel may be getting credit for organic demand

**When to use:** Channels you suspect are over-credited by attribution

#### PSA Test (Public Service Announcement)
1. Show real ads to control group
2. Show PSA/charity ads to test group (same targeting, same spend)
3. Compare conversion between groups
4. Difference = true incremental impact of your creative

**When to use:** Digital advertising, especially retargeting

---

## Part 12: Dark Funnel Measurement

### Key Indicators

Track these as proxies for dark social and word-of-mouth impact:

| Indicator | How to Track | What Growth Means |
|-----------|-------------|-------------------|
| **Branded search volume** | Google Search Console, SEMrush | More people know you exist |
| **Direct traffic** | Google Analytics (remove bots) | People typing your URL directly |
| **Self-reported "friend/colleague"** | HDYHAU surveys | Word-of-mouth is working |
| **Unattributed demo requests** | CRM (no marketing source) | Dark funnel is driving action |
| **Social mentions (unprompted)** | Social listening tools | Organic brand advocacy |
| **Branded search CTR** | Google Search Console | Brand recognition in search |

### Correlation Analysis

Look for correlations between content/social activity and dark funnel indicators:

1. Map content publishing dates and social activity spikes
2. Track branded search volume changes 1-2 weeks after major content pushes
3. Look for direct traffic spikes after podcast appearances
4. Track HDYHAU "word of mouth" percentage over time vs content investment

**If branded search grows while content investment grows → your dark funnel is working.**

---

## Part 13: Common Attribution Mistakes

| Mistake | Why It's Wrong | Fix |
|---------|---------------|-----|
| Only tracking last touch | Ignores 90% of the journey | Add first-touch and self-reported |
| Ignoring dark social | Podcasts, DMs, communities drive B2B — invisible to pixels | Self-reported attribution field |
| Over-crediting paid | Paid gets credit because it's trackable, not because it's most impactful | Compare with holdout tests |
| Not tracking offline | Events, word-of-mouth, partnerships disappear | Ask "how did you hear about us?" |
| Changing models mid-experiment | Breaks comparability | Lock model for at least one quarter |
| Attributing to the landing page, not the source | "Homepage" is not a source | Always look at the referrer/UTM |
| Zero attribution for brand | Brand spend shows up everywhere else | Run brand lift studies or holdout tests |

---

## Part 14: Attribution Maturity

### Attribution Maturity Model (Five Levels)

| Level | Description | Tools | Accuracy |
|:-----:|------------|-------|:--------:|
| **1 - Basic** | Last-click only | Google Analytics | Low |
| **2 - Aware** | First-touch + last-touch | GA + CRM | Medium-Low |
| **3 - Multi-Touch** | W-shaped or time-decay model | CRM + attribution tool | Medium |
| **4 - Blended** | Multi-touch + self-reported + dark funnel indicators | Full stack + HDYHAU | Medium-High |
| **5 - Advanced** | Multi-touch + self-reported + incrementality testing + MMM | Full stack + testing program | High |

**Target:** Most companies should aim for Level 4. Only $20M+ companies need Level 5.

### Attribution Maturity Roadmap

#### Level 1: Foundation (Month 1)
- [ ] UTM parameters on all external links
- [ ] Google Analytics 4 configured with conversion events
- [ ] Self-reported "how did you hear about us?" field on key forms
- [ ] Basic channel-level spending spreadsheet

#### Level 2: Multi-Touch (Month 2-3)
- [ ] CRM tracking full journey (first touch → close)
- [ ] Position-based attribution model implemented
- [ ] Monthly attribution report created
- [ ] Self-reported data compared to software attribution monthly

#### Level 3: Integrated (Month 4-6)
- [ ] Attribution data feeds budget allocation decisions
- [ ] Content-level attribution (which blog posts / assets appear in winning paths)
- [ ] Dark social tracking formalized
- [ ] Quarterly attribution review with leadership

#### Level 4: Advanced (Month 6-12)
- [ ] Incrementality tests running (holdout/geo tests)
- [ ] Multi-touch model calibrated with self-reported data
- [ ] Attribution integrated with LTV data
- [ ] Budget optimization driven by attribution insights

---

## Part 15: Budget Allocation Using Blended Attribution

### Decision Framework

| Attribution Signal | Budget Action |
|-------------------|-------------|
| Tool-based high + Self-reported high | Increase investment (validated impact) |
| Tool-based high + Self-reported low | Investigate — may be over-attributed |
| Tool-based low + Self-reported high | Increase investment (dark funnel channel) |
| Tool-based low + Self-reported low | Decrease or cut investment |

---

## Part 16: The Attribution Report Template

### Monthly Attribution Summary

| Source | First-Touch Conversions | Last-Touch Conversions | Self-Reported Mentions | Blended CAC | Trend |
|--------|:-:|:-:|:-:|:-:|:-:|
| Organic Search | | | | $ | |
| LinkedIn (organic) | | | | $ | |
| LinkedIn (paid) | | | | $ | |
| Email | | | | $ | |
| Direct / Brand | | | | $ | |
| Referral / WOM | | | | $ | |
| Podcast | | | | $ | |
| Events | | | | $ | |
| **Total** | | | | **$** | |

### Key Insights This Month

| Question | Answer |
|----------|--------|
| **Highest ROI channel:** | |
| **Fastest time-to-close channel:** | |
| **Biggest self-reported vs. tracked gap:** | |
| **Budget reallocation recommendation:** | |
| **New channel to test next month:** | |

### Monthly Attribution Review Template

```
ATTRIBUTION REVIEW — [Month] 20XX
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

TOOL-BASED ATTRIBUTION (top 5 sources)
1. [Source]: X% of pipeline ($___K)
2. [Source]: X% of pipeline ($___K)
3. [Source]: X% of pipeline ($___K)
4. [Source]: X% of pipeline ($___K)
5. [Source]: X% of pipeline ($___K)

SELF-REPORTED ATTRIBUTION (top 5 sources)
1. [Source]: X% of new customers
2. [Source]: X% of new customers
3. [Source]: X% of new customers
4. [Source]: X% of new customers
5. [Source]: X% of new customers

DARK FUNNEL INDICATORS
• Branded search: [trend] (+/- X% MoM)
• Direct traffic: [trend] (+/- X% MoM)
• Unattributed demos: [count] (+/- X% MoM)

GAPS / INSIGHTS
• [Insight about what self-reported reveals that tools miss]
• [Insight about dark funnel trends]

BUDGET RECOMMENDATIONS
• Increase: [channel] because [evidence]
• Maintain: [channel] because [evidence]
• Investigate: [channel] because [evidence]
• Decrease: [channel] because [evidence]
```

---

## Part 17: Integration with Marketing36z OS

| System Component | How This Playbook Connects |
|-----------------|---------------------------|
| **Strategy Planner** | Attribution data informs channel strategy and budget allocation |
| **Channel Operator** | UTM standards and channel tracking feed multi-touch reporting |
| **Quality Guardian** | Reviews attribution claims for accuracy before reporting |
| **Dashboard Templates** | Attribution data populates the marketing dashboard |
| **Experiment Framework** | Incrementality tests designed using experiment playbook |
| **Budget Template** | Attribution ROI drives budget reallocation decisions |

---

## Key Takeaways

1. **Traditional attribution captures <20% of B2B influence.** Build for the dark funnel.
2. **Self-reported attribution is your most honest signal.** Add HDYHAU to every form.
3. **Design content for private sharing.** If it's not screenshot-worthy, it won't spread.
4. **Blend attribution methods.** No single model tells the whole truth.
5. **Incrementality testing is the gold standard.** Run holdout tests on your biggest channels.
6. **Brand search volume is your dark funnel barometer.** Track it weekly.
7. **Don't cut what you can't measure.** Podcast and community may show $0 in attribution but drive 30% of customers.
