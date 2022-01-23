package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kimuson13/go-cp/copy"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if err := copy.Run(args); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err.Error()); err != nil {
			log.Print(err)
		}
		os.Exit(1)
	}
}
