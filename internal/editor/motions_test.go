package editor

import "testing"

func hasValidVerticalScrollPosition(m Model) bool {
	return m.CursorR >= m.ScrollBaseR && m.CursorR < m.ScrollBaseR+m.height
}
func hasValidHorizontalScrollPosition(m Model) bool {
	return m.CursorC >= m.ScrollBaseC && m.CursorC < m.ScrollBaseC+m.width
}

func TestRepeatMotionNegativeTimes(t *testing.T) {
	m := &Model{}

	motionWasRun := false
	motion := func() bool {
		motionWasRun = true
		return false
	}

	_ = m.RepeatMotion(-1, motion)

	if motionWasRun {
		t.Fatal("RepeatMotion(Model, -1, Motion): expected not to run, but ran at least once")
	}
}

func TestRepeatMotionTenTimes(t *testing.T) {
	m := &Model{}

	runCount := 0
	motion := func() bool {
		runCount++
		return true
	}

	_ = m.RepeatMotion(10, motion)

	if runCount != 10 {
		t.Fatalf("RepeatMotion(Model, -1, Motion): expected 10 executions, had %d", runCount)
	}
}
func TestRepeatMotionEarlyReturn(t *testing.T) {
	m := &Model{}

	runCount := 0
	motion := func() bool {
		runCount++
		return runCount < 5
	}

	_ = m.RepeatMotion(10, motion)

	if runCount != 5 {
		t.Fatalf("RepeatMotion(Model, 10, MotionThatFailsAfter5Executions): expected 5 executions, had %d", runCount)
	}
}

func TestCursorUpValidPosition(t *testing.T) {
	m := &Model{
		Buf:     []string{"", ""},
		CursorR: 1,
	}

	if ok := m.CursorUp(); !ok {
		t.Fatalf("CursorUp(Model): expected return true, got false")
	}

	if m.CursorR != 0 {
		t.Fatalf("CursorUp(Model): expected m.CursorR = 0, got %d", m.CursorR)
	}
}

func TestCursorUpInvalidPosition(t *testing.T) {
	m := &Model{
		Buf:     []string{"", ""},
		CursorR: 0,
	}

	if ok := m.CursorUp(); ok {
		t.Fatalf("CursorUp(Model): expected return false, got true")
	}

	if m.CursorR != 0 {
		t.Fatalf("CursorUp(Model): expected m.CursorR = 0, got %d", m.CursorR)
	}
}

func TestCursorUpScroll(t *testing.T) {
	m := &Model{
		Buf:         []string{"", "", ""},
		CursorR:     2,
		ScrollBaseR: 2,
	}

	if ok := m.CursorUp(); !ok {
		t.Fatalf("CursorUp(Model): expected return true, got false")
	}

	if m.ScrollBaseR != 1 {
		t.Fatalf("CursorUp(Model): expected m.ScrollBaseR = 1, got %d", m.ScrollBaseR)
	}
}

func TestCursorUpHorizontalAdjust(t *testing.T) {
	m := &Model{
		Buf:     []string{"a", "aa", ""},
		CursorR: 1,
		CursorC: 2,
	}

	if ok := m.CursorUp(); !ok {
		t.Fatalf("CursorUp(Model): expected return true, got false")
	}

	if m.CursorC != 1 {
		t.Fatalf("CursorUp(Model): expected m.CursorC = 0, got %d", m.ScrollBaseR)
	}
}

func TestCursorDownValidPosition(t *testing.T) {
	m := &Model{
		Buf:     []string{"", ""},
		CursorR: 0,
	}

	if ok := m.CursorDown(); !ok {
		t.Fatalf("CursorDown(Model): expected return true, got false")
	}

	if m.CursorR != 1 {
		t.Fatalf("CursorDown(Model): expected m.CursorR = 1, got %d", m.CursorR)
	}
}

func TestCursorDownInvalidPosition(t *testing.T) {
	m := &Model{
		Buf:     []string{"", ""},
		CursorR: 1,
	}

	if ok := m.CursorDown(); ok {
		t.Fatalf("CursorDown(Model): expected return false, got true")
	}

	if m.CursorR != 1 {
		t.Fatalf("CursorDown(Model): expected m.CursorR = 1, got %d", m.CursorR)
	}
}

