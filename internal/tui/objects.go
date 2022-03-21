package tui

import (
	"context"

	"github.com/charmbracelet/bubbles/list"
)

type Object string

const (
	objectAmmo  Object = "Ammo"
	objectArmor Object = "Armor"
)

type ObjectItem struct {
	title  string
	desc   string
	object interface{}
}

func (i ObjectItem) Title() string       { return i.title }
func (i ObjectItem) Description() string { return i.desc }
func (i ObjectItem) FilterValue() string { return i.title }

func (m model) getAmmosItems() ([]list.Item, error) {
	ammos, err := m.client.GetAmmos(context.Background())
	if err != nil {
		return nil, err
	}

	items := make([]list.Item, 0, len(ammos))
	for _, v := range ammos {
		items = append(items, ObjectItem{
			title:  v.Name,
			desc:   v.Description,
			object: v,
		})
	}

	return items, nil
}
