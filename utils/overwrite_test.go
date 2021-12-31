package utils_test

import (
	"os"
	"testing"

	"github.com/kimuson13/go-cp/utils"
)

func TestOverwrite(t *testing.T) {
	tempFile, closeFunc := SetUp(t)
	if err := utils.Overwrite(tempFile); err != nil {
		t.Error(err)
	}
	defer closeFunc()

	if err := utils.Overwrite(tempFile); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(tempFile); err == nil {
		t.Error("this test not want file")
	}
}

func SetUp(t *testing.T) (string, func()) {
	t.Helper()
	tempDir, err := os.MkdirTemp("./", "test")
	if err != nil {
		t.Fatal(err)
	}
	tempFile, err := os.CreateTemp(tempDir, "temp")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tempFile.Name())
	defer func() {
		if err := tempFile.Close(); err != nil {
			t.Log(err)
		}
	}()

	closeFunc := func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Log(err)
		}
	}

	return tempFile.Name(), closeFunc
}
