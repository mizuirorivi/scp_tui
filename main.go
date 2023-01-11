package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/shirou/gopsutil/v3/process"
)

var (
	tr     = tearoot{}
	models = []tea.Model{
		initialModel(),
		initialActionModel(),
	}
	views = []func(model) (string, model){
		ViewSeletPage,
		ViewSelectedSSHProcess,
	}
	manage_process = Node{
		data: Data{
			model:     models[0],
			model_num: 0,
		},
		next: nil,
		prev: nil}
)

type tickMsg time.Time
type tearoot struct {
	will_quite bool
}
type Data struct {
	model     tea.Model
	view      string
	model_num int
}
type Node struct {
	data Data
	cur  *Node
	next *Node
	prev *Node
}

type model struct {
	cursor     int
	choices    []string
	selected   map[int]struct{}
	will_quite bool
	is_no_ssh  bool
}

func (m Node) Manage_process() {

}

func (m Node) Move_next_node() *Node {
	var node Node
	if m.data.model_num+1 >= len(models) {
		return &m
	}
	node.data = Data{
		model:     models[m.data.model_num+1],
		model_num: m.data.model_num + 1,
	}
	node.next = nil
	node.prev = &m

	m.next = &node

	m = node
	return &node
}

func (m Node) Run() {
	p := tea.NewProgram(m.data.model)
	p.Run()
}

// func (m Node) Update {
// }

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

func initialModel() model {
	choise := get_process()
	is_no_ssh := false
	if len(choise) == 0 {
		choise = []string{"no ssh process"}
		is_no_ssh = true
	}
	return model{
		choices:   choise,
		is_no_ssh: is_no_ssh,
		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}
func initialActionModel() model {
	tr.will_quite = false
	return model{
		choices: []string{"remote_to_local", "local_to_remote", "show_files"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tick(), tea.EnterAltScreen)
}
func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			manage_process = *(manage_process.Move_next_node())
			manage_process.Run()
		}
	}

	if tr.will_quite {
		fmt.Println("is will quite")
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() string {
	view, model := views[manage_process.data.model_num](m)
	m = model
	return view
}

func ViewSeletPage(m model) (string, model) {
	s := "which do you select ssh process?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			// checked = "x"
			manage_process.Move_next_node()
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s, m
}

func ViewSelectedSSHProcess(m model) (string, model) {
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

func main() {
	manage_process.Run()
}
