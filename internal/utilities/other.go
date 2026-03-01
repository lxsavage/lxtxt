// utilities contains useful generalized mathematical helper functions
package utilities

import "os/exec"

func ShowMan() {
	call := exec.Command("man", "lxtxt")
	call.Run()
}
