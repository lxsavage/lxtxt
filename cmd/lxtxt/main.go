package main

import (
	"fmt"
	"log"
	"lxsavage/lxtxt/internal/fileio"
	"lxsavage/lxtxt/internal/ui"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("specify a file to open")
	}

	path := os.Args[1]
	buf := []string{""}
	if fbuf, err := fileio.LoadFileBuf(path); err == nil {
		buf = fbuf
	}

	if err := ui.Exec(path, buf); err != nil {
		log.Fatalf("LXTXT runtime exception: %v", err)
	}
}
