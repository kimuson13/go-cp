package validate_test

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/kimuson13/go-cp/utils"
	"github.com/kimuson13/go-cp/validate"
)

func TestArgs(t *testing.T) {
	temps, toDir, closeFunc := CreateTemps(t)
	defer closeFunc()

	tests := map[string]struct {
		inputArgs     []string
		isInteracitve bool
		overwritePerm bool
		in            io.Reader
	}{
		"copy_file_to_dir":                   {S(t, temps[0], toDir), false, false, toIn(t, "n")},
		"copy_two_file_to_dir":               {S(t, temps[0], temps[1], toDir), false, false, toIn(t, "n")},
		"copy_file_to_not_exist_file":        {S(t, temps[0], "hello"), false, false, toIn(t, "n")},
		"copy_file_to_file_with_overwrite":   {S(t, temps[0], temps[1]), true, false, toIn(t, "n")},
		"copy_file_to_file_with_interactive": {S(t, temps[0], temps[1]), false, true, toIn(t, "y")},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			mock := mockStd(t, tt.in, os.Stdout, os.Stderr)
			if err := validate.Args(tt.inputArgs, tt.isInteracitve, tt.overwritePerm, mock); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestArgsFailed(t *testing.T) {
	temps, toDir, closeFunc := CreateTemps(t)
	defer closeFunc()

	tests := map[string]struct {
		inputArgs     []string
		isInteracitve bool
		overwritePerm bool
		stdin         io.Reader
	}{
		"too_short_error":               {S(t, "hello"), true, true, toIn(t, "y")},
		"not_stat_file":                 {S(t, "hello", toDir), true, true, toIn(t, "y")},
		"not_stat_file_with_two":        {S(t, temps[0], "hello", toDir), true, true, toIn(t, "y")},
		"no_interactive_permission":     {S(t, temps[0], temps[1]), false, false, toIn(t, "y")},
		"not_overwrite_because_input_n": {S(t, temps[0], temps[1]), true, false, toIn(t, "n")},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			mock := mockStd(t, tt.stdin, os.Stdout, os.Stderr)

			if err := validate.Args(tt.inputArgs, tt.isInteracitve, tt.overwritePerm, mock); err == nil {
				t.Errorf("expected err, but got no error")
			} else {
				t.Log(err)
			}
		})
	}
}

func TestInteractive(t *testing.T) {
	tests := map[string]struct {
		stdin    io.Reader
		expected string
	}{
		"input_y":           {toIn(t, "y"), toE(t, validate.InteractiveStart)},
		"input_large_y":     {toIn(t, "Y"), toE(t, validate.InteractiveStart)},
		"input_no_and_y":    {toIn(t, "\ny"), toE(t, validate.InteractiveStart, validate.InteractiveContinue)},
		"input_other_and_y": {toIn(t, "o\ny"), toE(t, validate.InteractiveStart, validate.InteractiveContinue)},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			out := new(bytes.Buffer)
			mock := mockStd(t, tt.stdin, out, os.Stderr)
			if err := validate.Interactive(true, mock); err != nil {
				t.Errorf("not expected err: %v", err)
			}
			if out.String() != tt.expected {
				t.Errorf("Interactive want: %v\nbut got: %v", tt.expected, out.String())
			}
		})
	}
}

func TestInteractiveFailed(t *testing.T) {
	tests := map[string]struct {
		in              io.Reader
		isInteracitve   bool
		expectedMessage string
		expectedErr     error
	}{
		"no_interactive_flag": {toIn(t, "n"), false, toE(t), validate.ErrNoOverwritePermission},
		"input_n":             {toIn(t, "n"), true, toE(t, validate.InteractiveStart), validate.ErrDeniedToOverwrite},
		"input_other_and_n":   {toIn(t, "o\nn"), true, toE(t, validate.InteractiveStart, validate.InteractiveContinue), validate.ErrDeniedToOverwrite},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			out := new(bytes.Buffer)
			mock := mockStd(t, tt.in, out, os.Stderr)

			if err := validate.Interactive(tt.isInteracitve, mock); err == nil {
				t.Error("expected err but got no err")
			} else {
				if out.String() != tt.expectedMessage {
					t.Errorf("Interactive expected message: %v\nbut got: %v", tt.expectedMessage, out.String())
				}
				if !errors.As(err, &tt.expectedErr) {
					t.Errorf("Interactice expected err: %v\nbut got: %v", tt.expectedErr.Error(), err.Error())
				}
			}
		})
	}
}

func CreateTemps(t *testing.T) ([]string, string, func()) {
	tempFileNames := make([]string, 2)

	dir, err := os.MkdirTemp("./", "test")
	if err != nil {
		t.Fatal(err)
	}

	f1, err := os.CreateTemp(dir, "test1")
	if err != nil {
		t.Fatal(err)
	}

	f2, err := os.CreateTemp(dir, "test2")
	if err != nil {
		t.Fatal(err)
	}

	dir2, err := os.MkdirTemp(dir, "to")
	if err != nil {
		t.Fatal(err)
	}

	tempFileNames[0], tempFileNames[1] = f1.Name(), f2.Name()

	closeFunc := func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Error(err)
		}
	}

	defer func() {
		if err := f1.Close(); err != nil {
			t.Error(err)
		}

		if err := f2.Close(); err != nil {
			t.Error(err)
		}
	}()

	return tempFileNames, dir2, closeFunc
}

func S(t *testing.T, args ...string) []string {
	values := []string{}
	for _, v := range args {
		values = append(values, v)
	}

	return values
}

func toIn(t *testing.T, input string) io.Reader {
	return bytes.NewReader([]byte(input))
}

func toE(t *testing.T, args ...string) string {
	value := ""
	for _, v := range args {
		value += v + "\n"
	}

	return value
}

func mockStd(t *testing.T, in io.Reader, out, errOut io.Writer) utils.Std {
	return utils.Std{
		In:     in,
		Out:    out,
		ErrOut: errOut,
	}
}
