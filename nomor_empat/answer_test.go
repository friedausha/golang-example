package nomor_empat_test

import (
	"fmt"
	"golang-example/nomor_empat"
	"testing"
)

func TestFindExponentOfTwo(t *testing.T) {
	err := fmt.Errorf("bukan pangkat dari 2")
	var tests = []struct {
		a int
		want int
		expectedErr error
	}{
		{2, 1, nil},
		{8, 3, nil},
		{20, 0, err},
		{100, 0, err},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("%d,%d", tc.a, tc.want)
		t.Run(testname, func(t *testing.T) {
			ans, err := nomor_empat.FindExponentOfTwo(tc.a)
			if ans != tc.want {
				t.Errorf("got %d, want %d", ans, tc.want)
			}
			if tc.expectedErr != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("wrong error")
			}
		})

	}
}