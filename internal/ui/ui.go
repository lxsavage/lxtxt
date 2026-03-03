// ui implements the main UI for LXTXT
package ui

import (
	tea "charm.land/bubbletea/v2"
)

func Exec(path string, buf []string, experimental bool) error {
	p := tea.NewProgram(newModel(path, buf, experimental))
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
