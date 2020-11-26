package dayxx

import (
	"io/ioutil"
	"testing"
)

func TestPart1(t *testing.T) {
	want := "hello santa"

	if got := Part1(readFile(t, "input.txt")); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := "hello santa"

	if got := Part2(readFile(t, "input.txt")); got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func readFile(t *testing.T, filename string) string {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return string(input)
}
