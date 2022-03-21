package eldenring

import (
	"strconv"
	"strings"
)

func (a Ammo) Render() string {
	var s strings.Builder

	s.WriteString("Name: ")
	s.WriteString(a.Name)
	s.WriteByte('\n')

	s.WriteString("Description: ")
	s.WriteString(a.Description)
	s.WriteByte('\n')

	s.WriteString("Damage type: ")
	s.WriteString(a.Type)
	s.WriteByte('\n')

	s.WriteString("Passive: ")
	s.WriteString(a.Passive)
	s.WriteByte('\n')

	if a.AttackPower != nil {
		s.WriteString("Attack Powers: ")
		s.WriteByte('\n')
		for _, v := range a.AttackPower {
			if v.Amount != 0 {
				s.WriteString(v.Name)
				s.WriteString(": ")
				s.WriteString(strconv.Itoa(v.Amount))
				s.WriteByte('\n')
			}
		}
	}
	return s.String()
}
