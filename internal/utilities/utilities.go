package utilities

func NumWidth(n int) int {
	w := 1
	for n > 0 {
		n /= 10
		w++
	}

	return w
}

func IndentLevel(s string) int {
	c := 0
	for ; c < len(s) && s[c] == ' '; c++ {
	}
	return c
}
