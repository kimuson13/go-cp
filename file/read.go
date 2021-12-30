package file

import (
	"fmt"
	"io/ioutil"
)

func ReadCopy(copyFiles []string) ([][]byte, error) {
	copies := make([][]byte, len(copyFiles))
	for i, cf := range copyFiles {
		b, err := ReadFile(cf)
		if err != nil {
			return copies, fmt.Errorf("Copy err: %w", err)
		}

		copies[i] = b
	}

	return copies, nil
}

func ReadFile(fileName string) ([]byte, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return b, fmt.Errorf("Read error: %w", err)
	}

	return b, nil
}
