package editor

import (
	"lxsavage/lxtxt/internal/common"

	"charm.land/lipgloss/v2"
)

var (
	styleCursorNormal = lipgloss.NewStyle().
				Background(common.Blue)

	styleCursorInsert = lipgloss.NewStyle().
				Background(common.Green)

	styleLineNumber = lipgloss.NewStyle().
			Foreground(common.Green).
			Padding(0, 1).
			Align(lipgloss.Right)
)
