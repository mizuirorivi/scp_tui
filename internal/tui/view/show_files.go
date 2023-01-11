package view

import (
	"fmt"
	"github.com/mizuirorivi/scp_tui/internal/tui/view/show_files"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m show_files.mainModel) Init() tea.Cmd { // start the timer and spinner on program start return tea.Batch(m.timer.Init(), m.spinner.Tick)

	return tea.Batch(m.timer.Init(), m.spinner.Tick)
}

func (m show_files.mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.state == show_files.timerView {
				m.state = show_files.spinnerView
			} else {
				m.state = show_files.timerView
			}
		case "n":
			if m.state == show_files.timerView {
				m.timer = timer.New(show_files.defaultTime)
				cmds = append(cmds, m.timer.Init())
			} else {
				m.Next()
				m.resetSpinner()
				cmds = append(cmds, spinner.Tick)
			}
		}
		switch m.state {
		// update whichever model is focused
		case show_files.spinnerView:
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		default:
			m.timer, cmd = m.timer.Update(msg)
			cmds = append(cmds, cmd)
		}
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	case timer.TickMsg:
		m.timer, cmd = m.timer.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m show_files.mainModel) View() string {
	var s string
	model := m.currentFocusedModel()
	if m.state == show_files.timerView {
		s += lipgloss.JoinHorizontal(lipgloss.Top, show_files.focusedModelStyle.Render(fmt.Sprintf("%4s", m.timer.View())), show_files.modelStyle.Render(m.spinner.View()))
	} else {
		s += lipgloss.JoinHorizontal(lipgloss.Top, show_files.modelStyle.Render(fmt.Sprintf("%4s", m.timer.View())), show_files.focusedModelStyle.Render(m.spinner.View()))
	}
	s += show_files.helpStyle.Render(fmt.Sprintf("\ntab: focus next • n: new %s • q: exit\n", model))
	return s
}

func (m show_files.mainModel) currentFocusedModel() string {
	if m.state == show_files.timerView {
		return "timer"
	}
	return "spinner"
}

func (m *show_files.mainModel) Next() {
	if m.index == len(show_files.spinners)-1 {
		m.index = 0
	} else {
		m.index++
	}
}

func (m *show_files.mainModel) resetSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = show_files.spinnerStyle
	m.spinner.Spinner = show_files.spinners[m.index]
}
