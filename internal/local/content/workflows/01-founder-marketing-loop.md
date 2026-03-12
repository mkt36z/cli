# Workflow: Founder Marketing Loop (Daily/Weekly)

## Overview
The core operating loop of Marketing36z OS. This is the end-to-end workflow from a founder's marketing request to published campaign — designed to produce world-class output through strategic rigor, multi-agent collaboration, and mandatory quality gates.

## Step 0 — Strategic Intake

**Owner:** Orchestrator

Before any work begins, the Orchestrator diagnoses the request:

### Required Context Gathering
- Capture objective, timeline, channel priority, and constraints
- Pull latest context from GitHub/Notion/Drive and analytics connectors

### Strategic Pre-Flight Check
| Gate | Question | If NO |
|------|----------|-------|
| Desire exists | Is there a proven pre-existing desire we're channeling? | HALT — needs customer research |
| Offer is strong | Does it score high on the Hormozi Value Equation? | Route to Strategy Planner for offer redesign |
| Positioning is clear | Can we articulate why us vs. alternatives? | Route to Strategy Planner for positioning |
| Awareness is known | Where is the audience on Schwartz's 5-level scale? | Strategy Planner must assess first |
| Market sophistication assessed | What level is this market? | Strategy Planner must evaluate |
| Hypothesis defined | What are we testing with this campaign? | Cannot proceed without a hypothesis |
| Big Idea identified | What is the single dominant concept? | Strategy Planner must define |

**Output:** Validated intake brief with strategic context, or a redirect to foundational strategy work.

## Step 1 — LLM Council Routing

**Owner:** LLM Council (05)

- Orchestrator sends task profile to `05-llm-council`
- Council assesses task type (strategy/creative/QA/adaptation) and selects appropriate model profile
- Returns: selected model + fallback + rationale + estimated cost
- Council checks model memory for historical performance on similar task/channel combinations
- If campaign involves multiple task types, council may recommend different models for different phases

**Output:** Model selection with confidence score and cost estimate.

## Step 2 — Strategy Draft

**Owner:** Strategy Planner (01)

The Strategy Planner produces comprehensive strategic artifacts:

### Mandatory Deliverables
1. **Problem statement** — specific, in the customer's language
2. **ICP + Empathy Map** — demographics, psychographics, 3AM fears, secret desires, buying triggers
3. **Awareness level assessment** — where the audience is on Schwartz's 5 levels
4. **Market sophistication assessment** — what messaging approach is required
5. **Positioning** (April Dunford framework) — competitive alternatives, unique attributes, value, target customer, market category
6. **Offer evaluation** (Hormozi Value Equation) — dream outcome, likelihood, time delay, effort. Score and improvement recommendations
7. **Messaging pillars** (3 max) — differentiated, provable, rooted in empathy
8. **Big Idea** — the single, dominant campaign concept
9. **Campaign hypothesis** — "We believe [action] will [outcome] for [audience] because [reason]"
10. **Growth loop identification** — which loop does this campaign feed?

### Quality Gate
- Orchestrator reviews strategy output for completeness and internal consistency
- Flags unknowns and asks minimum viable follow-ups
- If offer score is low, recommend offer improvement before proceeding to content

**Output:** Complete campaign brief using the brief template, with all strategic fields populated.

## Step 3 — Asset Production

**Owner:** Content Engine (02) + Channel Operator (03)

### Content Engine produces:
1. **Long-form source draft** — Hook→Story→Offer structure, written in the customer's language
2. **Derivative snippets** — adapted (not copy-pasted) for each target channel
3. **Hook variants** (5+) — categorized by type (contrarian, specific result, curiosity gap, pattern interrupt, narrative)
4. **CTA variants** (3+) — direct, benefit-driven, risk-reversed
5. **Asset checklist** — creative, proof, and technical requirements

### Channel Operator produces:
1. **Channel-specific publishing plans** — format, cadence, sequencing per channel
2. **Channel KPIs and leading indicators** — measurable success criteria per channel
3. **Growth loop mapping** — how this campaign connects to compounding loops
4. **Cross-channel coordination** — how channels reinforce each other
5. **Testing plan** — what to A/B test per channel

