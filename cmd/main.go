package main

import (
	"flag"
)

var (
	b = flag.Bool("b", false, "make a backup of each existing destination file")
	i = flag.Bool("i", false, "prompt before overwrite")
	f = flag.Bool("f", false, "force overwrite")
	r = flag.Bool("r", false, "copy directory")
)

func main() {
	flag.Parse()
}
