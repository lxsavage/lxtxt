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
	segmentDirtyId   = "s-dirty"
	segmentPathId    = "s-path"
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
	SegmentVisual = statusbar.Segment(common.MODE_VISUAL.String(),
		statusbar.WithId(segmentModeId),
		statusbar.WithStyle(styleSegmentVisualMode),
	)
	SegmentIsNotDirty = statusbar.Segment("",
		statusbar.WithId(segmentDirtyId),
		statusbar.WithStyle(statusbar.StyleDefaultStatusBar),
	)
	SegmentIsDirty = statusbar.Segment("| +",
		statusbar.WithId(segmentDirtyId),
		statusbar.WithStyle(statusbar.StyleDefaultStatusBar),
	)
)

type Model struct {
	origBuf        []string
	numBuf         []byte
	path           string
	command        string
	commandMessage string
	status         *statusbar.Model
	Editor         *editor.Model
	mode           common.EditorMode
	width          int
	height         int
	dirty          bool
	experiments    bool
}

func newModel(path string, buf []string, experiments bool) Model {
	m := Model{
		mode:        common.MODE_NORMAL,
		path:        path,
		origBuf:     append([]string(nil), buf...),
		Editor:      editor.New(append([]string(nil), buf...)),
		experiments: experiments,
	}

	s := statusbar.StatusBar(
		statusbar.WithSegments(
			statusbar.SegmentWithBase(SegmentNormal,
				statusbar.WithId(segmentModeId),
			),
			statusbar.Segment(m.path,
				statusbar.WithId(segmentPathId),
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
		m.Editor.SetDimensions(msg.Width, msg.Height-1)
	case tea.KeyMsg:
		if msg.String() == "esc" {
			m.changeMode(common.MODE_NORMAL)
		}

		if m.mode == common.MODE_NORMAL || m.mode == common.MODE_INSERT || m.mode == common.MODE_VISUAL {
			switch msg.String() {
			case "up":
				m.Editor.CursorUp()
			case "down":
				m.Editor.CursorDown()
			case "left":
				m.Editor.CursorLeft()
			case "right":
				m.Editor.CursorRight()
			}
		}

		switch m.mode {
		case common.MODE_COMMAND:
			return m.updateCommand(msg)
		case common.MODE_INSERT:
			return m.updateInsert(msg)
		case common.MODE_VISUAL:
			fallthrough
		case common.MODE_NORMAL:
			fallthrough
		default:
			return m.updateNormal(msg)
		}
	case command.PrintMsg:
		m.commandMessage = msg.Value
	case command.SaveMsg:
		if len(m.path) > 0 {
			m.saveFile()
		} else {
			m.commandMessage = "no file specified; use :saveas <path> to save to a file."
		}
	case command.UpdateUIMsg:
		m.Editor.ApplyStateUI(common.EditorState(msg))
		m.setPath(msg.Path)
		m.setDirty(true)
	}

	m.computeFileStat()
	return m, nil
}

func (m Model) View() tea.View {
	var v strings.Builder

	v.WriteString(m.status.View())
	v.WriteString(m.Editor.View())

	if m.mode == common.MODE_COMMAND {
		fmt.Fprintf(&v, ":%s%s\n", m.command, styleCursorCommand.Render(" "))
	} else if len(m.numBuf) > 0 {
		for _, b := range m.numBuf {
			v.WriteByte(b)
		}
	} else if m.mode == common.MODE_VISUAL {
		fmt.Fprintf(&v, "%d,%d <- %d,%d",
			m.Editor.AnchorVisualR, m.Editor.AnchorVisualC,
			m.Editor.CursorR, m.Editor.CursorC,
		)
	} else {
		v.WriteString(m.commandMessage)
	}

	return tea.View{
		AltScreen: true,
		Content:   v.String(),
	}
}
