package gotanda

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// Run command string
func Run(cmd string) ([]byte, error) {
	command := exec.Command("sh", "-c", cmd)
	out, err := command.CombinedOutput()

	return out, err
}

// LoadTestFile from testdata directory
func LoadTestFile(t *testing.T, dir string, name string) (string, []byte) {
	path := filepath.Join(dir, name)
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatal(err)
	}

	return path, bytes
}

// Capture Output
func Capture(p func()) ([]byte, string) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = stdout
	}()

	p()
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.Bytes(), buf.String()
}
