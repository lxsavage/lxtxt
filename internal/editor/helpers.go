package editor

import (
	"lxsavage/lxtxt/internal/common"
	"lxsavage/lxtxt/internal/utilities"
)

// Unused for now until commands are fully implemented
func (m Model) ToState() common.StateUI {
	s := common.NewStateUI(m.width, m.height)
	s.Buf = append([]string(nil), m.Buf...)
	s.CursorR = m.CursorR
	s.CursorC = m.CursorC
	s.RScrollBase = m.ScrollBaseR
	s.Mode = m.Mode
	return s
}

// TODO - implement
func (m *Model) ApplyStateUI(s common.StateUI) {
}

// TODO - cache this
func (m Model) EditorWidth() int {
	gutterWidth := utilities.NumberWidth(len(m.Buf)) + 2
	return m.width - gutterWidth
}
