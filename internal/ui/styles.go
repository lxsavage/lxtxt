package ui

import (
	"lxsavage/lxtxt/internal/common"
	"lxsavage/lxtxt/internal/statusbar"

	"charm.land/lipgloss/v2"
)

var (
	StyleSegmentNormalMode = statusbar.StyleDefaultSegment.
				Background(common.Blue).
				Foreground(common.White)

	StyleSegmentCommandMode = statusbar.StyleDefaultSegment.
				Background(common.Purple).
				Foreground(common.White)

	StyleSegmentInsertMode = statusbar.StyleDefaultSegment.
				Background(common.Green).
				Foreground(common.White)

	StyleCursorNormal = lipgloss.NewStyle().
				Background(common.Blue)

	StyleCursorInsert = lipgloss.NewStyle().
				Background(common.Green)

	StyleLineNumber = lipgloss.NewStyle().
			Foreground(common.Green).
			Padding(0, 1).
			Align(lipgloss.Right)
)
