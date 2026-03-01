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
	case "k":
		m.Editor.RepeatMotion(m.readNumBuf(), m.Editor.CursorUp)
	case "j":
		m.Editor.RepeatMotion(m.readNumBuf(), m.Editor.CursorDown)
	case "h":
		m.Editor.RepeatMotion(m.readNumBuf(), m.Editor.CursorLeft)
	case "l":
		m.Editor.RepeatMotion(m.readNumBuf(), m.Editor.CursorRight)
	case "D":
		// TODO - evaluate if this should be repeatable
		if m.Editor.RepeatMotion(m.readNumBuf(), m.Editor.Deleteline) {
			m.setDirty(true)
		}
	case "_":
		m.Editor.CursorLineStart()
	case "$":
		m.Editor.CursorLineEnd()
	case "O":
		m.Editor.CursorLineStart()
		m.Editor.Newline(0)
		m.Editor.CursorUp()
		m.changeMode(common.MODE_INSERT)
	case "o":
		m.Editor.CursorLineEnd()
		m.Editor.Newline(utilities.IndentLevel(m.Editor.Buf[m.Editor.CursorR]))
		m.changeMode(common.MODE_INSERT)
	case "a":
		m.Editor.CursorRight()
		m.changeMode(common.MODE_INSERT)
	case "i":
		m.changeMode(common.MODE_INSERT)
	case ":":
		m.changeMode(common.MODE_COMMAND)
	case "!":
		if m.dirty {
			m.Editor.Buf = append([]string(nil), m.origBuf...)
			m.Editor.CursorR, m.Editor.CursorC = 0, 0

			m.setDirty(false)
		}
	case "W":
		if m.dirty {
			m.saveFile()
		}
	case "?":
		// NOTE - if install script has not been run, this won't do anything due to
		// the manpage not being installed to the proper directory
		utilities.ShowMan()
	case "Q":
		return m, tea.Quit
	}

	keyCodeByte := msg.String()[0]
	if keyCodeByte >= '0' && keyCodeByte <= '9' {
		m.numBuf = append(m.numBuf, keyCodeByte)
	} else {
		m.numBuf = m.numBuf[:0]
	}

	m.computeFileStat()
	return m, nil
}

func (m Model) updateCommand(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "backspace":
		if len(m.command) > 0 {
			m.command = m.command[:len(m.command)-1]
		}
	case "enter":
		state := m.Editor.ToState()
		cmd, err := command.Eval(state, m.command)
		if err != nil {
			m.CommandMessage = fmt.Sprintf("command error: %v", err)
			cmd = nil
		}
		m.changeMode(common.MODE_NORMAL)
		return m, cmd
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
		if m.Editor.Backspace() {
			m.setDirty(true)
		}
	case "delete":
		if m.Editor.Delete() {
			m.setDirty(true)
		}
	case "tab":
		if m.Editor.InsertText("\t") {
			m.setDirty(true)
		}
	case "enter":
		if m.Editor.Newline(utilities.IndentLevel(m.Editor.Buf[m.Editor.CursorR])) {
			m.setDirty(true)
		}
	default:
		if t := msg.Key().Text; t != "" {
			if m.Editor.InsertText(t) {
				m.setDirty(true)
			}
		}
	}
	return m, nil
}
