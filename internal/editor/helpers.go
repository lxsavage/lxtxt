package editor

import "lxsavage/lxtxt/internal/common"

// Unused for now until commands are fully implemented
func (m Model) ToState() common.StateUI {
	return common.StateUI{
		Buf: append([]string(nil), m.Buf...),
	}
}
