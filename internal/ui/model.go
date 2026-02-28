package ui

import (
	"fmt"
	"lxsavage/lxtxt/internal/command"
	"lxsavage/lxtxt/internal/common"
	"lxsavage/lxtxt/internal/editor"
	"lxsavage/lxtxt/internal/statusbar"
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
	SegmentNormal = statusbar.Segment(common.MODE_NORMAL.String(),
		statusbar.WithId(segmentModeId),
		statusbar.WithStyle(styleSegmentNormalMode),
	)
	SegmentInsert = statusbar.Segment(common.MODE_INSERT.String(),
		statusbar.WithId(segmentModeId),
		statusbar.WithStyle(styleSegmentInsertMode),
	)
	SegmentCommand = statusbar.Segment(common.MODE_COMMAND.String(),
		statusbar.WithId(segmentModeId),
		statusbar.WithStyle(styleSegmentCommandMode),
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

type Model struct {
	origBuf        []string
	status         *statusbar.Model
	Path           string
	Mode           common.EditorMode
	width          int
	height         int
	command        string
	CommandMessage string
	numBuf         []byte
	editor         *editor.Model
	dirty          bool
}

func newModel(path string, buf []string) Model {
	m := Model{
		Mode:    common.MODE_NORMAL,
		Path:    path,
		origBuf: append([]string(nil), buf...),
		editor:  editor.New(append([]string(nil), buf...)),
	}

	s := statusbar.StatusBar(
		statusbar.WithSegments(
			statusbar.SegmentWithBase(SegmentNormal,
				statusbar.WithId(segmentModeId),
			),
			statusbar.Segment(m.Path,
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

func (m Model) Init() tea.Cmd {
	return tea.RequestWindowSize
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		m.status.SetWidth(m.width)

		// TODO - find a non-hacky way of setting up the editor height
		m.editor.SetDimensions(msg.Width, msg.Height-1)
	case tea.KeyMsg:
		if m.Mode == common.MODE_NORMAL || m.Mode == common.MODE_INSERT {
			switch msg.String() {
			case "up":
				m.editor.CursorUp()
			case "down":
				m.editor.CursorDown()
			case "left":
				m.editor.CursorLeft()
			case "right":
				m.editor.CursorRight()
			}
		}
		switch msg.String() {
		case "esc":
			m.changeMode(common.MODE_NORMAL)
		}

		switch m.Mode {
		case common.MODE_COMMAND:
			return m.updateCommand(msg)
		case common.MODE_INSERT:
			return m.updateInsert(msg)
		case common.MODE_NORMAL:
			fallthrough
		default:
			return m.updateNormal(msg)
		}
	case command.PrintMsg:
		m.CommandMessage = msg.Value
	}

	m.computeFileStat()
	return m, nil
}

func (m Model) View() tea.View {
	var v strings.Builder

	v.WriteString(m.status.View())
	v.WriteString(m.editor.View())

	if m.Mode == common.MODE_COMMAND {
		fmt.Fprintf(&v, ":%s%s\n", m.command, styleCursorCommand.Render(" "))
	} else {
		for _, b := range m.numBuf {
			v.WriteByte(b)
		}
	}

	return tea.View{
		AltScreen: true,
		Content:   v.String(),
	}
}
