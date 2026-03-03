package common

type EditorMode int

const (
	MODE_NORMAL EditorMode = iota
	MODE_INSERT
	MODE_COMMAND
	MODE_VISUAL
)

var modeNames = map[EditorMode]string{
	MODE_NORMAL:  "NORMAL",
	MODE_INSERT:  "INSERT",
	MODE_COMMAND: "COMMAND",
	MODE_VISUAL:  "VISUAL",
}

func (m EditorMode) String() string {
	if name, ok := modeNames[m]; ok {
		return name
	}
	return "unknown"
}
