package validate

import (
	"errors"
	"os"
)

var (
	ErrAlreadyExist = errors.New("file already exist")
)

func Exists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func ExistSameFileInDir(path string, copyFiles []string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			if ExistFileName(copyFiles, file.Name()) {
				return ErrAlreadyExist
			}
		}
	}

	return nil
}

func ExistFileName(target []string, name string) bool {
	for _, v := range target {
		if v == name {
			return true
		}
	}

	return false
}