func TestCursorDownScroll(t *testing.T) {
	m := &Model{
		Buf:         []string{"", "", ""},
		CursorR:     1,
		ScrollBaseR: 1,
		height:      1,
	}

	if ok := m.CursorDown(); !ok {
		t.Fatalf("CursorDown(Model): expected return true, got false")
	}

	if m.ScrollBaseR != 2 {
		t.Fatalf("CursorDown(Model): expected m.ScrollBaseR = 2, got %d", m.ScrollBaseR)
	}
}

func TestCursorDownHorizontalAdjust(t *testing.T) {
	m := &Model{
		Buf:     []string{"", "aa", "a"},
		CursorR: 1,
		CursorC: 2,
	}

	if ok := m.CursorDown(); !ok {
		t.Fatalf("CursorDown(Model): expected return true, got false")
	}

	if m.CursorC != 1 {
		t.Fatalf("CursorDown(Model): expected m.CursorC = 0, got %d", m.ScrollBaseR)
	}
}

func TestCursorLeftValidPosition(t *testing.T) {
	m := &Model{
		Buf:     []string{"aaaaa"},
		CursorR: 0,
		CursorC: 2,
	}

	if ok := m.CursorLeft(); !ok {
		t.Fatalf("CursorLeft(Model): expected return true, got false")
	}

	if m.CursorC != 1 {
		t.Fatalf("CursorLeft(Model): expected m.CursorC = 1, got %d", m.ScrollBaseC)
	}
}

func TestCursorLeftInvalidPosition(t *testing.T) {
	m := &Model{
		Buf:     []string{"aaaaa"},
		CursorR: 0,
		CursorC: 0,
	}

	if ok := m.CursorLeft(); ok {
		t.Fatalf("CursorLeft(Model): expected return false, got true")
	}

	if m.CursorC != 0 {
		t.Fatalf("CursorLeft(Model): expected m.CursorC = 0, got %d", m.ScrollBaseC)
	}
}

func TestCursorRightValidPosition(t *testing.T) {
	m := &Model{
		Buf:     []string{"aaaaa"},
		CursorR: 0,
		CursorC: 1,
	}

	if ok := m.CursorRight(); !ok {
		t.Fatalf("CursorLineEnd(Model): expected return true, got false")
	}

	if m.CursorC != 2 {
		t.Fatalf("CursorLineEnd(Model): expected m.CursorC = 2, got %d", m.ScrollBaseC)
	}
}

func TestCursorRightInvalidPosition(t *testing.T) {
	m := &Model{
		Buf:     []string{"aaaaa"},
		CursorR: 0,
		CursorC: 5,
	}

	if ok := m.CursorRight(); ok {
		t.Fatalf("CursorLineEnd(Model): expected return false, got true")
	}

	if m.CursorC != 5 {
		t.Fatalf("CursorLineEnd(Model): expected m.CursorC = 5, got %d", m.ScrollBaseC)
	}
}

func TestCursorRightWithScroll(t *testing.T) {
	m := &Model{
		Buf:         []string{"aaaaa"},
		CursorR:     0,
		CursorC:     1,
		ScrollBaseC: 1,
		width:       4, // 1 editor width + (gutter: 1 digit + 2 padding)
	}

	if ok := m.CursorRight(); !ok {
		t.Fatalf("CursorLineEnd(Model): expected return true, got false")
	}

	if m.ScrollBaseC != 2 {
		t.Fatalf("CursorLineEnd(Model): expected m.ScrollBaseC = 2, got %d", m.ScrollBaseC)
	}
}

func TestCursorLineStartValidPosition(t *testing.T) {
	m := &Model{
		CursorC: 2,
	}

	if ok := m.CursorLineStart(); !ok {
		t.Fatalf("CursorLineStart(Model): expected return true, got false")
	}

	if m.CursorC != 0 {
		t.Fatalf("CursorLineStart(Model): expected m.CursorC = 0, got %d", m.ScrollBaseR)
	}
}

func TestCursorLineStartInvalidPosition(t *testing.T) {
	m := &Model{
		CursorC: 0,
	}

	if ok := m.CursorLineStart(); ok {
		t.Fatalf("CursorLineStart(Model): expected return false, got true")
	}

	if m.CursorC != 0 {
		t.Fatalf("CursorLineStart(Model): expected m.CursorC = 0, got %d", m.ScrollBaseR)
	}
}

