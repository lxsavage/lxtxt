package command

import (
	"errors"
	"lxsavage/lxtxt/internal/common"
	"lxsavage/lxtxt/internal/utilities"
	"os"
	"regexp"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type command func(state common.EditorState, args []string) (tea.Cmd, error)

var commandMap = map[string]command{
	"w":      saveBuf,
	"q":      quit,
	"q!":     quit,           // Kept due to being a Vim artifact
	"wq":     saveBufAndQuit, // Kept due to being a Vim artifact
	"saveas": saveAs,
	"sed":    sub,
	"_":      invalidCommand,
}

// For future use, adds commands to the table
func Register(name string, action command) bool {
	if _, ok := commandMap[name]; ok {
		return false
	}

	commandMap[name] = action
	return true
}

func saveBuf(_ common.EditorState, _ []string) (tea.Cmd, error) {
	return SaveCmd, nil
}

func saveAs(state common.EditorState, args []string) (tea.Cmd, error) {
	if len(args) < 1 {
		return nil, ErrInsufficientArguments
	}

	expandedPath := args[0]
	if strings.Contains(expandedPath, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		expandedPath = strings.Replace(expandedPath, "~/", home+"/", 1)
	}

	if !utilities.CreateOK(expandedPath) {
		return nil, errors.New("invalid path: " + expandedPath)
	}

	state.Path = expandedPath
	return tea.Batch(
		UpdateUICmdWithState(state),
		SaveCmd,
	), nil
}

func sub(state common.EditorState, args []string) (tea.Cmd, error) {
	if len(args) < 2 {
		return nil, ErrInsufficientArguments
	}

	re, err := regexp.Compile(args[0])
	if err != nil {
		return nil, errors.New("invalid search expression")
	}

	for i, line := range state.Buf {
		state.Buf[i] = re.ReplaceAllString(line, args[1])
	}

	return UpdateUICmdWithState(state), nil
}

func quit(_ common.EditorState, _ []string) (tea.Cmd, error) {
	return tea.Quit, nil
}

func saveBufAndQuit(_ common.EditorState, _ []string) (tea.Cmd, error) {
	return tea.Batch(SaveCmd, tea.Quit), nil
}

func invalidCommand(_ common.EditorState, _ []string) (tea.Cmd, error) {
	return nil, errors.New("unknown command")
}
