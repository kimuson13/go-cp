package utils

import (
	"io"
	"os"
)

type Std struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer
}

func New() Std {
	return Std{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}
}