func TestCursorLineStartWithScroll(t *testing.T) {
	m := &Model{
		CursorC:     2,
		ScrollBaseC: 1,
	}

	if ok := m.CursorLineStart(); !ok {
		t.Fatalf("CursorLineStart(Model): expected return true, got false")
	}

	if m.ScrollBaseC != 0 {
		t.Fatalf("CursorLineStart(Model): expected m.ScrollBaseC = 0, got %d", m.ScrollBaseR)
	}
}
func TestCursorLineEndValidPosition(t *testing.T) {
	m := &Model{
		Buf:     []string{"aaaaa"},
		CursorR: 0,
		CursorC: 2,
	}

	if ok := m.CursorLineEnd(); !ok {
		t.Fatalf("CursorLineEnd(Model): expected return true, got false")
	}

	if m.CursorC != 5 {
		t.Fatalf("CursorLineEnd(Model): expected m.CursorC = 5, got %d", m.CursorC)
	}
}

func TestCursorLineEndInvalidPosition(t *testing.T) {
	m := &Model{
		Buf:     []string{"aaaaa"},
		CursorR: 0,
		CursorC: 5,
	}

	if ok := m.CursorLineEnd(); ok {
		t.Fatalf("CursorLineEnd(Model): expected return false, got true")
	}

	if m.CursorC != 5 {
		t.Fatalf("CursorLineEnd(Model): expected m.CursorC = 5, got %d", m.CursorC)
	}
}

func TestCursorLineEndWithScroll(t *testing.T) {
	m := &Model{
		Buf:         []string{"aaaaa"},
		CursorR:     0,
		CursorC:     1,
		ScrollBaseC: 1,
		width:       4, // 1 editor width + (gutter: 1 digit + 2 padding)
	}

	if ok := m.CursorLineEnd(); !ok {
		t.Fatalf("CursorLineEnd(Model): expected return true, got false")
	}

	if m.ScrollBaseC != 5 {
		t.Fatalf("CursorLineEnd(Model): expected m.ScrollBaseC = 5, got %d", m.ScrollBaseC)
	}
}

func TestBackspaceValidInline(t *testing.T) {
	m := &Model{
		Buf:     []string{"abcdefg"},
		CursorC: 2,
	}

	if ok := m.Backspace(); !ok {
		t.Fatalf("Backspace(Model): expected return true, got false")
	}

	if m.Buf[0] != "acdefg" {
		t.Fatalf("Backspace(Model): expected m.Buf[0] = \"acdefg\", got \"%s\"", m.Buf[0])
	}
}

func TestBackspaceInvalidInline(t *testing.T) {
	m := &Model{
		Buf:     []string{"abcdefg"},
		CursorC: 0,
	}

	if ok := m.Backspace(); ok {
		t.Fatalf("Backspace(Model): expected return false, got true")
	}

	if m.Buf[0] != "abcdefg" {
		t.Fatalf("Backspace(Model): expected m.Buf[0] = \"abcdefg\", got \"%s\"", m.Buf[0])
	}
}

func TestBackspaceMultiline(t *testing.T) {
	m := &Model{
		Buf:     []string{"abcdefg", "hijk"},
		CursorR: 1,
		CursorC: 0,
	}

	if ok := m.Backspace(); !ok {
		t.Fatalf("Backspace(Model): expected return true, got false")
	}

	if len(m.Buf) != 1 {
		t.Fatalf("Backspace(Model): expected len(m.Buf) = 1, got \"%v\"", len(m.Buf))
	}

	if m.Buf[0] != "abcdefghijk" {
		t.Fatalf("Backspace(Model): expected m.Buf[0] = \"abcdefghijk\", got \"%s\"", m.Buf[0])
	}
}

func TestDeleteValidInline(t *testing.T) {
	m := &Model{
		Buf:     []string{"abcdefg"},
		CursorC: 0,
	}

	if ok := m.Delete(); !ok {
		t.Fatalf("Delete(Model): expected return true, got false")
	}

	if m.Buf[0] != "bcdefg" {
		t.Fatalf("Delete(Model): expected m.Buf[0] = \"bcdefg\", got \"%s\"", m.Buf[0])
	}
}

