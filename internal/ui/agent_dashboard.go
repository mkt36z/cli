package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// AgentStatus represents the status of one agent in the dashboard.
type AgentStatus struct {
	Name        string
	Label       string
	Status      AgentRunStatus
	Message     string
	StartedAt   time.Time
	CompletedAt time.Time
	Expanded    bool
	Output      string
	QAScore     int
}

// AgentRunStatus represents the state of an agent run.
type AgentRunStatus int

const (
	AgentPending AgentRunStatus = iota
	AgentRunning
	AgentComplete
	AgentFailed
	AgentSkipped
)

// AgentUpdateMsg is sent when an agent's status changes.
type AgentUpdateMsg struct {
	Name    string
	Status  AgentRunStatus
	Message string
	Output  string
	QAScore int
}

// AgentDoneMsg signals all agents have completed.
type AgentDoneMsg struct{}

// DashboardModel is a Bubble Tea model for the multi-agent dashboard.
type DashboardModel struct {
	agents   []AgentStatus
	cursor   int
	title    string
	done     bool
	err      error
	width    int
	height   int
	quitting bool
}

// NewDashboardModel creates a new agent dashboard.
func NewDashboardModel(title string, agentNames []AgentInfo) DashboardModel {
	agents := make([]AgentStatus, len(agentNames))
	for i, info := range agentNames {
		agents[i] = AgentStatus{
			Name:   info.Name,
			Label:  info.Label,
			Status: AgentPending,
		}
	}
	return DashboardModel{
		agents: agents,
		title:  title,
	}
}

// AgentInfo holds name and label for an agent.
type AgentInfo struct {
	Name  string
	Label string
}

func (m DashboardModel) Init() tea.Cmd {
	return nil
}

func (m DashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.agents)-1 {
				m.cursor++
			}
		case "enter", " ":
			if m.cursor < len(m.agents) {
				m.agents[m.cursor].Expanded = !m.agents[m.cursor].Expanded
			}
		}

	case AgentUpdateMsg:
		for i := range m.agents {
			if m.agents[i].Name == msg.Name {
				m.agents[i].Message = msg.Message
				if msg.Status != 0 {
					m.agents[i].Status = msg.Status
				}
				if msg.Status == AgentRunning && m.agents[i].StartedAt.IsZero() {
					m.agents[i].StartedAt = time.Now()
				}
				if msg.Status == AgentComplete || msg.Status == AgentFailed {
					m.agents[i].CompletedAt = time.Now()
				}
				if msg.Output != "" {
					m.agents[i].Output = msg.Output
				}
				if msg.QAScore > 0 {
					m.agents[i].QAScore = msg.QAScore
				}
				break
			}
		}

	case AgentDoneMsg:
		m.done = true
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m DashboardModel) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder

	// Title
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorBrandBlue).
		MarginBottom(1)
	b.WriteString(titleStyle.Render(m.title))
	b.WriteString("\n\n")

	// Agent rows
	for i, agent := range m.agents {
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}

		icon := statusIcon(agent.Status)
		name := fmt.Sprintf("%-22s", agent.Label)
		message := agent.Message

		// Duration for completed agents
		duration := ""
		if !agent.CompletedAt.IsZero() && !agent.StartedAt.IsZero() {
			d := agent.CompletedAt.Sub(agent.StartedAt)
			duration = fmt.Sprintf(" (%s)", d.Round(time.Millisecond))
		} else if agent.Status == AgentRunning && !agent.StartedAt.IsZero() {
			d := time.Since(agent.StartedAt)
			duration = fmt.Sprintf(" (%s...)", d.Round(time.Millisecond))
		}

		// QA score
		qa := ""
		if agent.QAScore > 0 {
			qa = fmt.Sprintf(" [QA: %d/100]", agent.QAScore)
		}

		row := fmt.Sprintf("%s%s %s %s%s%s", cursor, icon, name, message, duration, qa)
		b.WriteString(row)
		b.WriteString("\n")

		// Expanded output
		if agent.Expanded && agent.Output != "" {
			lines := strings.Split(agent.Output, "\n")
			maxLines := 10
			if len(lines) > maxLines {
				lines = lines[:maxLines]
				lines = append(lines, "    ...")
			}
			for _, line := range lines {
				b.WriteString("    " + line + "\n")
			}
		}
	}

	// Summary
	b.WriteString("\n")
	completed := 0
	failed := 0
	running := 0
	for _, a := range m.agents {
		switch a.Status {
		case AgentComplete:
			completed++
		case AgentFailed:
			failed++
		case AgentRunning:
			running++
		}
	}

	summary := fmt.Sprintf("%d/%d complete", completed, len(m.agents))
	if running > 0 {
		summary += fmt.Sprintf(", %d running", running)
	}
	if failed > 0 {
		summary += fmt.Sprintf(", %d failed", failed)
	}
	b.WriteString(Dim.Render(summary))

	if !m.done {
		b.WriteString(Dim.Render("  [j/k] navigate  [enter] expand  [q] quit"))
	}

	return b.String()
}

// Done returns true if all agents have completed.
func (m DashboardModel) Done() bool {
	return m.done
}

// Agents returns the current agent statuses.
func (m DashboardModel) Agents() []AgentStatus {
	return m.agents
}

func statusIcon(s AgentRunStatus) string {
	switch s {
	case AgentPending:
		return "○"
	case AgentRunning:
		return "◎"
	case AgentComplete:
		return "✓"
	case AgentFailed:
		return "✗"
	case AgentSkipped:
		return "-"
	default:
		return "?"
	}
}
