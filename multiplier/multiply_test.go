package multiplier_test

import (
	"fmt"
	"golang-example/multiplier"
	"testing"
)

func TestAddMultipleTimes(t *testing.T) {
	var tests = []struct {
		a, b        int
		want        int
		expectedErr error
	}{
		{12, 13, 0, fmt.Errorf("12 kurang dari 13")},
		{10, 1, 10, nil},
		{22, 20, 440, nil},
		{100, -1, -100, nil},
		{-100, 90, 0, fmt.Errorf("-100 kurang dari 90")},
	}

	for _, tc := range tests {

		testname := fmt.Sprintf("%d,%d", tc.a, tc.b)
		t.Run(testname, func(t *testing.T) {
			ans, err := multiplier.Multiply(tc.a, tc.b)
			if ans != tc.want {
				t.Errorf("got %d, want %d", ans, tc.want)
			}
			if tc.expectedErr != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("wrong error")
			}
		})

	}
}
