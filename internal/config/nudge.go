package config

import "fmt"

// Nudge represents a contextual, non-blocking tip shown to the user.
type Nudge struct {
	Message  string `json:"message"`
	Command  string `json:"command"`
	Priority int    `json:"priority"` // higher = more important
}

// GetNudge returns the most relevant nudge based on current context state.
// Returns nil if no nudge is warranted (all context filled or score >= 80).
func GetNudge() *Nudge {
	score := ContextScore()

	// No nudge if context is well-populated.
	if score >= 80 {
		return nil
	}

	populated := len(PopulatedContextFiles())

	// First-run nudge.
	if populated == 0 {
		return &Nudge{
			Message:  "Set up your brand context for personalized output",
			Command:  "mkt36z init",
			Priority: 100,
		}
	}

	// Nudge for highest-impact missing context.
	missing := HighestImpactMissing()
	if missing == "" {
		return nil
	}

	return &Nudge{
		Message: fmt.Sprintf("Fill %s for better results (score: %d/100)", missing.Label(), score),
		Command: fmt.Sprintf("mkt36z config context edit %s", string(missing)),
		Priority: contextWeights[missing],
	}
}
