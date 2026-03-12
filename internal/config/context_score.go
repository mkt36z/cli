package config

// contextWeights defines how much each context dimension impacts output quality.
// Total weights sum to 100. Ordered by impact on generated content.
var contextWeights = map[ContextType]int{
	ContextBusiness:     10, // Foundation — needed for everything
	ContextProduct:      15, // Core — value prop drives all copy
	ContextICP:          15, // Critical — who we're writing for
	ContextCompetitive:  8,  // Positioning context
	ContextBrandVoice:   12, // Voice consistency
	ContextChannels:     5,  // Channel optimization
	ContextContentSEO:   5,  // SEO integration
	ContextSales:        5,  // Conversion context
	ContextGoals:        8,  // Strategic alignment
	ContextLanguageBank: 10, // Customer voice authenticity
	ContextProofStack:   5,  // Social proof density
	ContextAntiPatterns: 2,  // Guardrails
}

// ContextScore calculates a 0–100 score based on which context files are populated.
// Weighted by impact on output quality — Product (15) matters more than Anti-Patterns (2).
func ContextScore() int {
	score := 0
	for _, ct := range AllContextTypes {
		if ContextFileExists(ct) {
			score += contextWeights[ct]
		}
	}
	return score
}

// ContextScoreBreakdown returns per-dimension scores for detailed display.
type ScoreBreakdown struct {
	Type      ContextType `json:"type"`
	Label     string      `json:"label"`
	MaxPoints int         `json:"max_points"`
	Earned    int         `json:"earned"`
	Filled    bool        `json:"filled"`
}

// Breakdown returns the detailed score breakdown for all 12 dimensions.
func Breakdown() []ScoreBreakdown {
	result := make([]ScoreBreakdown, 0, len(AllContextTypes))
	for _, ct := range AllContextTypes {
		filled := ContextFileExists(ct)
		earned := 0
		if filled {
			earned = contextWeights[ct]
		}
		result = append(result, ScoreBreakdown{
			Type:      ct,
			Label:     ct.Label(),
			MaxPoints: contextWeights[ct],
			Earned:    earned,
			Filled:    filled,
		})
	}
	return result
}

// HighestImpactMissing returns the unfilled context type with the highest weight.
// Returns empty string if all are filled.
func HighestImpactMissing() ContextType {
	var best ContextType
	bestWeight := 0
	for _, ct := range AllContextTypes {
		if !ContextFileExists(ct) && contextWeights[ct] > bestWeight {
			best = ct
			bestWeight = contextWeights[ct]
		}
	}
	return best
}
