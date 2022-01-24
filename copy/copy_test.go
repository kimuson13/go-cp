package copy_test

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/kimuson13/go-cp/copy"
)

func TestRunWithCopyFileByRelativePath(t *testing.T) {
	expected := "hoge hoge huga"
	if err := os.WriteFile("test", []byte(expected), 0777); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir("hoge", 0777); err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir("hoge"); err != nil {
		t.Fatal(err)
	}

	if err := os.Mkdir("huga", 0777); err != nil {
		t.Fatal(err)
	}

	if err := copy.Run([]string{"../test", "huga"}); err != nil {
		t.Error(err)
	}

	got, err := os.ReadFile("huga/test")
	if err != nil {
		t.Error(err)
	}
	if string(got) != expected {
		t.Errorf("got = %s, want = %s", string(got), expected)
	}

	abs, err := filepath.Abs("..")
	if err != nil {
		t.Error(err)
	}
	if err := os.Chdir(abs); err != nil {
		t.Error(err)
	}

	if err := os.RemoveAll("hoge"); err != nil {
		t.Error(err)
	}
	if err := os.Remove("test"); err != nil {
		t.Error(err)
	}
}

func TestRun(t *testing.T) {
	expected1 := "hoge hoge hoge"
	if err := os.WriteFile("test2", []byte(expected1), 0777); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test2")

	expected2 := "hoge hoge hoge hoge"
	if err := os.WriteFile("test3", []byte(expected2), 0777); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test3")

	tempD1, err := os.MkdirTemp("./", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempD1)

	tempD2, err := os.MkdirTemp(".", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempD2)

	tests := map[string]struct {
		inputArgs []string
		want      []string
	}{
		"case_input_one_file": {S(t, "test2", tempD1), S(t, expected1)},
		"case_input_2_files":  {S(t, "test2", "test3", tempD2), S(t, expected1, expected2)},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			if err := copy.Run(tt.inputArgs); err != nil {
				t.Error(err)
			}

			copyFiles := tt.inputArgs[:len(tt.inputArgs)-1]
			for i, v := range copyFiles {
				got, err := os.ReadFile(v)
				if err != nil {
					t.Error(err)
				}

				if string(got) != tt.want[i] {
					t.Errorf("want = %s, got = %s", tt.want[i], got)
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
		"case_not_exist_dir":     {S(t, "hello", "hoge"), true},
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

func S(t *testing.T, args ...string) []string {
	val := make([]string, len(args))
	for i, arg := range args {
		val[i] = arg
	}

	return val
}
