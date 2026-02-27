package ui

import (
	"lxsavage/lxtxt/internal/statusbar"

	"charm.land/lipgloss/v2"
)

var (
	blue   = lipgloss.Blue
	green  = lipgloss.Green
	purple = lipgloss.Magenta
	white  = lipgloss.Color("#ffffff")
)

var (
	StyleSegmentNormalMode = statusbar.StyleDefaultSegment.
				Background(blue).
				Foreground(white)

	StyleSegmentCommandMode = statusbar.StyleDefaultSegment.
				Background(purple).
				Foreground(white)

	StyleSegmentInsertMode = statusbar.StyleDefaultSegment.
				Background(green).
				Foreground(white)

	StyleCursorNormal = lipgloss.NewStyle().
				Background(blue)

	StyleCursorInsert = lipgloss.NewStyle().
				Background(green)

	StyleLineNumber = lipgloss.NewStyle().
			Foreground(green).
			Padding(0, 1).
			Align(lipgloss.Right)
)
