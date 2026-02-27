package command

import (
	"errors"
	"lxsavage/lxtxt/internal/common"
	"regexp"
	"strings"

	tea "charm.land/bubbletea/v2"
)

var spaceRegexp = regexp.MustCompile(`\s+`)
var argRegexp = regexp.MustCompile(`'[^']*'|"[^']*"|[^\s]+`)

var (
	ErrNoCommand      = errors.New("no command")
	ErrInvalidCommand = errors.New("command not found")
)

func parse(cmd string) (string, []string, error) {
	initialSplit := spaceRegexp.Split(cmd, 2)
	if len(initialSplit) == 0 {
		return "", nil, ErrNoCommand
	}

	action := initialSplit[0]

	var args []string
	if len(initialSplit) > 1 {
		initialSplit = append(initialSplit, "")
		argsCombined := initialSplit[1]
		args := argRegexp.FindAllString(argsCombined, -1)
		for i, v := range args {
			args[i] = strings.Trim(v, "'\"")
		}
	}

	return action, args, nil
}

func Eval(state common.StateUI, cmd string) (tea.Cmd, error) {
	cmd, args, err := parse(cmd)
	if err != nil {
		return nil, err
	}

	exec, ok := commandMap[strings.ToLower(cmd)]
	if !ok {
		return nil, ErrInvalidCommand
	}

	teaCmd, err := exec(state, args)
	if err != nil {
		return nil, err
	}

	return teaCmd, nil
}
