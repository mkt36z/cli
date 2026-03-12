package local

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// PreCheckResult holds the results of a local QA pre-check.
type PreCheckResult struct {
	Pass     bool
	Warnings []string
	Errors   []string
	Score    int // 0-100
}

// ForbiddenWords that should not appear in marketing copy.
var ForbiddenWords = []string{
	"synergy", "leverage", "paradigm", "disrupt",
	"revolutionary", "game-changing", "world-class",
	"best-in-class", "cutting-edge", "next-generation",
}

// PreCheck runs local QA validation on generated content.
func PreCheck(content string) PreCheckResult {
	result := PreCheckResult{Score: 100}

	// Word count check
	words := strings.Fields(content)
	wordCount := len(words)

	if wordCount < 5 {
		result.Errors = append(result.Errors, fmt.Sprintf("Too short: %d words (minimum 5)", wordCount))
		result.Score -= 30
	} else if wordCount > 5000 {
		result.Warnings = append(result.Warnings, fmt.Sprintf("Very long: %d words", wordCount))
		result.Score -= 10
	}

	// Readability: average sentence length
	sentences := splitSentences(content)
	if len(sentences) > 0 {
		avgWords := wordCount / len(sentences)
		if avgWords > 30 {
			result.Warnings = append(result.Warnings,
				fmt.Sprintf("Long sentences: avg %d words/sentence (aim for <25)", avgWords))
			result.Score -= 10
		}
	}

	// Forbidden words
	lower := strings.ToLower(content)
	for _, word := range ForbiddenWords {
		if strings.Contains(lower, word) {
			result.Warnings = append(result.Warnings,
				fmt.Sprintf("Buzzword detected: %q — consider more specific language", word))
			result.Score -= 5
		}
	}

	// Structure validation: check for headings in longer content
	if wordCount > 200 && !strings.Contains(content, "#") && !strings.Contains(content, "**") {
		result.Warnings = append(result.Warnings, "Long content without headings or bold text — consider adding structure")
		result.Score -= 5
	}

	// UTF-8 validation
	if !utf8.ValidString(content) {
		result.Errors = append(result.Errors, "Content contains invalid UTF-8 characters")
		result.Score -= 20
	}

	// Clamp score
	if result.Score < 0 {
		result.Score = 0
	}
	result.Pass = len(result.Errors) == 0 && result.Score >= 50

	return result
}

var sentenceSplitter = regexp.MustCompile(`[.!?]+\s+`)

func splitSentences(text string) []string {
	parts := sentenceSplitter.Split(text, -1)
	var sentences []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			sentences = append(sentences, p)
		}
	}
	return sentences
}