### Quality Checks at This Stage
- Content Engine verifies all content matches the target awareness level
- Content Engine ensures all claims have evidence at Proof Hierarchy level 2+
- Channel Operator verifies all content is channel-native (not cross-posted)
- Both agents verify alignment with the Big Idea and messaging pillars

**Output:** Complete copy pack (using copy-pack template) + channel execution plans.

## Step 4 — QA Gate

**Owner:** Quality Guardian (04)

The Quality Guardian runs a comprehensive 8-dimension review:

1. **Template completeness** — all required fields populated
2. **Strategic alignment** — awareness level match, Big Idea coherence, messaging pillar adherence
3. **Copywriting quality** — Delete Test, Bar Stool Test, "So What?" Test, specificity check
4. **Brand voice & tone** — adherence to policies.yaml voice guidelines
5. **Proof & evidence quality** — every claim scored on the Proof Hierarchy
6. **Compliance & policy** — forbidden claims, required disclaimers, platform policies
7. **Channel format & fit** — platform-native formatting, character limits, CTA placement
8. **Remarkability score** — Seth Godin "Purple Cow" test (minimum score: 3/5)

### Failure Handling
- **Blockers:** Cannot proceed. Returns to responsible agent with specific edit requirements
- **Majors:** Should not proceed. Returns with strong recommendations
- **Minors:** Can proceed with notes for improvement
- Remarkability score below 3 triggers mandatory revision cycle
- Claims at proof level 0 are automatic blockers

**Output:** QA report (using qa-report template) with overall verdict: PASS / FAIL / PASS_WITH_NOTES.

## Step 5 — Human Approval

**Owner:** Founder (Human)

The Orchestrator presents a consolidated approval package:

### Approval Package Contents
- Campaign brief summary (Big Idea, awareness level, hypothesis)
- All content pieces with QA status
- Channel execution plans
- Model trace (which models were used and why)
- QA report (all checks with results)
- Estimated LLM cost for this campaign

### Founder Actions
- **Approve:** Move to publish prep
- **Approve with edits:** Specify changes → return to relevant agent → re-run QA on changes
- **Reject:** Provide reason → system logs rejection for learning → return to strategy or content phase
- **Score (1-5):** Quality rating for LLM Council memory

**Output:** Approval status + human score + specific feedback (if any).

## Step 6 — Memory Update

**Owner:** LLM Council (05)

### On Approval
- Append model/task/channel outcome to LLM council memory
- Record human score and any specific praise
- Boost confidence for this model on similar future tasks

### On Rejection
- Append failure reason and context to memory
- Decrease model confidence for this task/channel combination
- Analyze: was this a model issue or a prompt/strategy issue?
- Feed learning back to improve future routing

**Output:** Updated model memory records.

## Step 7 — Publish Prep

**Owner:** Channel Operator (03)

- Generate final payloads formatted for each publishing platform
- Create publishing checklists with timing and sequencing
- Prepare tracking setup (UTM parameters, conversion pixels, analytics events)
- Deliver final files/copy in platform-ready format

**Output:** Ready-to-publish assets + checklists. System does NOT auto-publish unless explicitly enabled.

## Step 8 — Postmortem

**Owner:** Orchestrator + Strategy Planner

### Postmortem triggers at campaign end (or at 2-week mark):

1. **Performance review** — actual vs. predicted KPIs per channel
2. **Hypothesis validation** — was the campaign hypothesis confirmed or rejected?
3. **Growth loop assessment** — did the campaign feed the intended loop? How well?
4. **Content analysis** — which hooks/variants performed best? Why?
5. **Model performance** — did the LLM Council make good routing decisions?
6. **Lessons learned** — specific, actionable insights for the next campaign
7. **Template/policy updates** — should any templates or policies be updated based on learnings?
8. **Next campaign recommendation** — based on what was learned, what should we do next?

**Output:** Postmortem report with actionable insights that feed back into the next cycle.

---

## Loop Summary

```
Intake → Diagnosis → LLM Selection → Strategy → Content + Channels → QA → Human Approval → Memory → Publish → Postmortem
   ↑                                                                                                              │
   └──────────────────────────────── Learnings feed back into next cycle ──────────────────────────────────────────┘
```

Every cycle makes the system smarter. The model memory improves routing. The postmortems improve strategy. The QA feedback improves content quality. This is a learning machine, not a one-shot generator.
