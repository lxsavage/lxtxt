package command

import (
	"lxsavage/lxtxt/internal/common"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type command func(state common.StateUI, args []string) (tea.Cmd, error)

var commandMap = map[string]command{
	"echo":  echo,
	"w":     write,
	"write": write,
}

func write(state common.StateUI, args []string) (tea.Cmd, error) {
	return tea.Batch(
		PrintCmdWithMessage("Writing output..."),
		SaveCmd,
	), nil
}

func echo(state common.StateUI, args []string) (tea.Cmd, error) {
	res := strings.Join(args, " ")
	return PrintCmdWithMessage(res), nil
}

func _invalidCmd(state common.StateUI, args []string) (tea.Cmd, error) {
	return PrintCmdWithMessage("Invalid command"), nil
}
