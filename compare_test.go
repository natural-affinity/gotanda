package gotanda_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/natural-affinity/gotanda"
)

func TestCompareCommand(t *testing.T) {
	cases := []struct {
		Name         string
		Update       bool
		GoldenResult *gotanda.GoldenResult
	}{
		{
			"compare.command.same", false,
			&gotanda.GoldenResult{
				Match:    true,
				Updated:  false,
				Command:  []byte("echo \"Same\"\n"),
				Actual:   []byte("Same\n"),
				Expected: []byte("Same\n")},
		},
		{
			"compare.command.different", false,
			&gotanda.GoldenResult{
				Match:    false,
				Updated:  false,
				Command:  []byte("echo \"A\"\n"),
				Actual:   []byte("A\n"),
				Expected: []byte("B\n")},
		},
		{
			"compare.command.updated", true,
			&gotanda.GoldenResult{
				Match:    true,
				Updated:  true,
				Command:  []byte("echo \"updated\"\n"),
				Actual:   []byte("updated\n"),
				Expected: []byte("updated\n")},
		},
	}

	for _, tc := range cases {
		ra := gotanda.CompareCommand(t, gotanda.TestCase{tc.Name}, &tc.Update)
		gr := tc.GoldenResult

		out := !(ra.Match == gr.Match &&
			ra.Updated == gr.Updated &&
			bytes.Equal(ra.Actual, gr.Actual) &&
			bytes.Equal(ra.Expected, gr.Expected))

		if out {
			t.Errorf("Test: %s\nActual: %v\nExpected: %v\n", tc.Name, ra, gr)
		}
	}
}

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
			t.Errorf("Test: %s\nActual: %t\nExpected: %t\n", tc.Name, actual, tc.Expected)
		}
	}
}
