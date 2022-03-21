package tui

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jon4hz/eldencli/pkg/eldenring"
)

type errMsg error

type state int

const (
	stateReady state = iota
	stateSelectObject
	stateViewObject
)

type model struct {
	ctx    context.Context
	state  state
	list   list.Model
	client *eldenring.Client
	err    error
}

func InitalModel() tea.Model {
	del := list.NewDefaultDelegate()
	del.Styles.SelectedDesc.Foreground(FocusedStyle.GetForeground()).BorderForeground(FocusedStyle.GetForeground())
	del.Styles.SelectedTitle.Foreground(FocusedStyle.GetForeground()).BorderForeground(FocusedStyle.GetForeground())
	m := model{
		ctx:    context.Background(),
		state:  stateReady,
		list:   list.NewModel(getInitialItems(), del, 0, 0),
		client: eldenring.NewClient(),
	}
	m.list.Styles.FilterCursor = FocusedStyle
	m.list.SetFilteringEnabled(true)
	m.list.SetShowStatusBar(true)
	m.list.Title = "eldencli"
	m.list.Styles.Title = TitleStyle
	return m
}

func Start() error {
	return tea.NewProgram(InitalModel(), tea.WithAltScreen()).Start()
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case stateReady:
			switch msg.String() {
			case "enter":
				switch m.list.SelectedItem().(MenuItem).title {
				case string(objectAmmo):
					items, err := m.getAmmosItems()
					if err != nil {
						return m, func() tea.Msg { return errMsg(err) }
					}
					m.state = stateSelectObject
					cmds = append(cmds, m.list.SetItems(items))
				}
			}

		case stateSelectObject:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.state = stateReady
				cmds = append(cmds, m.list.SetItems(getInitialItems()))
				return m, tea.Batch(cmds...)
			case "enter":
				m.state = stateViewObject
				return m, nil
			}

		case stateViewObject:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.state = stateSelectObject
				return m, nil
			}
		}

	case tea.WindowSizeMsg:
		top, right, bottom, left := DocStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)

	case errMsg:
		m.err = error(msg)
	}

	var cmd tea.Cmd
	switch m.state {
	case stateReady, stateSelectObject:
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.err != nil {
		return DocStyle.Render(m.err.Error())
	}
	switch m.state {
	case stateViewObject:
		switch o := m.list.SelectedItem().(ObjectItem).object.(type) {
		case eldenring.Ammo:
			return o.Render()
		}
		return ""
	default:
		return DocStyle.Render(m.list.View())
	}
}
