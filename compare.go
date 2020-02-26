package gotanda

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

// Assertable result
type Assertable interface {
	Assert()
}

// TestCase properties
type TestCase struct {
	Name string
}

// GoldenResult for command
type GoldenResult struct {
	Match    bool
	Updated  bool
	Command  []byte
	Actual   []byte
	Expected []byte
}

// Assert test result
func (gr *GoldenResult) Assert(t *testing.T, tc TestCase) {
	if !gr.Match {
		t.Errorf("Test: %s\n Expected: %s\n Actual: %s\n", tc.Name, gr.Expected, gr.Actual)
	}
}

// CompareCommand (does .input command result in .golden output?)
func CompareCommand(t *testing.T, tc TestCase, update *bool) *GoldenResult {
	r := &GoldenResult{}
	golden := filepath.Join("testdata", tc.Name+".golden")

	_, r.Command = LoadTestFile(t, "testdata", tc.Name+".input")
	r.Actual, _ = Run(string(r.Command))

	if *update {
		err := ioutil.WriteFile(golden, r.Actual, 0644)
		r.Updated = (err == nil)
	}

	r.Expected, _ = ioutil.ReadFile(golden)
	r.Match = bytes.Equal(r.Actual, r.Expected)

	return r
}

// CompareError (are the errors the same?)
func CompareError(actual error, expected error) bool {
	a := (actual != nil && expected != nil && actual.Error() == expected.Error())
	b := (actual == nil && expected == nil)

	return a || b
}
