package day02

import (
	"lib/testutils"
	"testing"
)

func TestPart1(t *testing.T) {
	want := "5305097"

	if got := Part1(testutils.ReadFile(t, "input.txt")); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := "4925"

	if got := Part2(testutils.ReadFile(t, "input.txt")); got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func Test_run(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"addition", args{"1,0,0,0,99"}, "2,0,0,0,99"},
		{"multiplication", args{"2,3,0,3,99"}, "2,3,0,6,99"},
		{"edge case", args{"2,4,4,5,99,0"}, "2,4,4,5,99,9801"},
		{"chaining opcodes", args{"1,1,1,4,99,5,6,0,99"}, "30,1,1,4,2,5,6,0,99"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.input); got != tt.want {
				t.Errorf("runProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func run(input string) string {
	opcodes := deserialize(input)

	nextOpcodes, err := runProgram(opcodes)
	check(err)

	return serialize(nextOpcodes)
}
