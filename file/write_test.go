package file_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kimuson13/go-cp/file"
)

func TestCreateBackUp(t *testing.T) {
	tests := map[string]struct {
		inputFiles []string
		expected   []string
	}{
		"one_file_input": {S(t, "hello"), S(t, "hello~")},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			inputFiles, expected, closeFunc := CreateTempFile(t, tt.inputFiles, tt.expected)

			if err := file.CreateBakcUp(inputFiles); err != nil {
				t.Error(err)
			}

			for _, v := range expected {
				if _, err := os.Stat(v); err != nil {
					t.Error(err)
				}
			}

			closeFunc()
		})
	}
}

func CreateTempFile(t *testing.T, values []string, exepcted []string) ([]string, []string, func()) {
	files := []string{}
	expectedFiles := []string{}
	tempDir, err := os.MkdirTemp("./", "test")
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range values {
		file, err := os.Create(filepath.Join(tempDir, v))
		if err != nil {
			t.Fatal(err)
		}

		files = append(files, file.Name())
	}

	for _, v := range exepcted {
		newFile := filepath.Join(tempDir, v)
		expectedFiles = append(expectedFiles, newFile)
	}

	closeFunc := func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Log(err)
		}
	}

	return files, expectedFiles, closeFunc
}
