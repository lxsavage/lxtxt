package editor

import (
	"lxsavage/lxtxt/internal/common"

	"charm.land/lipgloss/v2"
)

var (
	StyleCursorNormal = lipgloss.NewStyle().
				Background(common.Blue)

	StyleCursorInsert = lipgloss.NewStyle().
				Background(common.Green)

	StyleLineNumber = lipgloss.NewStyle().
			Foreground(common.Green).
			Padding(0, 1).
			Align(lipgloss.Right)
)
