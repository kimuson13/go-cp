package utils_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/kimuson13/go-cp/utils"
)

func TestNew(t *testing.T) {
	expected := utils.Std{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}
	actual := utils.New()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("New want utils.Std{In: %v, Out: %v, ErrOut: %v}\nBut got {In: %v, Out: %v, ErrOut, %v}",
			expected.In, expected.Out, expected.ErrOut, actual.In, actual.Out, actual.ErrOut)
	}
}
