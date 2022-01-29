package copy

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kimuson13/go-cp/validate"
)

var (
	ErrNotExistCopyFile = errors.New("not exist copy file")
	ErrTooShort         = errors.New("input args too short, need more than 2 args")
)

func Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("go-cp: %w", ErrTooShort)
	}
	copyFiles := make([]string, len(args)-1)
	for i, arg := range args[:len(args)-1] {
		absolutePath, err := filepath.Abs(arg)
		if err != nil {
			return fmt.Errorf("go-cp: copyFiles: %w", err)
		}
		copyFiles[i] = absolutePath
	}

	for _, file := range copyFiles {
		if !validate.Exists(file) {
			return fmt.Errorf("go-cp: %w: %s", ErrNotExistCopyFile, file)
		}
	}

	pasteDir, err := filepath.Abs(args[len(args)-1])
	if err != nil {
		return fmt.Errorf("go-cp: pasteDir: %w", err)
	}

	if err := validate.ExistSameFileInDir(pasteDir, copyFiles); err != nil {
		return fmt.Errorf("go-cp: %w", err)
	}

	for _, file := range copyFiles {
		if err := MakeCopy(file, pasteDir); err != nil {
			return err
		}
	}

	return nil
}

func MakeCopy(fileName, pasteDir string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("MakeCopy: ReadFile: %w", err)
	}

	if err := os.WriteFile(filepath.Join(pasteDir, filepath.Base(fileName)), file, 0777); err != nil {
		return fmt.Errorf("MakeCopy: WriteFile: %w", err)
	}

	return nil
}
