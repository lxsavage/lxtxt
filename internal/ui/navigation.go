package ui

import "strings"

func cursorUp(m model) model {
	if m.cursorR <= 0 {
		return m
	}

	if m.rScrollBase > 0 && m.cursorR == m.rScrollBase {
		m.rScrollBase--
	}

	m.cursorR--
	if m.cursorC > len(m.buf[m.cursorR])-1 {
		m.cursorC = len(m.buf[m.cursorR])
	}
	return m
}

func cursorDown(m model) model {
	if m.cursorR >= len(m.buf)-1 {
		return m
	}

	m.cursorR++
	if m.cursorC > len(m.buf[m.cursorR])-1 {
		m.cursorC = len(m.buf[m.cursorR])
	}
	return m
}

func cursorLeft(m model) model {
	if m.cursorC > 0 {
		m.cursorC--
	}
	return m
}

func cursorRight(m model) model {
	if m.cursorC < len(m.buf[m.cursorR]) {
		m.cursorC++
	}
	return m
}

func cursorLineStart(m model) model {
	m.cursorC = 0
	return m
}

func cursorLineEnd(m model) model {
	m.cursorC = len(m.buf[m.cursorR])
	return m
}

func backspace(m model) model {
	if m.cursorC > 0 {
		oldLine := m.buf[m.cursorR]
		newLine := ""

		newLine += oldLine[:m.cursorC-1]
		if m.cursorC < len(oldLine) {
			newLine += oldLine[m.cursorC:]
		}
		m.buf[m.cursorR] = newLine
		m.cursorC--
		m = m.setDirty(true)
	} else if m.cursorR > 0 {
		oldLine := m.buf[m.cursorR]
		newBuf := append(m.buf[:m.cursorR], m.buf[m.cursorR+1:]...)
		m.cursorR--
		m.cursorC = len(m.buf[m.cursorR])

		newBuf[m.cursorR] += oldLine
		m.buf = newBuf
		m = m.setDirty(true)
	}
	return m
}

func deleteline(m model) model {
	if m.cursorR == 0 && len(m.buf) == 0 {
		return m
	}

	if m.cursorR == 0 {
		m.buf = m.buf[1:]
	} else {

		if m.cursorR < len(m.buf) {
			m.buf = append(m.buf[:m.cursorR], m.buf[m.cursorR+1:]...)
		} else {
			m.buf = append([]string(nil), m.buf[:m.cursorR]...)
		}

		if m.cursorR >= len(m.buf) {
			m.cursorR = len(m.buf)
			m = cursorUp(m)
		}
	}

	return m
}

func newline(m model, indent int) model {
	oldLine := ""
	newLine := ""
	if indent > 0 {
		newLine = strings.Repeat(" ", indent)
	}

	if m.cursorC > 0 {
		oldLine = m.buf[m.cursorR][:m.cursorC]
	}
	if m.cursorC < len(m.buf[m.cursorR])-1 {
		newLine += m.buf[m.cursorR][m.cursorC:]
	}

	var newBuf []string
	if m.cursorR > 0 {
		newBuf = append(newBuf, m.buf[:m.cursorR]...)
	}
	newBuf = append(newBuf, oldLine, newLine)

	if m.cursorR < len(m.buf)-1 {
		newBuf = append(newBuf, m.buf[m.cursorR+1:]...)
	}

	m.cursorR++
	if m.cursorC > len(newLine) {
		m.cursorC = len(newLine)
	}
	m.buf = newBuf
	m.cursorC = indent
	return m.setDirty(true)
}

func insertText(m model, t string) model {
	if len(m.buf[m.cursorR]) == 0 {
		m.buf[m.cursorR] = t
	} else {
		oldLine := m.buf[m.cursorR]
		newLine := ""
		if m.cursorC > 0 {
			newLine += oldLine[:m.cursorC]
		}
		newLine += t
		if m.cursorC < len(oldLine)-1 {
			newLine += oldLine[m.cursorC:]
		}

		m.buf[m.cursorR] = newLine
	}

	m.cursorC++
	return m.setDirty(true)
}
