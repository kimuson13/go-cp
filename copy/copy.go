package copy

import (
	"errors"
	"fmt"
)

var ErrTooShort = errors.New("input args too short, need more than 2 args")

func Run(args []string) error {
	if len(args) < 2 {
		return ErrTooShort
	}
	copyFiles := args[:len(args)-1]
	pasteDir := args[len(args)-1]
	fmt.Println("copyFiles: ", copyFiles)
	fmt.Println("pasteDir: ", pasteDir)

	return nil
}
