package command

import (
	"lxsavage/lxtxt/internal/common"

	tea "charm.land/bubbletea/v2"
)

type UpdateUIMsg common.StateUI

type SaveMsg struct{}

type PrintMsg struct {
	Value string
}

func PrintCmdWithMessage(msg string) tea.Cmd {
	return func() tea.Msg {
		return PrintMsg{
			Value: msg,
		}
	}
}

func SaveCmd() tea.Msg {
	return SaveMsg{}
}

func UpdateUICmdWithState(s common.StateUI) tea.Cmd {
	return func() tea.Msg {
		return s
	}
}
