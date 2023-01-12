package manage_process

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mizuirorivi/scp_tui/internal/tui/model"
)

var (
	Manage_process = Node{
		data: Data{
			model:     model.Models[0],
			model_num: 0,
		},
		next: nil,
		prev: nil}
)

func init() {
	Manage_process = Node{
		data: Data{
			model:     model.Models[0],
			model_num: 0,
		},
		next: nil,
		prev: nil}

}

type Data struct {
	model     model.Model
	view      string
	Model_num int
}

type Node struct {
	data Data
	cur  *Node
	next *Node
	prev *Node
}

func init() {

}
func (m Node) Manage_process() {

}
func (m Node) Move_next_node() *Node {
	var node Node
	if m.data.Model_num+1 >= len(model.Models) {
		return &m
	}

	node.data = Data{
		Model:     model.Models[m.data.Model_num+1],
		Model_num: m.data.Model_num + 1,
	}
	node.next = nil
	node.prev = &m

	m.next = &node

	m = node
	return &node
}

func (m Node) Run() {
	p := tea.NewProgram(tea.Model(m.data.Model))
	p.Run()
}

// func (m Node) Update {
// }
