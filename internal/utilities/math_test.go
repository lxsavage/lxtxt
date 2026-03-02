package utilities

import "testing"

func TestNumberWidth0(t *testing.T) {
	of := NumberWidth(0)

	if of != 1 {
		t.Fatalf("NumberWidth(0): expected 1, got %d", of)
	}
}

func TestNumberWidth1(t *testing.T) {
	of := NumberWidth(1)

	if of != 1 {
		t.Fatalf("NumberWidth(1): expected 1, got %d", of)
	}
}

func TestNumberWidth9(t *testing.T) {
	of := NumberWidth(9)

	if of != 1 {
		t.Fatalf("NumberWidth(9): expected 1, got %d", of)
	}
}

func TestNumberWidth10(t *testing.T) {
	of := NumberWidth(10)

	if of != 2 {
		t.Fatalf("NumberWidth(10): expected 2, got %d", of)
	}
}

func TestNumberWidth19(t *testing.T) {
	of := NumberWidth(19)

	if of != 2 {
		t.Fatalf("NumberWidth(19): expected 2, got %d", of)
	}
}

func TestNumberWidth10000(t *testing.T) {
	of := NumberWidth(10000)

	if of != 5 {
		t.Fatalf("NumberWidth(10000): expected 2, got %d", of)
	}
}

func TestIndentLevelNone(t *testing.T) {
	of := IndentLevel("asdf")

	if of != 0 {
		t.Fatalf("IndentLevel(\"asdf\"): expected 0, got %d", of)
	}
}

func TestIndentLevelNoneNoText(t *testing.T) {
	of := IndentLevel("")

	if of != 0 {
		t.Fatalf("IndentLevel(\"\"): expected 0, got %d", of)
	}
}

func TestIndentLevel3(t *testing.T) {
	of := IndentLevel("   asdf")

	if of != 3 {
		t.Fatalf("IndentLevel(\"   asdf\"): expected 3, got %d", of)
	}
}

func TestIndentLevel3NoText(t *testing.T) {
	of := IndentLevel("   ")

	if of != 3 {
		t.Fatalf("IndentLevel(\"   \"): expected 3, got %d", of)
	}
}
