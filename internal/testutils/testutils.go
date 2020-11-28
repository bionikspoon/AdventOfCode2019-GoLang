package testutils

import (
	"io/ioutil"
	"strings"
	"testing"
)

// ReadFile attempts to read a file or fails the test
func ReadFile(t *testing.T, filename string) string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(string(input))
}
