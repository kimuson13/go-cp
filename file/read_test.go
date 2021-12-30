package file_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/kimuson13/go-cp/file"
)

func TestReadCopy(t *testing.T) {
	tests := map[string]struct {
		input []string
	}{
		"read_one_file": {S(t, "hello world")},
		"read_two_file": {S(t, "hello world", "hoge")},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			tempFileNames, expected, closeFunc := MakeTemps(t, tt.input...)

			actual, err := file.ReadCopy(tempFileNames)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("ReadCopy expected: %v actual: %v", expected, actual)
			}

			closeFunc()
		})
	}
}

func TestReadFile(t *testing.T) {
	tests := map[string]struct {
		input string
	}{
		"input_value": {"hello world"},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			fileNames, expected, closeFunc := MakeTemps(t, tt.input)
			actual, err := file.ReadFile(fileNames[0])
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(actual, expected[0]) {
				t.Errorf("ReadFile want %v, but got %v", expected[0], actual)
			}

			closeFunc()
		})
	}
}

func S(t *testing.T, values ...string) []string {
	val := make([]string, len(values))
	for i, v := range values {
		val[i] = v
	}

	return val
}

func MakeTemps(t *testing.T, values ...string) ([]string, [][]byte, func()) {
	tempFileNames := make([]string, len(values))
	expected := make([][]byte, len(values))
	tempDir, err := os.MkdirTemp("./", "test")
	if err != nil {
		t.Fatal(err)
	}

	for i, v := range values {
		tempFile, err := os.CreateTemp(tempDir, "temp")
		if err != nil {
			t.Fatal(err)
		}
		tempFileNames[i] = tempFile.Name()

		input := []byte(v)
		expected[i] = input

		if _, err := tempFile.Write(input); err != nil {
			t.Fatal(err)
		}

		if err := tempFile.Close(); err != nil {
			t.Error(err)
		}
	}

	closeFunc := func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Error(err)
		}
	}

	return tempFileNames, expected, closeFunc
}
