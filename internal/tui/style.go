package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	DocStyle            = lipgloss.NewStyle().Margin(1, 2)
	ModuleWrapper       = lipgloss.NewStyle().Margin(1, 4)
	FocusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#c8a35c"))
	BlurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	CursorStyle         = FocusedStyle
	NoStyle             = lipgloss.NewStyle()
	HelpStyle           = BlurredStyle
	CursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	TitleStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#c8a35c")).Background(lipgloss.Color("#1c4341")).Bold(true)
	FooterStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	ErrStyle            = lipgloss.NewStyle().Foreground(lipgloss.Color("#f00")).Bold(true)
)
