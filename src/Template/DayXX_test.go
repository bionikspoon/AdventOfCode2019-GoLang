package dayxx

import (
	"testing"

	"lib/testutils"
)

func TestPart1(t *testing.T) {
	want := "hello santa"

	if got := Part1(testutils.ReadFile(t, "input.txt")); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := "hello santa"

	if got := Part2(testutils.ReadFile(t, "input.txt")); got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}
