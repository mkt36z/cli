package api

import (
	"bufio"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// Event represents a Server-Sent Event from the Workers API.
type Event struct {
	// Type is the SSE event type (e.g. "agent_start", "content_chunk", "error", "done").
	Type string
	// Data is the JSON payload of the event.
	Data string
	// ID is the optional event ID for reconnection.
	ID string
}

// SSE event type constants.
const (
	EventAgentStart      = "agent_start"
	EventAgentProgress   = "agent_progress"
	EventAgentOutput     = "agent_output"
	EventAgentComplete   = "agent_complete"
	EventContentChunk    = "content_chunk"
	EventQAScore         = "qa_score"
	EventError           = "error"
	EventDone            = "done"
	EventSEOAudit        = "seo_audit"
	EventEmailAnalysis   = "email_analysis"
	EventCompetitiveMap  = "competitive_map"
	EventRevenueDiag     = "revenue_diagnostic"
	EventBrandHealth     = "brand_health"
	EventGovernanceAlert = "governance_alert"
)

// StreamSSE reads an SSE stream from resp and sends parsed events on the returned channel.
// The channel is closed when the stream ends, ctx is cancelled, or an error occurs.
// The caller must have already verified resp.StatusCode == 200.
func StreamSSE(ctx context.Context, resp *http.Response) <-chan Event {
	ch := make(chan Event, 16)

	go func() {
		defer close(ch)
		defer resp.Body.Close()

		scanner := bufio.NewScanner(resp.Body)
		// Allow up to 1MB per SSE line
		scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

		var event Event

		for scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			default:
			}

			line := scanner.Text()

			// Empty line = dispatch event
			if line == "" {
				if event.Type != "" || event.Data != "" {
					if event.Type == "" {
						event.Type = "message"
					}
					ch <- event
					event = Event{}
				}
				continue
			}

			// Parse SSE fields
			if strings.HasPrefix(line, "event:") {
				event.Type = strings.TrimSpace(line[6:])
			} else if strings.HasPrefix(line, "data:") {
				data := strings.TrimSpace(line[5:])
				if event.Data != "" {
					event.Data += "\n" + data
				} else {
					event.Data = data
				}
			} else if strings.HasPrefix(line, "id:") {
				event.ID = strings.TrimSpace(line[3:])
			}
			// Ignore "retry:" and comments (lines starting with ":")
		}

		// Flush any remaining event
		if event.Type != "" || event.Data != "" {
			if event.Type == "" {
				event.Type = "message"
			}
			ch <- event
		}

		if err := scanner.Err(); err != nil {
			// SECURITY (VULN-15): Use json.Marshal for proper escaping instead of
			// fmt.Sprintf, which produces malformed JSON if the error contains quotes.
			errMsg := "stream read error"
			errJSON, _ := json.Marshal(map[string]string{"error": errMsg})
			ch <- Event{
				Type: EventError,
				Data: string(errJSON),
			}
		}
	}()

	return ch
}
