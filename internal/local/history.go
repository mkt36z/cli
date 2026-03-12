package local

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mkt36z/cli/internal/config"
)

// HistoryEntry records a generation event for local history.
type HistoryEntry struct {
	Timestamp string `json:"timestamp"`
	Command   string `json:"command"`
	Input     string `json:"input"`
	Output    string `json:"output"`
	Model     string `json:"model,omitempty"`
	Tokens    int    `json:"tokens,omitempty"`
	Duration  string `json:"duration,omitempty"`
	QAScore   int    `json:"qa_score,omitempty"`
}

// AppendHistory appends a generation entry to the local history file.
func AppendHistory(entry HistoryEntry) error {
	path := config.HistoryFilePath()

	if err := config.EnsureDir(config.CacheDir()); err != nil {
		return fmt.Errorf("creating cache directory: %w", err)
	}

	if entry.Timestamp == "" {
		entry.Timestamp = time.Now().UTC().Format(time.RFC3339)
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("opening history file: %w", err)
	}
	defer f.Close()

	_, err = f.Write(append(data, '\n'))
	return err
}

// ReadHistory reads the last n entries from the history file.
// If n <= 0, returns all entries.
func ReadHistory(n int) ([]HistoryEntry, error) {
	path := config.HistoryFilePath()

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("opening history file: %w", err)
	}
	defer f.Close()

	var entries []HistoryEntry
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var entry HistoryEntry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			continue // skip malformed lines
		}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return entries, err
	}

	// Return last n entries
	if n > 0 && len(entries) > n {
		entries = entries[len(entries)-n:]
	}

	return entries, nil
}
