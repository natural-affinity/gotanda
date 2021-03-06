package gotanda

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"testing"

	capturer "github.com/kami-zh/go-capturer"
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
	str := capturer.CaptureStdout(p)
	return []byte(str), str
}
