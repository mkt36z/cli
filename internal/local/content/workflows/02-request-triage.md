# Workflow: Request Triage

## Purpose
The intelligent routing layer that ensures every founder request gets handled with the right level of strategic depth. Not every request needs the full marketing loop — but every request needs strategic thinking.

## Decision Tree

### Gate 1: What type of request is this?

```
Founder Request
      │
      ├── Strategy (positioning, offers, audience, planning)
      │       → Route to Strategy Planner (01) → Full strategy artifacts
      │
      ├── Execution (write copy, create ads, build emails)
      │       → Check: Is strategy foundation in place? (see Gate 2)
      │
      ├── Audit (review existing content, QA check, compliance)
      │       → Route to Quality Guardian (04) → Full QA report
      │
      ├── Analysis (campaign performance, postmortem, market research)
      │       → Route to Strategy Planner (01) → Analysis + recommendations
      │
      └── Mixed / Ambiguous
              → Orchestrator asks clarifying questions (max 3)
              → Re-triage after clarification
```

### Gate 2: Is the strategic foundation in place?

Before any execution work, verify these exist:

| Foundation Element | Status Check | If Missing |
|-------------------|-------------|-----------|
| ICP + Empathy Map | Defined with psychographics and language? | → Strategy Planner builds it first |
| Awareness Level | Assessed for this audience? | → Strategy Planner assesses |
| Market Sophistication | Evaluated? | → Strategy Planner evaluates |
| Positioning | April Dunford framework completed? | → Strategy Planner creates positioning |
| Offer | Scored on Hormozi Value Equation? | → Strategy Planner evaluates offer |
| Brand Voice | Configured in policies.yaml? | → Founder must configure |

**If 2+ elements are missing:** Route to full strategy loop before execution.
**If 1 element is missing:** Can proceed with flagged gap — Orchestrator notes the risk.
**If all present:** Proceed to execution.

### Gate 3: What is the channel scope?

```
Single Channel Request
      │
      ├── Known channel with existing strategy
      │       → Focused path: Content Engine + Channel Operator + QA Gate
      │       → Faster turnaround, fewer agents involved
      │
      └── Unknown channel or new channel
              → Channel Operator evaluates fit first
              → Then focused execution path

Multi-Channel Request
      │
      └── Full workflow: Strategy → Content → All Channels → QA → Approval
              → Cross-channel coordination required
              → Channel Operator ensures consistency + mutual reinforcement
```

### Gate 4: New campaign or optimization?

```
New Campaign
      │
      └── Full marketing loop (Workflow 01)
              → All steps: intake through postmortem
              → Campaign hypothesis required

Optimization of Existing Campaign
      │
      ├── Performance data available?
      │       ├── YES → Analysis first → identify what to change → targeted revision
      │       └── NO → Flag: need performance data to optimize intelligently
      │
      └── What kind of optimization?
              ├── Hook/headline testing → Content Engine generates new variants
              ├── Channel expansion → Channel Operator evaluates + plans
              ├── Audience refinement → Strategy Planner updates ICP
              ├── Offer improvement → Strategy Planner re-scores Value Equation
              └── Full pivot → Treat as new campaign
```

### Gate 5: Which LLM profile is required?

| Request Type | LLM Profile | Rationale |
|-------------|-------------|-----------|
| Strategy, positioning, research | Deterministic | Needs accuracy, structure, factual grounding |
| Hook generation, email copy, social posts | Creative | Needs linguistic flair, variety, emotional pull |
| QA review, compliance check | Audit | Needs precision, consistency, rule-following |
| Channel reformatting, adaptation | Adaptation | Needs format compliance, message preservation |
| Mixed (strategy + creative) | Multiple | Council selects different models for each phase |

**Rule:** Any generation request → LLM council model selection happens FIRST.

## Routing Summary

| Scenario | Path | Estimated Time |
|----------|------|---------------|
| Full campaign from scratch | Full loop (all 8 steps) | 30-60 min active + review |
| Single-channel content piece | Focused: Content + Channel + QA | 15-20 min active + review |
| Strategy refresh | Strategy Planner → Orchestrator review | 15-20 min active |
| Content audit / QA check | Quality Guardian analysis | 10-15 min |
| Campaign postmortem | Analysis → Strategy recommendations | 15-20 min |
| Quick hook/variant generation | Content Engine → QA quick check | 5-10 min |
| Missing context | Clarifying questions → re-triage | 5 min + founder response |

## Anti-Patterns the Triage Must Prevent

1. **Execution without strategy:** "Write me an email" → Check: do we know the audience, awareness level, offer, and Big Idea?
2. **Strategy without action:** Endless planning → Every strategy output must include concrete next steps.
3. **Channel mismatch:** Recommending paid ads to a pre-revenue founder → Hormozi Matrix check.
4. **Vanity campaigns:** "Let's go viral" → Redirect to measurable hypothesis-driven campaign.
5. **Copy-paste distribution:** Same content across all channels → Enforce channel-native adaptation.
