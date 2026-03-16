package setup

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/theOldZoom/gofm/internal/config"
)

type step int

const (
	StepUsername step = iota
	StepAPIKey
	StepSaving
	StepDone
)

type errMsg struct {
	err error
}

type doneMsg struct {
	cfg *config.Config
}

type Model struct {
	step     step
	username textinput.Model
	apiKey   textinput.Model
	focused  int
	err      error
	width    int
	height   int
	saving   bool
	result   *config.Config
}

func NewModel() Model {
	username := textinput.New()
	username.Placeholder = "Last.fm username"
	username.Focus()
	username.CharLimit = 96
	username.Width = 40

	apiKey := textinput.New()
	apiKey.Placeholder = "Last.fm API key"
	apiKey.CharLimit = 64
	apiKey.Width = 40

	return Model{
		step:     StepUsername,
		username: username,
		apiKey:   apiKey,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case errMsg:
		m.err = msg.err
		m.saving = false
		m.step = StepAPIKey
		m.apiKey.SetValue("")
		m.apiKey.Focus()
		return m, nil
	case doneMsg:
		m.result = msg.cfg
		m.err = nil
		m.saving = false
		m.step = StepDone
		return m, nil
	case tea.KeyMsg:
		if m.step == StepDone {
			return m, tea.Quit
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			switch m.step {
			case StepUsername:
				if strings.TrimSpace(m.username.Value()) == "" {
					m.err = fmt.Errorf("username is required")
					return m, nil
				}
				m.err = nil
				m.step = StepAPIKey
				m.username.Blur()
				m.apiKey.Focus()
				return m, nil
			case StepAPIKey:
				if strings.TrimSpace(m.apiKey.Value()) == "" {
					m.err = fmt.Errorf("API key is required")
					return m, nil
				}
				m.err = nil
				m.step = StepSaving
				m.saving = true
				m.apiKey.Blur()
				username := strings.TrimSpace(m.username.Value())
				apiKey := strings.TrimSpace(m.apiKey.Value())
				return m, submitSetup(username, apiKey)
			}
		}
	}

	var cmd tea.Cmd
	switch m.step {
	case StepUsername:
		m.username, cmd = m.username.Update(msg)
	case StepAPIKey:
		m.apiKey, cmd = m.apiKey.Update(msg)
	}
	return m, cmd
}

func (m Model) Result() (*config.Config, error) {
	return m.result, m.err
}

func submitSetup(username, apiKey string) tea.Cmd {
	return func() tea.Msg {
		if err := config.ValidateAPIKey(apiKey); err != nil {
			return errMsg{err: err}
		}

		if err := config.ValidateUsername(username, apiKey); err != nil {
			return errMsg{err: err}
		}

		cfg := &config.Config{
			Username: username,
			ApiKey:   apiKey,
		}

		if err := config.Save(cfg); err != nil {
			return errMsg{err: err}
		}

		return doneMsg{cfg: cfg}
	}
}
