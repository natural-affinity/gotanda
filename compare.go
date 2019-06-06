package gotanda

import (
	"bytes"
	"io/ioutil"
	"testing"
)

// TestCase properties
type TestCase struct {
	Name string
}

// CompareFile (are the *.input and *.golden files the same?)
func CompareFile(t *testing.T, tc TestCase, update *bool) {
	_, command := LoadTestFile(t, "testdata", tc.Name+".input")
	golden, expected := LoadTestFile(t, "testdata", tc.Name+".golden")
	aout, _ := Run(string(command))

	if *update {
		ioutil.WriteFile(golden, aout, 0644)
	}

	expected, _ = ioutil.ReadFile(golden)
	out := !bytes.Equal(aout, expected)

	if out {
		t.Errorf("Test: %s\n Expected: %s\n Actual: %s\n", tc.Name, expected, aout)
	}
}

// CompareError (are the errors the same?)
func CompareError(actual error, expected error) bool {
	a := (actual != nil && expected != nil && actual.Error() == expected.Error())
	b := (actual == nil && expected == nil)

	return a || b
}
