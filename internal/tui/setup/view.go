package setup

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	if m.step == StepDone {
		return "Setup complete!\n\n" +
			"You can now use `gofm` for more information.\n\n" +
			"Press any key to exit..."
	}

	var b strings.Builder

	b.WriteString("Welcome to GoFM\n\n")

	switch m.step {
	case StepUsername:
		b.WriteString("Please enter your Last.fm username\n\n")
		b.WriteString(m.username.View())
		b.WriteString("\n\n")
		b.WriteString("Press Enter to continue")
	case StepAPIKey:
		b.WriteString("Please enter your Last.fm API key\n\n")
		b.WriteString(m.apiKey.View())
		b.WriteString("\n\n")
		b.WriteString("Press Enter to validate and save")
	case StepSaving:
		b.WriteString("Saving configuration...\n\n")
	}

	if m.err != nil {
		b.WriteString(fmt.Sprintf("\n\nError: %s", m.err.Error()))
	}

	return b.String()
}
