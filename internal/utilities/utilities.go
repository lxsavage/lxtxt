// utilities contains useful generalized mathematical helper functions
package utilities

import "os/exec"

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

func ShowMan() {
	call := exec.Command("man", "lxtxt")
	call.Run()
}
