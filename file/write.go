package file

import (
	"fmt"
	"io/ioutil"
)

func CreateBakcUp(backUpFiles []string) error {
	for _, bf := range backUpFiles {
		data, err := ioutil.ReadFile(bf)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(fmt.Sprintf("%s~", bf), data, 0644); err != nil {
			return err
		}
	}

	return nil
}
