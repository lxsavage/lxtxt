package editor

import (
	"lxsavage/lxtxt/internal/common"
	"lxsavage/lxtxt/internal/utilities"
)

func (m Model) ToState() common.EditorState {
	return common.EditorState{
		Buf:         append([]string(nil), m.Buf...),
		CursorR:     m.CursorR,
		CursorC:     m.CursorC,
		ScrollBaseR: m.ScrollBaseR,
		ScrollBaseC: m.ScrollBaseC,
		Width:       m.width,
		Height:      m.height,
	}
}

func (m *Model) ApplyStateUI(s common.EditorState) {
	m.Buf = s.Buf
	m.CursorR = s.CursorR
	m.CursorC = s.CursorC
	m.ScrollBaseR = s.ScrollBaseR
	m.ScrollBaseC = s.ScrollBaseC

	m.correctHorizontalScrolling()
	m.correctVerticalScrolling()
}

func (m Model) EditorWidth() int {
	gutterWidth := utilities.NumberWidth(len(m.Buf)) + 2
	return m.width - gutterWidth
}
