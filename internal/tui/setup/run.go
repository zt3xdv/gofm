package setup

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/theOldZoom/gofm/internal/config"
)

func Run() (*config.Config, error) {
	m := NewModel()

	p := tea.NewProgram(m, tea.WithAltScreen())
	finalModel, err := p.Run()
	if err != nil {
		return nil, err
	}

	result := finalModel.(Model)
	cfg, err := result.Result()
	if err != nil {
		return nil, err
	}
	if cfg == nil {
		return nil, errors.New("setup cancelled")
	}

	return cfg, nil
}
