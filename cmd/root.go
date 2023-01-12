package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mizuirorivi/mizuirorivi/scp_tui/internal/tui/model"
	"github.com/mizuirorivi/scp_tui/internal/tui/manage_process"

	"github.com/mizuirorivi/mizuirorivi/scp_tui/internal/tui/manage_process"
	"github.com/shirou/gopsutil/v3/process"
)

type tickMsg time.Time

type tearoot struct {
	will_quite bool
}

func get_process() []string {
	var ssh_process_list []string
	processes, err := process.Processes()
	if err != nil {
		fmt.Errorf("Error: %v", err)
		os.Exit(1)
	}

	for _, process := range processes {
		nme, _ := process.Cmdline()
		if strings.Contains(nme, "ssh") && strings.Contains(nme, "@") {
			ssh_process_list = append(ssh_process_list, nme)
		}
	}
	return ssh_process_list
}

func (m model.Model) Init() tea.Cmd {
	return tea.Batch(tick(), tea.EnterAltScreen)
}
func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
func (m model.Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			manage_process.Manage_process = *(manage_process.Manage_process.Move_next_node())
			manage_process.Manage_process.Run()
		}
	}

	if Tr.will_quite {
		fmt.Println("is will quite")
		return m, tea.Quit
	}

	return m, nil
}
