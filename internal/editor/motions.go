package editor

import (
	"strings"
)

// A Motion is a function that might change the state of the editor, which
// returns true if successful or false if not
type Motion func() bool

func (m *Model) adjustVerticalScrolling() {
	if m.ScrollBaseR > m.CursorR {
		m.ScrollBaseR = m.CursorR
	} else if m.ScrollBaseR+m.height-1 <= m.CursorR {
		m.ScrollBaseR = m.CursorR - m.height + 1
	}
}

func (m *Model) adjustHorizontalScrolling() {
	ew := m.EditorWidth()
	if m.ScrollBaseC > m.CursorC {
		m.ScrollBaseC = m.CursorC
	} else if m.ScrollBaseC+ew-1 < m.CursorC {
		m.ScrollBaseC = m.CursorC - ew + 1
	}
}

// RepeatMotion runs a motion up to a specified number of times, but will stop
// early if the motion fails
func (m *Model) RepeatMotion(repeat int, motion Motion) bool {
	if repeat <= 0 {
		return false
	}

	didAnything := false
	for range repeat {
		if !motion() {
			break
		}
		didAnything = true
	}
	return didAnything
}

func (m *Model) CursorUp() bool {
	if m.CursorR <= 0 {
		return false
	}
	m.CursorR--

	if m.CursorC > len(m.Buf[m.CursorR])-1 {
		m.CursorC = len(m.Buf[m.CursorR])
	}

	m.adjustVerticalScrolling()
	m.adjustHorizontalScrolling()
	return true
}

func (m *Model) CursorDown() bool {
	if m.CursorR >= len(m.Buf)-1 {
		return false
	}
	m.CursorR++

	if m.CursorC > len(m.Buf[m.CursorR])-1 {
		m.CursorC = len(m.Buf[m.CursorR])
	}

	m.adjustVerticalScrolling()
	m.adjustHorizontalScrolling()
	return true
}

func (m *Model) CursorLeft() bool {
	if m.CursorC <= 0 {
		return false
	}
	m.CursorC--

	m.adjustHorizontalScrolling()
	return true
}

func (m *Model) CursorRight() bool {
	if m.CursorC >= len(m.Buf[m.CursorR]) {
		return false
	}
	m.CursorC++

	m.adjustHorizontalScrolling()
	return true
}

func (m *Model) CursorLineStart() bool {
	// ScrollBase does not need to be checked because it will always be 0 if CursorC is
	if m.CursorC == 0 {
		return false
	}

	m.CursorC = 0
	m.ScrollBaseC = 0
	return true
}

func (m *Model) CursorLineEnd() bool {
	target := len(m.Buf[m.CursorR])
	if m.CursorC == target {
		return false
	}

	m.CursorC = target

	m.adjustHorizontalScrolling()
	return true
}

func (m *Model) Backspace() bool {
	if m.CursorC > 0 {
		oldLine := m.Buf[m.CursorR]
		newLine := ""

		newLine += oldLine[:m.CursorC-1]
		if m.CursorC < len(oldLine) {
			newLine += oldLine[m.CursorC:]
		}

		m.Buf[m.CursorR] = newLine
		m.CursorC--

		m.adjustHorizontalScrolling()
		return true
	} else if m.CursorR > 0 {
		oldLine := m.Buf[m.CursorR]
		newBuf := append(m.Buf[:m.CursorR], m.Buf[m.CursorR+1:]...)
		m.CursorR--
		m.CursorC = len(m.Buf[m.CursorR])

		newBuf[m.CursorR] += oldLine
		m.Buf = newBuf

		return true
	}
	return false
}

func (m *Model) Delete() bool {
	if m.CursorC >= len(m.Buf[m.CursorR]) {
		if m.CursorR >= len(m.Buf)-1 {
			return false
		}

		nextLine := m.Buf[m.CursorR+1]
		newBuf := append(m.Buf[:m.CursorR], m.Buf[m.CursorR]+nextLine)
		m.Buf = newBuf
	} else {
		oldLine := m.Buf[m.CursorR]
		newLine := oldLine[:m.CursorC]
		if len(oldLine) > m.CursorR+1 {
			newLine += oldLine[m.CursorC+1:]
		}

		m.Buf[m.CursorR] = newLine
	}

	return true
}

func (m *Model) DeleteLine() bool {
	if len(m.Buf) <= 1 {
		if m.Buf[0] == "" {
			return false
		}

		m.Buf = []string{""}
		return true
	}

	if m.CursorR == 0 {
		m.Buf = m.Buf[1:]
	} else {
		if m.CursorR < len(m.Buf) {
			m.Buf = append(m.Buf[:m.CursorR], m.Buf[m.CursorR+1:]...)
		}

		if m.CursorR >= len(m.Buf) {
			m.CursorR = len(m.Buf)
			m.CursorUp()
		}
	}

	return true
}

func (m *Model) NewLine(indent int) bool {
	oldLine := ""
	newLine := ""
	if indent > 0 {
		newLine = strings.Repeat(" ", indent)
	}

	if m.CursorC > 0 {
		oldLine = m.Buf[m.CursorR][:m.CursorC]
	}
	if m.CursorC < len(m.Buf[m.CursorR])-1 {
		newLine += m.Buf[m.CursorR][m.CursorC:]
	}

	var newBuf []string
	if m.CursorR > 0 {
		newBuf = append(newBuf, m.Buf[:m.CursorR]...)
	}
	newBuf = append(newBuf, oldLine, newLine)

	if m.CursorR < len(m.Buf)-1 {
		newBuf = append(newBuf, m.Buf[m.CursorR+1:]...)
	}

	m.CursorR++
	if m.CursorC > len(newLine) {
		m.CursorC = len(newLine)
	}
	m.Buf = newBuf
	m.CursorC = indent

	m.adjustVerticalScrolling()
	return true
}

func (m *Model) InsertText(t string) bool {
	if len(m.Buf[m.CursorR]) == 0 {
		m.Buf[m.CursorR] = t
	} else {
		oldLine := m.Buf[m.CursorR]
		newLine := ""
		if m.CursorC > 0 {
			newLine += oldLine[:m.CursorC]
		}
		newLine += t
		if m.CursorC < len(oldLine) {
			newLine += oldLine[m.CursorC:]
		}

		m.Buf[m.CursorR] = newLine
	}

	m.CursorC++
	return true
}
