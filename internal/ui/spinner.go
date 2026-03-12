package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Spinner wraps a Bubble Tea spinner for consistent CLI progress indication.
// Automatically falls back to static text when not in a TTY.
type Spinner struct {
	program *tea.Program
	model   *spinnerModel
}

type spinnerModel struct {
	spinner spinner.Model
	message string
	done    bool
	final   string
}

type stopMsg struct{ final string }
type updateSpinnerMsg struct{ message string }

func (m *spinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case stopMsg:
		m.done = true
		m.final = msg.final
		return m, tea.Quit
	case updateSpinnerMsg:
		m.message = msg.message
		return m, nil
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			m.done = true
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m *spinnerModel) View() string {
	if m.done {
		if m.final != "" {
			return m.final + "\n"
		}
		return ""
	}
	return m.spinner.View() + " " + m.message + "\n"
}

// NewSpinner creates and starts a spinner with the given message.
// If stdout is not a TTY, prints the message as static text.
func NewSpinner(message string) *Spinner {
	if !IsTTY() {
		fmt.Fprintln(os.Stderr, message)
		return &Spinner{}
	}

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(ColorBrandBlue)

	m := &spinnerModel{
		spinner: s,
		message: message,
	}

	p := tea.NewProgram(m, tea.WithOutput(os.Stderr))
	sp := &Spinner{program: p, model: m}

	go p.Run() //nolint:errcheck
	return sp
}

// UpdateMessage changes the spinner's display message while running.
func (s *Spinner) UpdateMessage(message string) {
	if s.program == nil {
		fmt.Fprintln(os.Stderr, message)
		return
	}
	s.program.Send(updateSpinnerMsg{message: message})
}

// Stop halts the spinner and shows a final message.
func (s *Spinner) Stop(finalMsg string) {
	if s.program == nil {
		if finalMsg != "" {
			fmt.Fprintln(os.Stderr, finalMsg)
		}
		return
	}
	s.program.Send(stopMsg{final: finalMsg})
	s.program.Wait()
}
