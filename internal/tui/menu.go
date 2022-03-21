package tui

import "github.com/charmbracelet/bubbles/list"

type MenuItem struct {
	title string
	desc  string
}

func (i MenuItem) Title() string       { return i.title }
func (i MenuItem) Description() string { return i.desc }
func (i MenuItem) FilterValue() string { return i.title }

func getInitialItems() []list.Item {
	return []list.Item{
		newMenuItem(string(objectAmmo), "ammo"),
		newMenuItem(string(objectArmor), "armor"),
	}
}

func newMenuItem(title, desc string) MenuItem {
	return MenuItem{
		title: title,
		desc:  desc,
	}
}
