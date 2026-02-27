package command

import tea "charm.land/bubbletea/v2"

type PrintMsg struct {
	Value string
}

type SaveMsg struct{}

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
