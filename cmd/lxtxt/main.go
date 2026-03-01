package main

import (
	"flag"
	"fmt"
	"log"
	"lxsavage/lxtxt/internal/ui"
	"lxsavage/lxtxt/internal/utilities"
	"os"
)

const Version = "localbuild"

func main() {
	showVersion := flag.Bool("version", false, "gets the version of LXTXT")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "LXTXT Editor %s\n\n", Version)
		fmt.Fprintf(os.Stderr,
			"Usage: %s [file]\tedit a specific file\n"+
				"   or: %s\t\topen an empty buffer\n"+
				"   or: %s [arguments]\n\n",
			os.Args[0],
			os.Args[0],
			os.Args[0],
		)
		flag.PrintDefaults()
	}
	flag.Parse()

	if *showVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	path := ""
	buf := []string{""}
	if len(flag.Args()) > 0 {
		path = flag.Arg(0)
		if fbuf, err := utilities.LoadFileBuf(path); err == nil {
			buf = fbuf
		}
	}

	if err := ui.Exec(path, buf); err != nil {
		log.Fatalf("LXTXT runtime exception: %v", err)
	}
}
