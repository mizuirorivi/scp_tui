package model

import (
	"github.com/mizuirorivi/scp_tui/internal/cmd"
)

type Model struct {
	cursor     int
	choices    []string
	selected   map[int]struct{}
	will_quite bool
	is_no_ssh  bool
}

var Models []Model

func init() {
	Models = []Model{
		initialModel(),
		initialActionModel(),
	}
}

func initialModel() Model {
	choise := cmd.get_process()
	is_no_ssh := false
	if len(choise) == 0 {
		choise = []string{"no ssh process"}
		is_no_ssh = true
	}
	return Model{
		choices:   choise,
		is_no_ssh: is_no_ssh,
		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}
func initialActionModel() Model {
	cmd.Tr.will_quite = false
	return Model{
		choices: []string{"remote_to_local", "local_to_remote", "show_files"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}