func TestDeleteLastCharEOB(t *testing.T) {
	m := &Model{
		Buf:     []string{"abcdefg"},
		CursorC: 7,
		CursorR: 0,
	}

	if ok := m.Delete(); ok {
		t.Fatalf("Delete(Model): expected return false, got true")
	}

	if m.Buf[0] != "abcdefg" {
		t.Fatalf("Delete(Model): expected m.Buf[0] = \"abcdefg\", got \"%s\"", m.Buf[0])
	}
}

func TestDeleteValidMultiline(t *testing.T) {
	m := &Model{
		Buf:     []string{"abcdefg", "hijk"},
		CursorC: 7,
		CursorR: 0,
	}

	if ok := m.Delete(); !ok {
		t.Fatalf("Delete(Model): expected return true, got false")
	}

	if len(m.Buf) != 1 {
		t.Fatalf("Delete(Model): expected len(m.Buf) = 1, got \"%d\"", len(m.Buf))
	}

	if m.Buf[0] != "abcdefghijk" {
		t.Fatalf("Delete(Model): expected m.Buf[0] = \"abcdefghijk\", got \"%s\"", m.Buf[0])
	}
}

func TestDeleteLineFirstValid(t *testing.T) {
	m := &Model{
		Buf:     []string{"abcdefg", "hijk"},
		CursorR: 0,
	}

	if ok := m.DeleteLine(); !ok {
		t.Fatalf("DeleteLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 1 {
		t.Fatalf("DeleteLine(Model): expected len(m.Buf) = 1, got \"%d\"", len(m.Buf))
	}

	if m.Buf[0] != "hijk" {
		t.Fatalf("DeleteLine(Model): expected m.Buf[0] = \"hijk\", got \"%s\"", m.Buf[0])
	}
}

func TestDeleteLineLastValid(t *testing.T) {
	m := &Model{
		Buf:     []string{"abcdefg", "hijk"},
		CursorR: 1,
	}

	if ok := m.DeleteLine(); !ok {
		t.Fatalf("DeleteLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 1 {
		t.Fatalf("DeleteLine(Model): expected len(m.Buf) = 1, got \"%d\"", len(m.Buf))
	}

	if m.CursorR != 0 {
		t.Fatalf("DeleteLine(Model): expected m.CursorR = 0, got %d", m.CursorR)
	}

	if m.Buf[0] != "abcdefg" {
		t.Fatalf("DeleteLine(Model): expected m.Buf[0] = \"abcdefg\", got \"%s\"", m.Buf[0])
	}
}

func TestDeleteLineWithScroll(t *testing.T) {
	m := &Model{
		Buf:         []string{"abcdefg", "hijk", "lmn"},
		CursorR:     2,
		ScrollBaseR: 2,
	}

	if ok := m.DeleteLine(); !ok {
		t.Fatalf("DeleteLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 2 {
		t.Fatalf("DeleteLine(Model): expected len(m.Buf) = 2, got \"%d\"", len(m.Buf))
	}

	if m.ScrollBaseR != 1 {
		t.Fatalf("DeleteLine(Model): expected m.ScrollBaseR = 1, got %d", m.ScrollBaseR)
	}
}

func TestDeleteOnlyLine(t *testing.T) {
	m := &Model{
		Buf:     []string{"a"},
		CursorR: 0,
	}

	if ok := m.DeleteLine(); !ok {
		t.Fatalf("DeleteLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 1 {
		t.Fatalf("DeleteLine(Model): expected len(m.Buf) = 1, got \"%d\"", len(m.Buf))
	}

	if m.Buf[0] != "" {
		t.Fatalf("DeleteLine(Model): expected m.Buf[0] = \"\", got \"%s\"", m.Buf[0])
	}
}

func TestDeleteLineMiddle(t *testing.T) {
	m := &Model{
		Buf:     []string{"a", "b", "c"},
		CursorR: 1,
	}

	if ok := m.DeleteLine(); !ok {
		t.Fatalf("DeleteLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 2 {
		t.Fatalf("DeleteLine(Model): expected len(m.Buf) = 2, got \"%d\"", len(m.Buf))
	}

	if m.Buf[0] != "a" || m.Buf[1] != "c" {
		t.Fatalf("DeleteLine(Model): expected m.Buf = {\"a\",\"c\"}, got \"%s\"", m.Buf[0])
	}
}

func TestDeleteLineInvalid(t *testing.T) {
	m := &Model{
		Buf:     []string{""},
		CursorR: 0,
	}

	if ok := m.DeleteLine(); ok {
		t.Fatalf("DeleteLine(Model): expected return false, got true")
	}

	if len(m.Buf) != 1 {
		t.Fatalf("DeleteLine(Model): expected len(m.Buf) = 1, got \"%d\"", len(m.Buf))
	}

	if m.Buf[0] != "" {
		t.Fatalf("DeleteLine(Model): expected m.Buf[0] = \"\", got \"%s\"", m.Buf[0])
	}
}

func TestNewLineBaseNoIndent(t *testing.T) {
	m := &Model{
		Buf:     []string{"asdf"},
		CursorR: 0,
	}

	if ok := m.NewLine(0); !ok {
		t.Fatalf("NewLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 2 {
		t.Fatalf("NewLine(Model): expected len(m.Buf) = 2, got \"%v\"", len(m.Buf))
	}

	if m.Buf[0] != "" || m.Buf[1] != "asdf" {
		t.Fatalf("NewLine(Model): expected m.Buf = {\"\",\"asdf\"}, got \"%v\"", m.Buf)
	}
}

func TestNewLineBaseWithIndent(t *testing.T) {
	m := &Model{
		Buf:     []string{"asdf"},
		CursorR: 0,
	}

	if ok := m.NewLine(1); !ok {
		t.Fatalf("NewLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 2 {
		t.Fatalf("NewLine(Model): expected len(m.Buf) = 2, got \"%v\"", len(m.Buf))
	}

	if m.Buf[0] != "" || m.Buf[1] != " asdf" {
		t.Fatalf("NewLine(Model): expected m.Buf = {\"\",\" asdf\"}, got \"%v\"", m.Buf)
	}
}

func TestNewLineMiddleNoIndent(t *testing.T) {
	m := &Model{
		Buf:     []string{"asdf"},
		CursorR: 0,
		CursorC: 2,
	}

	if ok := m.NewLine(0); !ok {
		t.Fatalf("NewLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 2 {
		t.Fatalf("NewLine(Model): expected len(m.Buf) = 2, got \"%v\"", len(m.Buf))
	}

	if m.Buf[0] != "as" || m.Buf[1] != "df" {
		t.Fatalf("NewLine(Model): expected m.Buf = {\"as\",\"df\"}, got \"%v\"", m.Buf)
	}
}

func TestNewLineMiddleWithIndent(t *testing.T) {
	m := &Model{
		Buf:     []string{"asdf"},
		CursorR: 0,
		CursorC: 2,
	}

	if ok := m.NewLine(1); !ok {
		t.Fatalf("NewLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 2 {
		t.Fatalf("NewLine(Model): expected len(m.Buf) = 2, got \"%v\"", len(m.Buf))
	}

	if m.Buf[0] != "as" || m.Buf[1] != " df" {
		t.Fatalf("NewLine(Model): expected m.Buf = {\"as\",\" df\"}, got \"%v\"", m.Buf)
	}
}
func TestNewLineEndNoIndent(t *testing.T) {
	m := &Model{
		Buf:     []string{"asdf"},
		CursorR: 0,
		CursorC: 4,
	}

	if ok := m.NewLine(0); !ok {
		t.Fatalf("NewLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 2 {
		t.Fatalf("NewLine(Model): expected len(m.Buf) = 2, got \"%v\"", len(m.Buf))
	}

	if m.Buf[0] != "asdf" || m.Buf[1] != "" {
		t.Fatalf("NewLine(Model): expected m.Buf = {\"asdf\",\"\"}, got \"%v\"", m.Buf)
	}
}

func TestNewLineEndWithIndent(t *testing.T) {
	m := &Model{
		Buf:     []string{"asdf"},
		CursorR: 0,
		CursorC: 4,
	}

	if ok := m.NewLine(1); !ok {
		t.Fatalf("NewLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 2 {
		t.Fatalf("NewLine(Model): expected len(m.Buf) = 2, got \"%v\"", len(m.Buf))
	}

	if m.Buf[0] != "asdf" || m.Buf[1] != " " {
		t.Fatalf("NewLine(Model): expected m.Buf = {\"asdf\",\" \"}, got \"%v\"", m.Buf)
	}
}

func TestNewLineMiddleUpperSide(t *testing.T) {
	m := &Model{
		Buf:     []string{"asdf", "fdsa"},
		CursorC: 4,
		CursorR: 0,
	}

	if ok := m.NewLine(0); !ok {
		t.Fatalf("NewLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 3 {
		t.Fatalf("NewLine(Model): expected len(m.Buf) = 3, got \"%v\"", len(m.Buf))
	}

	if m.Buf[0] != "asdf" || m.Buf[1] != "" || m.Buf[2] != "fdsa" {
		t.Fatalf("NewLine(Model): expected m.Buf = {\"asdf\",\"\",\"fdsa\"}, got %v", m.Buf)
	}
}

func TestNewLineMiddleLowerSide(t *testing.T) {
	m := &Model{
		Buf:     []string{"asdf", "fdsa"},
		CursorC: 0,
		CursorR: 1,
	}

	if ok := m.NewLine(0); !ok {
		t.Fatalf("NewLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 3 {
		t.Fatalf("NewLine(Model): expected len(m.Buf) = 3, got \"%v\"", len(m.Buf))
	}

	if m.Buf[0] != "asdf" || m.Buf[1] != "" || m.Buf[2] != "fdsa" {
		t.Fatalf("NewLine(Model): expected m.Buf = {\"asdf\",\"\",\"fdsa\"}, got %v", m.Buf)
	}
}

func TestNewLineWithScroll(t *testing.T) {
	m := &Model{
		Buf:         []string{"asdf", "fdsa"},
		CursorR:     1,
		CursorC:     4,
		ScrollBaseR: 1,
		height:      1,
	}

	if ok := m.NewLine(0); !ok {
		t.Fatalf("NewLine(Model): expected return true, got false")
	}

	if len(m.Buf) != 3 {
		t.Fatalf("NewLine(Model): expected len(m.Buf) = 3, got \"%v\"", len(m.Buf))
	}

	if m.ScrollBaseR != 2 {
		t.Fatalf("NewLine(Model): expected m.ScrollBaseR = 2, got %d", m.ScrollBaseR)
	}
}

func TestInsertTextEmptyLine(t *testing.T) {
	m := &Model{
		Buf:     []string{""},
		CursorR: 0,
		CursorC: 0,
	}

	if ok := m.InsertText("a"); !ok {
		t.Fatalf("InsertText(Model): expected return true, got false")
	}

	if m.Buf[0] != "a" {
		t.Fatalf("InsertText(Model): expected m.Buf[0] = \"a\", got \"%s\"", m.Buf[0])
	}
}

func TestInsertTextBase(t *testing.T) {
	m := &Model{
		Buf:     []string{"sdf"},
		CursorR: 0,
		CursorC: 0,
	}

	if ok := m.InsertText("a"); !ok {
		t.Fatalf("InsertText(Model): expected return true, got false")
	}

	if m.Buf[0] != "asdf" {
		t.Fatalf("InsertText(Model): expected m.Buf[0] = \"asdf\", got \"%s\"", m.Buf[0])
	}
}

func TestInsertTextMiddle(t *testing.T) {
	m := &Model{
		Buf:     []string{"asf"},
		CursorR: 0,
		CursorC: 2,
	}

	if ok := m.InsertText("d"); !ok {
		t.Fatalf("InsertText(Model): expected return true, got false")
	}

	if m.Buf[0] != "asdf" {
		t.Fatalf("InsertText(Model): expected m.Buf[0] = \"asdf\", got \"%s\"", m.Buf[0])
	}
}

func TestInsertTextEnd(t *testing.T) {
	m := &Model{
		Buf:     []string{"asd"},
		CursorR: 0,
		CursorC: 3,
	}

	if ok := m.InsertText("f"); !ok {
		t.Fatalf("InsertText(Model): expected return true, got false")
	}

	if m.Buf[0] != "asdf" {
		t.Fatalf("InsertText(Model): expected m.Buf[0] = \"asdf\", got \"%s\"", m.Buf[0])
	}
}
