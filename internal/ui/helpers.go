package ui

import (
	"fmt"
	"lxsavage/lxtxt/internal/common"
	"lxsavage/lxtxt/internal/fileio"
	"lxsavage/lxtxt/internal/statusbar"
)

func (m *Model) SaveFile() {
	if newBuf, err := fileio.WriteFile(m.Path, m.editor.Buf); err == nil {
		m.editor.Buf = newBuf
		if m.editor.CursorR > len(m.editor.Buf) {
			m.editor.CursorR = len(m.editor.Buf)
			m.editor.CursorUp()
		}

		m.setDirty(false)
	}
}

func (m *Model) computeFileStat() {
	msg := fmt.Sprintf("%d:%d", m.editor.CursorR+1, m.editor.CursorC+1)

	m.status.AddSegmentOptionsById(segmentNavStatId,
		statusbar.WithText(msg),
	)
}

func (m *Model) setDirty(d bool) {
	m.dirty = d

	if m.dirty {
		m.status.SetSegmentById(segmentDirty, SegmentIsDirty)
	} else {
		m.status.SetSegmentById(segmentDirty, SegmentIsNotDirty)
	}
}

func (m *Model) changeMode(em common.EditorMode) {
	m.Mode = em
	m.editor.Mode = em
	m.command = ""

	switch em {
	case common.MODE_NORMAL:
		m.status.SetSegmentById(segmentModeId, SegmentNormal)
	case common.MODE_COMMAND:
		m.status.SetSegmentById(segmentModeId, SegmentCommand)
	case common.MODE_INSERT:
		m.status.SetSegmentById(segmentModeId, SegmentInsert)
	}
}
