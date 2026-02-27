// editor implements the main text editor component used in LXTXT
package editor

import (
	"fmt"
	"lxsavage/lxtxt/internal/common"
	"lxsavage/lxtxt/internal/utilities"
	"strconv"
	"strings"
)

type Model struct {
	Buf         []string
	CursorR     int
	CursorC     int
	RScrollBase int
	Mode        common.EditorMode
	width       int
	height      int
}

func New(buf []string) Model {
	return Model{
		Buf: buf,
	}
}

func (m *Model) SetWidth(w int) {
	m.width = w
}

func (m *Model) SetHeight(h int) {
	m.height = h
}

func (m *Model) SetDimensions(w, h int) {
	m.SetWidth(w)
	m.SetHeight(h)
}

func (m Model) View() string {
	gutterWidth := utilities.NumberWidth(len(m.Buf)) + 2
	gutterStyle := StyleLineNumber.Width(gutterWidth)

	var lines []string
	endIdx := m.RScrollBase + m.height
	if endIdx >= len(m.Buf) {
		lines = m.Buf[m.RScrollBase:]
	} else {
		lines = m.Buf[m.RScrollBase : m.RScrollBase+m.height]
	}

	var view strings.Builder
	for i, line := range lines {
		lineNum := i + m.RScrollBase
		if m.Mode != common.MODE_COMMAND && m.CursorR == lineNum {
			if m.CursorC >= len(line) {
				if m.Mode == common.MODE_INSERT {
					line += StyleCursorInsert.Render(" ")
				} else {
					line += StyleCursorNormal.Render(" ")
				}
			} else {
				newLine := ""
				if m.CursorC > 0 {
					newLine += line[:m.CursorC]
				}

				if m.Mode == common.MODE_INSERT {
					newLine += StyleCursorInsert.Render(string(line[m.CursorC]))
				} else {
					newLine += StyleCursorNormal.Render(string(line[m.CursorC]))
				}

				if m.CursorC < len(line)-1 {
					newLine += line[m.CursorC+1:]
				}
				line = newLine
			}
		}
		fmt.Fprintf(&view, "%s%s\n", gutterStyle.Render(strconv.Itoa(lineNum+1)), line)
	}

	return view.String()
}
