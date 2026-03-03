package ui

import (
	"fmt"
	"lxsavage/lxtxt/internal/common"
	"lxsavage/lxtxt/internal/statusbar"
	"lxsavage/lxtxt/internal/utilities"
	"strconv"
)

func (m *Model) saveFile() {
	if newBuf, err := utilities.WriteFile(m.path, m.Editor.Buf); err == nil {
		m.Editor.Buf = newBuf
		if m.Editor.CursorR > len(m.Editor.Buf) {
			m.Editor.CursorR = len(m.Editor.Buf)
			m.Editor.CursorUp()
		}

		m.setDirty(false)
	}
}

func (m Model) readNumBuf() int {
	res, err := strconv.Atoi(string(m.numBuf))
	if err != nil {
		return 1
	}

	return res
}

func (m *Model) computeFileStat() {
	msg := fmt.Sprintf("%d:%d", m.Editor.CursorR+1, m.Editor.CursorC+1)

	m.status.AddSegmentOptionsById(segmentNavStatId,
		statusbar.WithText(msg),
	)
}

func (m *Model) setDirty(d bool) {
	m.dirty = d

	if m.dirty {
		m.status.SetSegmentById(segmentDirtyId, SegmentIsDirty)
	} else {
		m.status.SetSegmentById(segmentDirtyId, SegmentIsNotDirty)
	}
}

func (m *Model) setPath(path string) {
	m.path = path
	m.status.AddSegmentOptionsById(segmentPathId,
		statusbar.WithText(path),
	)
}

func (m *Model) changeMode(em common.EditorMode) {
	m.mode = em
	m.Editor.Mode = em
	m.command = ""

	switch em {
	case common.MODE_NORMAL:
		m.status.SetSegmentById(segmentModeId, SegmentNormal)
	case common.MODE_COMMAND:
		m.status.SetSegmentById(segmentModeId, SegmentCommand)
	case common.MODE_INSERT:
		m.status.SetSegmentById(segmentModeId, SegmentInsert)
	case common.MODE_VISUAL:
		m.status.SetSegmentById(segmentModeId, SegmentVisual)
	}
}
