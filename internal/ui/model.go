package ui

import (
	"fmt"
	"log"
	"lxsavage/lxtxt/internal/command"
	"lxsavage/lxtxt/internal/fileio"
	"lxsavage/lxtxt/internal/statusbar"
	"lxsavage/lxtxt/internal/utilities"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	segmentModeId    = "s-mode"
	segmentNavStatId = "s-stat"
	segmentDirty     = "s-dirty"
)

var (
	SegmentNormal = statusbar.Segment(MODE_NORMAL.String(),
		statusbar.WithId(segmentModeId),
		statusbar.WithStyle(StyleSegmentNormalMode),
	)
	SegmentInsert = statusbar.Segment(MODE_INSERT.String(),
		statusbar.WithId(segmentModeId),
		statusbar.WithStyle(StyleSegmentInsertMode),
	)
	SegmentCommand = statusbar.Segment(MODE_COMMAND.String(),
		statusbar.WithId(segmentModeId),
		statusbar.WithStyle(StyleSegmentCommandMode),
	)
	SegmentIsNotDirty = statusbar.Segment("",
		statusbar.WithId(segmentDirty),
		statusbar.WithStyle(statusbar.StyleDefaultStatusBar),
	)
	SegmentIsDirty = statusbar.Segment("| +",
		statusbar.WithId(segmentDirty),
		statusbar.WithStyle(statusbar.StyleDefaultStatusBar),
	)
)

type model struct {
	origBuf     []string
	buf         []string
	status      *statusbar.Model
	path        string
	mode        mode
	width       int
	height      int
	command     string
	commandMsg  string
	cursorR     int
	cursorC     int
	rScrollBase int
	dirty       bool
}

func newModel(path string, buf []string) model {
	m := model{
		mode:    MODE_NORMAL,
		path:    path,
		origBuf: append([]string(nil), buf...),
		buf:     append([]string(nil), buf...),
	}

	s := statusbar.StatusBar(
		statusbar.WithSegments(
			statusbar.SegmentWithBase(SegmentNormal,
				statusbar.WithId(segmentModeId),
			),
			statusbar.Segment(m.path,
				statusbar.WithStyle(statusbar.StyleDefaultStatusBar.Padding(0, 1)),
			),
			SegmentIsNotDirty,
			statusbar.Segment("",
				statusbar.WithId(segmentNavStatId),
				statusbar.WithPosition(lipgloss.Right),
			),
		),
	)

	m.status = &s
	return m
}

func (m model) Init() tea.Cmd {
	return tea.RequestWindowSize
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		m.status.SetWidth(m.width)
	case tea.KeyMsg:
		if m.mode == MODE_NORMAL || m.mode == MODE_INSERT {
			switch msg.String() {
			case "up":
				m = cursorUp(m)
			case "down":
				m = cursorDown(m)
			case "left":
				m = cursorLeft(m)
			case "right":
				m = cursorRight(m)
			}
		}
		switch msg.String() {
		case "esc":
			m = m.changeMode(MODE_NORMAL)
		}

		switch m.mode {
		case MODE_COMMAND:
			return m.updateCommand(msg)
		case MODE_INSERT:
			return m.updateInsert(msg)
		case MODE_NORMAL:
			fallthrough
		default:
			return m.updateNormal(msg)
		}
	case command.PrintMsg:
		m.commandMsg = msg.Value
	}

	return m.computeFileStat(), nil
}

func (m model) View() tea.View {
	var v strings.Builder

	v.WriteString(m.status.View())
	v.WriteRune('\n')

	lineWidth := utilities.NumWidth(len(m.buf)) + 2
	lineStyle := StyleLineNumber.Width(lineWidth)

	for i, line := range m.buf[m.rScrollBase:] {
		lineNum := i + m.rScrollBase
		if m.mode != MODE_COMMAND && m.cursorR == lineNum {
			if m.cursorC >= len(line) {
				if m.mode == MODE_INSERT {
					line += StyleCursorInsert.Render(" ")
				} else {
					line += StyleCursorNormal.Render(" ")
				}
			} else {
				newLine := ""
				if m.cursorC > 0 {
					newLine += line[:m.cursorC]
				}

				if m.mode == MODE_INSERT {
					newLine += StyleCursorInsert.Render(string(line[m.cursorC]))
				} else {
					newLine += StyleCursorNormal.Render(string(line[m.cursorC]))
				}

				if m.cursorC < len(line)-1 {
					newLine += line[m.cursorC+1:]
				}
				line = newLine
			}
		}
		fmt.Fprintf(&v, "%s%s\n", lineStyle.Render(strconv.Itoa(lineNum+1)), line)
	}

	emptyLineCount := m.height - strings.Count(v.String(), "\n") - 1
	if emptyLineCount > 0 {
		v.WriteString(strings.Repeat("\n", emptyLineCount))
		if m.mode == MODE_COMMAND {
			fmt.Fprintf(&v, ":%s%s\n", m.command, StyleCursorNormal.Render(" "))
		} else {
			if len(m.commandMsg) > m.width {
				v.WriteString(m.commandMsg[:m.width-3])
				v.WriteString("...")
			} else {
				v.WriteString(m.commandMsg)
			}
		}
	}

	return tea.View{
		AltScreen: true,
		Content:   v.String(),
	}
}

