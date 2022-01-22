package copy_test

import (
	"testing"

	"github.com/kimuson13/go-cp/copy"
)

func TestRun(t *testing.T) {
	tests := map[string]struct {
		inputArgs []string
	}{
		"case_input_2_args":      {S(t, "hello", "hoge")},
		"case_input_over_2_args": {S(t, "hello", "hoge", "hoge")},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			if err := copy.Run(tt.inputArgs); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestRunFailed(t *testing.T) {
	tests := map[string]struct {
		inputArgs []string
		want      error
	}{
		"case_too_short_args(1)": {S(t, "hello"), copy.ErrTooShort},
		"case_too_short_args(0)": {S(t), copy.ErrTooShort},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			got := copy.Run(tt.inputArgs).Error()
			if got != tt.want.Error() {
				t.Errorf("want = %s, but got = %s", tt.want.Error(), got)
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
