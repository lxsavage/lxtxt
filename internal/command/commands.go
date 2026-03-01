package command

import (
	"errors"
	"lxsavage/lxtxt/internal/common"
	"regexp"

	tea "charm.land/bubbletea/v2"
)

type command func(state common.EditorState, args []string) (tea.Cmd, error)

var commandMap = map[string]command{
	"w":   saveBuf,
	"q":   quit,
	"q!":  quit, // Kept due to being a Vim artifact
	"wq":  saveBufAndQuit,
	"sed": sub,
	"_":   invalidCommand,
}

// For future use, adds commands to the table
func Register(name string, action command) bool {
	if _, ok := commandMap[name]; ok {
		return false
	}

	commandMap[name] = action
	return true
}

func saveBuf(state common.EditorState, args []string) (tea.Cmd, error) {
	return SaveCmd, nil
}

func sub(state common.EditorState, args []string) (tea.Cmd, error) {
	if len(args) < 2 {
		return nil, errors.New("insufficient arguments")
	}

	re, err := regexp.Compile(args[0])
	if err != nil {
		return nil, errors.New("Invalid search expression")
	}

	for i, line := range state.Buf {
		state.Buf[i] = re.ReplaceAllString(line, args[1])
	}

	return UpdateUICmdWithState(state), nil
}

func quit(state common.EditorState, args []string) (tea.Cmd, error) {
	return tea.Quit, nil
}

func saveBufAndQuit(state common.EditorState, args []string) (tea.Cmd, error) {
	return tea.Batch(SaveCmd, tea.Quit), nil
}

func invalidCommand(state common.EditorState, args []string) (tea.Cmd, error) {
	return nil, errors.New("Invalid command")
}
