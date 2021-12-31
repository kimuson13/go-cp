package utils

import "os"

func Overwrite(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return nil
}
