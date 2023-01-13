package view

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	mp "github.com/mizuirorivi/scp_tui/internal/tui/manage_process"
	tm "github.com/mizuirorivi/scp_tui/internal/tui/model"
)

var (
	Views = []func(tm.Model) (string, tm.Model){
		ViewSelectPage,
		ViewSelectedSSHProcess,
	}
)

func init() {
	Views = []func(tm.Model) (string, tm.Model){
		ViewSelectPage,
		ViewSelectedSSHProcess,
	}
}

func (m tm.Model) View() string {
	view, model := Views[mp.Manage_process.Ndata.Dmodel_num](m)
	m = model
	return view
}

func ViewSelectPage(m tm.Model) (string, tm.Model) {
	s := "which do you select ssh process?\n\n"

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.Selected[i]; ok {
			// checked = "x"
			mp.Manage_process.Move_next_node()
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s, m
}

func ViewSelectedSSHProcess(m tm.Model) (string, tm.Model) {
	s := "You selected:\n\n"

	for i, choice := range m.Choices {
		if _, ok := m.Selected[i]; ok {
			s += fmt.Sprintf("  %s\n", choice)
		}
	}
	s += "which do you select action?\n\n"

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.Selected[i]; ok {
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
