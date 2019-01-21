package gotanda_test

import (
	"errors"
	"testing"

	"github.com/natural-affinity/gotanda"
)

func TestCompareError(t *testing.T) {
	cases := []struct {
		Name     string
		A        error
		B        error
		Expected bool
	}{
		{"nil", nil, nil, true},
		{"a.nil", nil, errors.New("B"), false},
		{"b.nil", errors.New("A"), nil, false},
		{"same", errors.New("e"), errors.New("e"), true},
		{"diff", errors.New("A"), errors.New("B"), false},
	}

	for _, tc := range cases {
		actual := gotanda.CompareError(tc.A, tc.B)
		if actual != tc.Expected {
			t.Errorf("Test: %s\nActual: %t\nExpected: %t\n",
				tc.Name, actual, tc.Expected)
		}
	}
}
