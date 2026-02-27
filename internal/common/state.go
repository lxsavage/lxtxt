package common

type StateUI struct {
	Buf         []string
	CursorR     int
	CursorC     int
	RScrollBase int
	Mode        EditorMode
	width       int
	height      int
}

func NewStateUI(width, height int) StateUI {
	return StateUI{
		width:  width,
		height: height,
	}
}

func (s StateUI) GetDimensions() (int, int) {
	return s.width, s.height
}
