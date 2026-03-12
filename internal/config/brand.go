package config

// BrandVoice defines the brand's voice and identity for content generation.
type BrandVoice struct {
	// Core identity
	CompanyName string `yaml:"company_name,omitempty" json:"company_name,omitempty"`
	Industry    string `yaml:"industry,omitempty" json:"industry,omitempty"`
	Archetype   string `yaml:"archetype,omitempty" json:"archetype,omitempty"` // one of 12 Jungian archetypes

	// Voice definition
	Tone           []string `yaml:"tone,omitempty" json:"tone,omitempty"`             // e.g. ["confident", "warm", "direct"]
	Personality    []string `yaml:"personality,omitempty" json:"personality,omitempty"` // e.g. ["expert", "approachable"]
	ForbiddenWords []string `yaml:"forbidden_words,omitempty" json:"forbidden_words,omitempty"`

	// Content-type specific tone overrides
	ContentTones map[string][]string `yaml:"content_tones,omitempty" json:"content_tones,omitempty"` // e.g. {"email": ["casual"], "landing-page": ["urgent"]}

	// Brand values
	Values       []string `yaml:"values,omitempty" json:"values,omitempty"`
	Differentiator string `yaml:"differentiator,omitempty" json:"differentiator,omitempty"`
}
