package main

import (
	"flag"
	"fmt"
	"log"
	"lxsavage/lxtxt/internal/fileio"
	"lxsavage/lxtxt/internal/ui"
	"os"
)

const Version = "localbuild"

func main() {
	showVersion := flag.Bool("version", false, "gets the version of LXTXT")
	flag.Parse()

	if *showVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		fmt.Println("specify a file to open")
	}

	path := flag.Arg(0)
	buf := []string{""}
	if fbuf, err := fileio.LoadFileBuf(path); err == nil {
		buf = fbuf
	}

	if err := ui.Exec(path, buf); err != nil {
		log.Fatalf("LXTXT runtime exception: %v", err)
	}
}
