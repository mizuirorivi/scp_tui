package manage_process

import (
	tea "github.com/charmbracelet/bubbletea"
	tm "github.com/mizuirorivi/scp_tui/internal/tui/model"
)

var (
	Manage_process = Node{
		Ndata: Data{
			Dmodel:     tm.Models[0],
			Dmodel_num: 0,
		},
		next: nil,
		prev: nil}
)

func init() {
	Manage_process = Node{
		Ndata: Data{
			Dmodel:     tm.Models[0],
			Dmodel_num: 0,
		},
		next: nil,
		prev: nil}
}

type Data struct {
	Dmodel     tm.Model
	Dview      string
	Dmodel_num int
}

type Node struct {
	Ndata Data
	cur   *Node
	next  *Node
	prev  *Node
}

func init() {

}

func (m Node) Manage_process() {

}

func (m Node) Move_next_node() *Node {
	var node Node
	if m.Ndata.Dmodel_num+1 >= len(tm.Models) {
		return &m
	}

	node.Ndata = Data{
		Dmodel:     tm.Models[m.Ndata.Dmodel_num+1],
		Dmodel_num: m.Ndata.Dmodel_num + 1,
	}
	node.next = nil
	node.prev = &m

	m.next = &node

	m = node
	return &node
}

func (m Node) Run() {
	p := tea.NewProgram(m.Ndata.Dmodel)
	p.Run()
}

// func (m Node) Update {
// }
