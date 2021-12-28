package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kimuson13/go-cp/copy"
)

var (
	b = flag.Bool("b", false, "make a backup of each existing destination file")
	f = flag.Bool("f", false, "force overwrite")
	i = flag.Bool("i", false, "prompt before overwrite")
	r = flag.Bool("r", false, "copy directory")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if err := copy.Run(*b, *f, *i, *r, args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
