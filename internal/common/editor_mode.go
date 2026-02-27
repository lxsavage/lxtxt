package common

type EditorMode int

const (
	MODE_NORMAL EditorMode = iota
	MODE_INSERT
	MODE_COMMAND
)

var modeNames = map[EditorMode]string{
	MODE_NORMAL:  "NORMAL",
	MODE_INSERT:  "INSERT",
	MODE_COMMAND: "COMMAND",
}

func (m EditorMode) String() string {
	if name, ok := modeNames[m]; ok {
		return name
	}
	return "unknown"
}
