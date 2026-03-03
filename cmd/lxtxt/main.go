package main

import (
	"flag"
	"fmt"
	"log"
	"lxsavage/lxtxt/internal/ui"
	"lxsavage/lxtxt/internal/utilities"
	"os"
	"time"
)

const Version = "localbuild"

func showExperimentsWarning(done chan<- any) {
	fmt.Println("WARNING: Experiments are enabled; unexpected or broken behavior may occur.")
	time.Sleep(time.Second * 2)
	done <- struct{}{}
}

func main() {
	experiments := flag.Bool("experiments", false, "enable incomplete/experimental features")
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

	msgDoneCh := make(chan any, 1)
	if *experiments {
		go showExperimentsWarning(msgDoneCh)
	}

	path := ""
	buf := []string{""}
	if len(flag.Args()) > 0 {
		path = flag.Arg(0)
		if fbuf, err := utilities.LoadFileBuf(path); err == nil {
			buf = fbuf
		}
	}

	if *experiments {
		<-msgDoneCh
	}

	if err := ui.Exec(path, buf, *experiments); err != nil {
		log.Fatalf("LXTXT runtime exception: %v", err)
	}
}
