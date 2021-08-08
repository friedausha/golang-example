package concat_string_test

import (
	"fmt"
	"golang-example/concat_string"
	"testing"
)

func TestConcatMultipleTimes(t *testing.T) {
	var tests = []struct {
		a, b        string
		want        string
		expectedErr error
	}{
		{"aku ", "uka", "aku uka", nil},
		{"aku", "aku", "", fmt.Errorf("string sama yaitu aku")},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("%s,%s", tc.a, tc.b)
		t.Run(testname, func(t *testing.T) {
			ans, _, _, err := concat_string.ConcatString(tc.a, tc.b)
			if ans != tc.want {
				t.Errorf("got %s, want %s", ans, tc.want)
			}
			if tc.expectedErr != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("wrong error")
			}
		})
	}
}
