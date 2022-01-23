package validate_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/kimuson13/go-cp/validate"
)

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
			got := validate.Exists(tt.path)
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

	if err := validate.ExistSameFileInDir(path, files); err != nil {
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
			if err := validate.ExistSameFileInDir(tt.path, tt.copyFiles); err == nil {
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
			got := validate.ExistFileName(tt.target, tt.name)
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
