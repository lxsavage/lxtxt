package command

import (
	"lxsavage/lxtxt/internal/common"

	tea "charm.land/bubbletea/v2"
)

type UpdateUIMsg common.EditorState
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

func UpdateUICmdWithState(s common.EditorState) tea.Cmd {
	return func() tea.Msg {
		return UpdateUIMsg(s)
	}
}

func SaveCmd() tea.Msg {
	return SaveMsg{}
}
