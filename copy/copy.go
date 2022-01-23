package copy

import (
	"errors"
	"fmt"
	"io"
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
		return ErrTooShort
	}
	copyFiles := args[:len(args)-1]
	for _, file := range copyFiles {
		if !validate.Exists(file) {
			return fmt.Errorf("%w: %s", ErrNotExistCopyFile, file)
		}
	}
	pasteDir := args[len(args)-1]
	if err := validate.ExistSameFileInDir(pasteDir, copyFiles); err != nil {
		return err
	}

	for _, file := range copyFiles {
		if err := MakeCopy(file, pasteDir); err != nil {
			return err
		}
	}

	return nil
}

func MakeCopy(fileName, pasteDir string) error {
	fileInfo, err := os.Open(fileName)
	if err != nil {
		return err
	}

	pasteFile, err := os.Create(filepath.Join(pasteDir, fileInfo.Name()))
	if err != nil {
		return err
	}

	b, err := io.ReadAll(fileInfo)
	if err != nil {
		return err
	}

	if _, err := pasteFile.Write(b); err != nil {
		return err
	}

	return nil
}

// func Exists(path string) bool {
// 	_, err := os.Stat(path)

// 	return !os.IsNotExist(err)
// }

// func ExistSameFileInDir(path string, copyFiles []string) error {
// 	files, err := os.ReadDir(path)
// 	if err != nil {
// 		return err
// 	}

// 	for _, file := range files {
// 		if !file.IsDir() {
// 			if ExistFileName(copyFiles, file.Name()) {
// 				return ErrAlreadyExist
// 			}
// 		}
// 	}

// 	return nil
// }

// func ExistFileName(target []string, name string) bool {
// 	for _, v := range target {
// 		if v == name {
// 			return true
// 		}
// 	}

// 	return false
// }
