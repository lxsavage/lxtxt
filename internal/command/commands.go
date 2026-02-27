package command

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

type command func(args []string) (tea.Cmd, error)

var commandMap = map[string]command{
	"echo":  echo,
	"w":     write,
	"write": write,
}

func write(args []string) (tea.Cmd, error) {
	return tea.Batch(
		PrintCmdWithMessage("Writing output..."),
		SaveCmd,
	), nil
}

func echo(args []string) (tea.Cmd, error) {
	res := strings.Join(args, " ")
	return PrintCmdWithMessage(res), nil
}
