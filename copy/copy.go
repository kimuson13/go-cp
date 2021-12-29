package copy

import (
	"github.com/kimuson13/go-cp/utils"
	"github.com/kimuson13/go-cp/validate"
)

func Run(b, f, i, r bool, args []string, std utils.Std) error {
	if err := validate.Args(args, f, i, std); err != nil {
		return err
	}

	return nil
}