func (m model) changeMode(new mode) model {
	m.mode = new
	m.command = ""

	switch new {
	case MODE_NORMAL:
		m.status.SetSegmentById(segmentModeId, SegmentNormal)
	case MODE_COMMAND:
		m.status.SetSegmentById(segmentModeId, SegmentCommand)
	case MODE_INSERT:
		m.status.SetSegmentById(segmentModeId, SegmentInsert)
	}

	return m
}

func (m model) updateNormal(msg tea.KeyMsg) (model, tea.Cmd) {
	// TODO - implement fully-fledged scrolling using u/d f/b
	switch msg.String() {
	case "u":
		if m.rScrollBase > 0 {
			m.rScrollBase--
		}
	case "d":
		if m.rScrollBase < len(m.buf)-1 {
			m.rScrollBase++

			if m.cursorR < m.rScrollBase {
				m.cursorR = m.rScrollBase
			}
		}
	case "_", "0":
		m = cursorLineStart(m)
	case "$":
		m = cursorLineEnd(m)
	case "k":
		m = cursorUp(m)
	case "j":
		m = cursorDown(m)
	case "h":
		m = cursorLeft(m)
	case "l":
		m = cursorRight(m)
	case "O":
		m = cursorLineStart(m)
		m = newline(m, 0)
		m = cursorUp(m)
		m = m.changeMode(MODE_INSERT)
	case "o":
		m = cursorLineEnd(m)
		m = newline(m, utilities.IndentLevel(m.buf[m.cursorR]))
		m = m.changeMode(MODE_INSERT)
	case "a":
		m = cursorRight(m)
		m = m.changeMode(MODE_INSERT)
	case "i":
		m = m.changeMode(MODE_INSERT)
	// case ":":
	// 	m = m.changeMode(MODE_COMMAND)
	case "!":
		if m.dirty {
			m.buf = append([]string(nil), m.origBuf...)
			m.cursorR, m.cursorC = 0, 0
			m = m.setDirty(false)
		}
	case "D":
		m = deleteline(m)
	case "W":
		if m.dirty {
			if newBuf, err := fileio.WriteFile(m.path, m.buf); err == nil {
				m.buf = newBuf
				if m.cursorR > len(m.buf) {
					m.cursorR = len(m.buf)
					m = cursorUp(m)
				}

				m = m.setDirty(false)
			}
		}
	case "q":
		return m, tea.Quit
	}
	return m.computeFileStat(), nil
}

func (m model) updateCommand(msg tea.KeyMsg) (model, tea.Cmd) {
	switch msg.String() {
	case "backspace":
		if len(m.command) > 0 {
			m.command = m.command[:len(m.command)-1]
		}
	case "enter":
		cmd, err := command.Eval(m.command)
		if err != nil {
			log.Printf("command error: %v", err)
			cmd = nil // ensure that no command is sent as a result
		}

		return m.changeMode(MODE_NORMAL), cmd
	}

	if t := msg.Key().Text; t != "" {
		m.command += t
	}
	return m, nil
}

func (m model) updateInsert(msg tea.KeyMsg) (model, tea.Cmd) {
	switch msg.String() {
	case "backspace":
		return backspace(m), nil
	case "enter":
		return newline(m, utilities.IndentLevel(m.buf[m.cursorR])), nil
	}

	if t := msg.Key().Text; t != "" {
		return insertText(m, t), nil
	}
	return m, nil
}

func (m model) computeFileStat() model {
	msg := fmt.Sprintf("%d:%d", m.cursorR+1, m.cursorC+1)

	m.status.AddSegmentOptionsById(segmentNavStatId,
		statusbar.WithText(msg),
	)
	return m
}

func (m model) setDirty(d bool) model {
	m.dirty = d

	if m.dirty {
		m.status.SetSegmentById(segmentDirty, SegmentIsDirty)
	} else {
		m.status.SetSegmentById(segmentDirty, SegmentIsNotDirty)
	}

	return m
}
