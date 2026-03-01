package utilities

// NumberWidth determines how many digits wide n is
func NumberWidth(n int) int {
	w := 1
	for n > 0 {
		n /= 10
		w++
	}

	return w
}

// IndentLevel determines how many spaces of indentation s has
func IndentLevel(s string) int {
	c := 0
	for ; c < len(s) && s[c] == ' '; c++ {
	}
	return c
}
