package gotanda_test

import (
	"errors"
	"fmt"
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

func TestRun(t *testing.T) {
	cases := []struct {
		Name    string
		Command string
		Output  string
		Error   error
	}{
		{"stdout", "echo hello", "hello\n", nil},
		{"stderr", "fake-exe", "sh: fake-exe: command not found\n", errors.New("exit status 127")},
	}

	for _, tc := range cases {
		obyte, e := gotanda.Run(tc.Command)
		ostr := string(obyte)

		if ostr != tc.Output {
			t.Errorf("Test: %s\nActual: %s\nExpected: %s\n",
				tc.Name, ostr, tc.Output)
		}

		if !gotanda.CompareError(e, tc.Error) {
			t.Errorf("Test: %s\nActual: %s\nExpected: %s\n",
				tc.Name, e, tc.Error)
		}
	}
}

func TestLoadTestFile(t *testing.T) {
	cases := []struct {
		Name     string
		Expected string
	}{
		{"load.file.test", "example load file\n"},
	}

	for _, tc := range cases {
		dir := "testdata"
		fp, data := gotanda.LoadTestFile(t, dir, tc.Name+".input")

		expectedPath := fmt.Sprintf("%s\\%s.input", dir, tc.Name)
		if fp != expectedPath {
			t.Errorf("Test: %s\nActual: %s\nExpected: %s\n",
				tc.Name, fp, expectedPath)
		}

		datastr := string(data)
		if datastr != tc.Expected {
			t.Errorf("Test: %s\nActual: %s\nExpected: %s\n",
				tc.Name, datastr, tc.Expected)
		}
	}
}
