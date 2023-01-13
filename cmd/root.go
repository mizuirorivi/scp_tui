package cmd

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	mp "github.com/mizuirorivi/mizuirorivi/scp_tui/internal/tui/manage_process"
	tm "github.com/mizuirorivi/mizuirorivi/scp_tui/internal/tui/model"
)

type tickMsg time.Time

type tearoot struct {
	Will_quite bool
}

func (m tm.Model) Init() tea.Cmd {
	return tea.Batch(tick(), tea.EnterAltScreen)
}
func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
func (m tm.Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			fmt.Println("bye-bye")
			os.Exit(0)
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			// _, ok := m.selected[m.cursor]
			// if ok {
			//   delete(m.selected, m.cursor)
			// } else {
			//   m.selected[m.cursor] = struct{}{}
			// }
			if m.is_no_ssh {
				break
			}
			mp.Manage_process = *(mp.Manage_process.Move_next_node())
			mp.Manage_process.Run()
		}
	}

	if Tr.Will_quite {
		fmt.Println("is will quite")
		return m, tea.Quit
	}

	return m, nil
}
