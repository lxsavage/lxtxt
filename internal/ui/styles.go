package ui

import (
	"lxsavage/lxtxt/internal/common"
	"lxsavage/lxtxt/internal/statusbar"

	"charm.land/lipgloss/v2"
)

var (
	styleSegmentNormalMode = statusbar.StyleDefaultSegment.
				Background(common.Blue).
				Foreground(common.White)

	styleSegmentCommandMode = statusbar.StyleDefaultSegment.
				Background(common.Purple).
				Foreground(common.White)

	styleSegmentInsertMode = statusbar.StyleDefaultSegment.
				Background(common.Green).
				Foreground(common.White)

	styleSegmentVisualMode = statusbar.StyleDefaultSegment.
				Background(common.Orange).
				Foreground(common.White)

	styleCursorCommand = lipgloss.NewStyle().
				Background(common.White).
				Foreground(common.Black)
)
