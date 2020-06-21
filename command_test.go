package gotanda_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/natural-affinity/gotanda"
)

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

func TestCapture(t *testing.T) {
	cases := []struct {
		Name     string
		Text     string
		Expected []byte
	}{
		{"empty", "", []byte("")},
		{"text", "hello", []byte("hello")},
	}

	for _, tc := range cases {
		buf := gotanda.Capture(func() {
			fmt.Print(tc.Text)
		})

		byt := !bytes.Equal(buf.Bytes(), tc.Expected)
		str := !(buf.String() == tc.Text)

		if byt || str {
			t.Errorf("Test: %s\nActual: %s\nExpected: %s\n",
				tc.Name, buf.Bytes(), tc.Expected)
		}
	}
}
