package ui

type mode int

const (
	MODE_NORMAL mode = iota
	MODE_INSERT
	MODE_COMMAND
)

var modeNames = map[mode]string{
	MODE_NORMAL:  "NORMAL",
	MODE_INSERT:  "INSERT",
	MODE_COMMAND: "COMMAND",
}

func (m mode) String() string {
	if name, ok := modeNames[m]; ok {
		return name
	}
	return "unknown"
}
