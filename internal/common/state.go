package common

type EditorState struct {
	Buf         []string
	CursorR     int
	CursorC     int
	ScrollBaseR int
	ScrollBaseC int

	// Not modifiable by command

	Width  int
	Height int
}
