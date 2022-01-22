package copy_test

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/kimuson13/go-cp/copy"
)

func TestRun(t *testing.T) {
	tests := map[string]struct {
		inputArgs []string
	}{
		"case_input_2_args":      {S(t, "hello", t.TempDir())},
		"case_input_over_2_args": {S(t, "hello", "hoge", t.TempDir())},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			files := tt.inputArgs[:len(tt.inputArgs)-1]
			for _, v := range files {
				if _, err := os.Create(v); err != nil {
					t.Fatal(err)
				}
			}
			if err := copy.Run(tt.inputArgs); err != nil {
				t.Fatal(err)
			}

			for _, v := range files {
				if err := os.Remove(v); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestRunFailed(t *testing.T) {
	tests := map[string]struct {
		inputArgs []string
		needFile  bool
	}{
		"case_too_short_args(1)": {S(t, "hello"), false},
		"case_too_short_args(0)": {S(t), false},
		"case_not_exist_file":    {S(t, "hello", "hoge"), false},
		"case_not_exist_dir":     {S(t, "heelo", "hoge"), true},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			if tt.needFile {
				if _, err := os.Create(tt.inputArgs[0]); err != nil {
					t.Fatal(err)
				}
			}
			if err := copy.Run(tt.inputArgs); err == nil {
				t.Error("want err")
			}

			if tt.needFile {
				if err := os.Remove(tt.inputArgs[0]); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

func TestMakeCopy(t *testing.T) {
	input, err := os.Create("hoge")
	if _, err := input.Write([]byte("this is test")); err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	pasteDir := t.TempDir()

	if err := copy.MakeCopy(input.Name(), pasteDir); err != nil {
		t.Fatal(err)
	}

	got, err := os.Open(filepath.Join(pasteDir, input.Name()))
	if err != nil {
		t.Fatal(err)
	}

	if b, err := io.ReadAll(got); err != nil {
		t.Fatal(err)
	} else if string(b) != "this is test" {
		t.Errorf("want = this is test, but got %s", string(b))
	}

	if err := input.Close(); err != nil {
		t.Error(err)
	}
	if err := os.Remove(input.Name()); err != nil {
		t.Error(err)
	}
}

func TestExists(t *testing.T) {
	tests := map[string]struct {
		path string
		want bool
	}{
		"case_exist":     {t.TempDir(), true},
		"case_not_exist": {"hoge", false},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			got := copy.Exists(tt.path)
			if got != tt.want {
				t.Errorf("want = %v, but got = %v", tt.want, got)
			}
		})
	}
}

func TestExistFileInDir(t *testing.T) {
	path := t.TempDir()
	files := S(t, "hoge", "huga")
	if _, err := os.Create(filepath.Join(path, "temp")); err != nil {
		t.Fatal(err)
	}

	if err := copy.ExistSameFileInDir(path, files); err != nil {
		t.Error(err)
	}
}

func TestExistFileInDirFailed(t *testing.T) {
	tests := map[string]struct {
		path      string
		copyFiles []string
		needFile  bool
	}{
		"case_not_dir":    {"hoge", S(t, "hello"), false},
		"case_exist_fiel": {t.TempDir(), S(t, "hoge", "huga"), true},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			if tt.needFile {
				if _, err := os.Create(fmt.Sprintf("%s/%s", tt.path, tt.copyFiles[0])); err != nil {
					t.Fatal(err)
				}
			}
			if err := copy.ExistSameFileInDir(tt.path, tt.copyFiles); err == nil {
				t.Error("want err")
			}
		})
	}
}

func TestExistFileName(t *testing.T) {
	tests := map[string]struct {
		target []string
		name   string
		want   bool
	}{
		"case_not_exist": {S(t, "hoge", "huga"), "hello", false},
		"case_exist":     {S(t, "hoge", "huga"), "hoge", true},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			got := copy.ExistFileName(tt.target, tt.name)
			if got != tt.want {
				t.Errorf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}

func S(t *testing.T, args ...string) []string {
	val := make([]string, len(args))
	for i, arg := range args {
		val[i] = arg
	}

	return val
}
