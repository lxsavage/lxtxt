package ui

import (
	"fmt"
	"lxsavage/lxtxt/internal/command"
	"lxsavage/lxtxt/internal/common"
	"lxsavage/lxtxt/internal/utilities"

	tea "charm.land/bubbletea/v2"
)

func (m Model) updateNormal(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "_", "0":
		m.editor.CursorLineStart()
	case "$":
		m.editor.CursorLineEnd()
	case "k":
		m.editor.CursorUp()
	case "j":
		m.editor.CursorDown()
	case "h":
		m.editor.CursorLeft()
	case "l":
		m.editor.CursorRight()
	case "O":
		m.editor.CursorLineStart()
		m.editor.Newline(0)
		m.editor.CursorUp()
		m = m.changeMode(common.MODE_INSERT)
	case "o":
		m.editor.CursorLineEnd()
		m.editor.Newline(utilities.IndentLevel(m.editor.Buf[m.editor.CursorR]))
		m = m.changeMode(common.MODE_INSERT)
	case "a":
		m.editor.CursorRight()
		m = m.changeMode(common.MODE_INSERT)
	case "i":
		m = m.changeMode(common.MODE_INSERT)
	// case ":":
	// 	m = m.changeMode(common.MODE_COMMAND)
	case "!":
		if m.dirty {
			m.editor.Buf = append([]string(nil), m.origBuf...)
			m.editor.CursorR, m.editor.CursorC = 0, 0

			m = m.setDirty(false)
		}
	case "D":
		if m.editor.Deleteline() {
			m = m.setDirty(true)
		}
	case "W":
		if m.dirty {
			m = m.SaveFile()
		}
	case "q":
		return m, tea.Quit
	}
	return m.computeFileStat(), nil
}

func (m Model) updateCommand(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "backspace":
		if len(m.command) > 0 {
			m.command = m.command[:len(m.command)-1]
		}
	case "enter":
		state := m.editor.ToState()
		cmd, err := command.Eval(state, m.command)
		if err != nil {
			m.CommandMessage = fmt.Sprintf("command error: %v", err)
			cmd = nil
		}

		return m.changeMode(common.MODE_NORMAL), cmd
	default:
		if t := msg.Key().Text; t != "" {
			m.command += t
		}
	}
	return m, nil
}

func (m Model) updateInsert(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "backspace":
		if m.editor.Backspace() {
			m = m.setDirty(true)
		}
	case "enter":
		if m.editor.Newline(utilities.IndentLevel(m.editor.Buf[m.editor.CursorR])) {
			m = m.setDirty(true)
		}
	default:
		if t := msg.Key().Text; t != "" {
			if m.editor.InsertText(t) {
				m = m.setDirty(true)
			}
		}
	}
	return m, nil
}
