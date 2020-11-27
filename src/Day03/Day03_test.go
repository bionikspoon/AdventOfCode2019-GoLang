package day03

import (
	"testing"

	"bionikspoon/go-advent-of-code-2019/lib/testutils"
)

func TestPart1(t *testing.T) {
	t.Skip("skipping")
	want := "hello santa"

	if got := Part1(testutils.ReadFile(t, "input.txt")); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Skip("skipping")
	want := "hello santa"

	if got := Part2(testutils.ReadFile(t, "input.txt")); got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}
