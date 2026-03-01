// editor implements the main text editor component used in LXTXT
package editor

import (
	"fmt"
	"lxsavage/lxtxt/internal/common"
	"lxsavage/lxtxt/internal/utilities"
	"strconv"
	"strings"
)

// tabSeq is what a tab is rendered as in the editor
var tabSeq = strings.Repeat(" ", 4)

type Model struct {
	Buf         []string
	CursorR     int
	CursorC     int
	ScrollBaseR int
	ScrollBaseC int
	Mode        common.EditorMode
	width       int
	height      int
}

func New(buf []string) *Model {
	return &Model{
		Buf:  buf,
		Mode: common.MODE_NORMAL,
	}
}

func (m *Model) SetWidth(w int) {
	m.width = w
}

func (m *Model) SetHeight(h int) {
	m.height = h
}

func (m *Model) SetDimensions(w, h int) {
	m.width = w
	m.height = h
}

func (m Model) View() string {
	gutterWidth := utilities.NumberWidth(len(m.Buf)) + 2
	gutterStyle := styleLineNumber.Width(gutterWidth)

	var lines []string
	endIdx := m.ScrollBaseR + m.height
	if endIdx >= len(m.Buf) {
		lines = m.Buf[m.ScrollBaseR:]
	} else {
		lines = m.Buf[m.ScrollBaseR : m.ScrollBaseR+m.height]
	}

	var view strings.Builder
	for i, line := range lines {
		if len(line) < m.ScrollBaseC {
			line = ""
		} else if m.ScrollBaseC > 0 {
			line = line[m.ScrollBaseC:]
		}
		relativeCursorC := m.CursorC - m.ScrollBaseC

		lineNum := i + m.ScrollBaseR
		if m.Mode != common.MODE_COMMAND && m.CursorR == lineNum {
			if relativeCursorC >= len(line) {
				if m.Mode == common.MODE_INSERT {
					line += styleCursorInsert.Render(" ")
				} else {
					line += styleCursorNormal.Render(" ")
				}
			} else if relativeCursorC >= 0 {
				newLine := ""
				if m.CursorC > 0 {
					newLine += line[:relativeCursorC]
				}

				highlighted := string(line[relativeCursorC])
				if highlighted[0] == '\t' {
					highlighted = tabSeq
				}
				if m.Mode == common.MODE_INSERT {
					newLine += styleCursorInsert.Render(highlighted)
				} else {
					newLine += styleCursorNormal.Render(highlighted)
				}

				if relativeCursorC < len(line)-1 {
					newLine += line[relativeCursorC+1:]
				}
				line = newLine
			}
		}
		line = strings.ReplaceAll(line, "\t", tabSeq)
		fmt.Fprintf(&view, "%s%s\n", gutterStyle.Render(strconv.Itoa(lineNum+1)), line)
	}

	vs := view.String()
	emptyLineCount := m.height - strings.Count(vs, "\n") - 1
	if emptyLineCount > 0 {
		return vs + styleEmptyLine.Render(strings.Repeat("~\n", emptyLineCount))
	}

	return vs
}
