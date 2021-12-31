package validate

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/kimuson13/go-cp/utils"
)

var (
	ErrDeniedToOverwrite     = errors.New("overwrite denied")
	ErrNoOverwritePermission = errors.New("no permission to overwrite this file")
	ErrNotDir                = errors.New("if you want to copy more than two files, you need to specify a directory")
	ErrTooShort              = errors.New("need more than two args(want to copy files and the place or name you want to paste)")
)

var (
	InteractiveStart    string = "already exists file, you want to overwrite? [y/n]"
	InteractiveContinue string = "please type y or n"
)

func Args(inputArgs []string, overwritePerm, isInteractive bool, std utils.Std) error {
	if len(inputArgs) == 1 {
		return fmt.Errorf("Validation Err: %w", ErrTooShort)
	}

	copyFiles := inputArgs[:len(inputArgs)-1]
	paste := inputArgs[len(inputArgs)-1]

	for _, cf := range copyFiles {
		if _, err := os.Stat(cf); err != nil {
			return fmt.Errorf("Validation Err: %w", err)
		}
	}

	//ペースト先がファイルとして存在した場合、上書き許可があるかないかの確認
	f, err := os.Stat(paste)
	if err == nil {
		if len(copyFiles) >= 2 && !f.IsDir() {
			return fmt.Errorf("Validation Err: %w", ErrNotDir)
		}

		if overwritePerm {
			if err := os.Remove(paste); err != nil {
				return err
			}
		}

		if !f.IsDir() && !overwritePerm {
			if err := Interactive(isInteractive, std); err != nil {
				return fmt.Errorf("Validation Err: %w", err)
			}

			if err := os.Remove(paste); err != nil {
				return err
			}
		}
	}

	return nil
}

func Interactive(isInteracitve bool, std utils.Std) error {
	if !isInteracitve {
		return fmt.Errorf("Interactive Err: %w", ErrNoOverwritePermission)
	}

	if _, err := fmt.Fprintln(std.Out, InteractiveStart); err != nil {
		return fmt.Errorf("unexpected err at interactive: %w", err)
	}
	scanner := bufio.NewScanner(std.In)
	for scanner.Scan() {
		isPerm := scanner.Text()
		switch isPerm {
		case "y":
			return nil
		case "Y":
			return nil
		case "n":
			return fmt.Errorf("Interactiver Err: %v", ErrDeniedToOverwrite)
		default:
			if _, err := fmt.Fprintln(std.Out, InteractiveContinue); err != nil {
				return fmt.Errorf("unexpected err at interactive: %w", err)
			}
		}
	}

	return nil
}
