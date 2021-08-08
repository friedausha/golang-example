package addition_test

import (
	"fmt"
	"golang-example/addition"
	"testing"
)

func TestAdd(t *testing.T) {
	actualAns, actualErr := addition.Add(2, 3)
	if actualErr.Error() != fmt.Errorf("kurang dari 10").Error() {
		t.Errorf("got actualErr %+v, expected %v ", actualErr, fmt.Errorf("kurang dari 10"))
	}
	if actualAns != 0 {
		t.Errorf("it returns actualAns, supposed to be error")
	}
}

func TestAddSuccess(t *testing.T) {
	actualAns, actualErr := addition.Add(20, 30)
	expectedAns := 50
	if actualErr !=  nil {
		t.Errorf("ada error padahal harusnya sukses")
	}
	if actualAns != expectedAns {
		t.Errorf("hasilnya %d padahal harusnya %d", actualAns, expectedAns)
	}
}

func TestAddMultipleTimes(t *testing.T) {
	err := fmt.Errorf("kurang dari 10")
	var tests = []struct {
		a, b int
		want int
		expectedErr error
	}{
		{12, 13, 25, nil},
		{10, 1, 0, err},
		{20, 22, 42, nil},
		{100, -1, 0, err},
		{-100, 90, 0, err},
	}

	for _, tc := range tests {
		testname := fmt.Sprintf("%d,%d", tc.a, tc.b)
		t.Run(testname, func(t *testing.T) {
			ans, err := addition.Add(tc.a, tc.b)
			if ans != tc.want {
				t.Errorf("got %d, want %d", ans, tc.want)
			}
			if tc.expectedErr != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("wrong error")
			}
		})

	}
}
