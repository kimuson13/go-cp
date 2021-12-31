package file

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func CreateBakcUp(backUpFiles []string) error {
	for _, bf := range backUpFiles {
		data, err := ioutil.ReadFile(bf)
		if err != nil {
			return err
		}

		newFile := filepath.Base(bf[:len(bf)-len(filepath.Ext(bf))])
		if err := ioutil.WriteFile(fmt.Sprintf("%s~%s", newFile, filepath.Ext(bf)), data, 0644); err != nil {
			return err
		}
	}

	return nil
}
