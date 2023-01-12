package view

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mizuirorivi/scp_tui/internal/tui/manage_process"
	"github.com/mizuirorivi/scp_tui/internal/tui/model"
)

var (
	Views = []func(model.Model) (string, model.Model){
		ViewSeletPage,
		ViewSelectedSSHProcess,
	}
)

func init() {
	Views = []func(model.Model) (string, model.Model){
		ViewSeletPage,
		ViewSelectedSSHProcess,
	}
}

func (m model.Model) View() string {
	view, model := Views[manage_process.Manage_process.data.Model_num](m)
	m = model
	return view
}

func ViewSeletPage(m model.Model) (string, model.Model) {
	s := "which do you select ssh process?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			// checked = "x"
			Manage_process.Move_next_node()
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s, m
}

func ViewSelectedSSHProcess(m model.Model) (string, model.Model) {
	s := "You selected:\n\n"

	for i, choice := range m.choices {
		if _, ok := m.selected[i]; ok {
			s += fmt.Sprintf("  %s\n", choice)
		}
	}
	s += "which do you select action?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			// checked = "x"
			// tr.will_quite = true
			// os.Exit(1)
			tea.Quit()
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s, m
}
