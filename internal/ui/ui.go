package ui

import (
	tea "charm.land/bubbletea/v2"
)

func Exec(path string, buf []string) error {
	p := tea.NewProgram(newModel(path, buf))
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
