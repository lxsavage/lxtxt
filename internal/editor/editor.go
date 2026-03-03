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
const tabSeq = "    "

type Model struct {
	Buf           []string
	CursorR       int
	CursorC       int
	ScrollBaseR   int
	ScrollBaseC   int
	Mode          common.EditorMode
	width         int
	height        int
	AnchorVisualR int
	AnchorVisualC int
}

func New(buf []string) *Model {
	return &Model{
		Buf:  buf,
		Mode: common.MODE_NORMAL,
	}
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
			toRight := m.getIsHighlightingToTheRight()
			anchoredAtCursor := m.AnchorVisualR == m.CursorR && m.AnchorVisualC == m.CursorC
			if relativeCursorC >= len(line) {
				switch m.Mode {
				case common.MODE_INSERT:
					line += styleCursorInsert.Render(" ")
				case common.MODE_VISUAL:
					if !anchoredAtCursor && toRight {
						line = line[:m.AnchorVisualC] +
							styleHighlight.Render(line[m.AnchorVisualC:])
					}

					line += styleCursorVisual.Render(" ")
				default:
					line += styleCursorNormal.Render(" ")
				}
			} else if relativeCursorC >= 0 {
				newLine := ""
				if m.CursorC > 0 {
					if m.Mode == common.MODE_VISUAL && !anchoredAtCursor && toRight {
						newLine += line[:m.AnchorVisualC] +
							styleHighlight.Render(line[m.AnchorVisualC:m.CursorC])
					} else {
						newLine += line[:relativeCursorC]
					}
				}

				highlighted := string(line[relativeCursorC])
				if highlighted[0] == '\t' {
					highlighted = tabSeq
				}

				switch m.Mode {
				case common.MODE_INSERT:
					newLine += styleCursorInsert.Render(highlighted)
				case common.MODE_VISUAL:
					newLine += styleCursorVisual.Render(highlighted)
				default:
					newLine += styleCursorNormal.Render(highlighted)
				}
				// newLine += highlighted

				if relativeCursorC < len(line)-1 {
					if m.Mode == common.MODE_VISUAL && !anchoredAtCursor && !toRight {
						newLine += styleHighlight.Render(line[relativeCursorC+1 : m.AnchorVisualC+1])

						if len(line) > m.AnchorVisualC+1 {
							newLine += line[m.AnchorVisualC+1:]
						}
					} else {
						newLine += line[relativeCursorC+1:]
					}
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

func (m Model) getIsHighlightingToTheRight() bool {
	isOnLineBelow := m.AnchorVisualR < m.CursorR
	isRightOfAnchor := m.AnchorVisualR == m.CursorR && m.AnchorVisualC < m.CursorC

	return isOnLineBelow || isRightOfAnchor
}
