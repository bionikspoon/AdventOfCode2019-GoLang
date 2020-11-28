package day03

import (
	"bionikspoon/go-advent-of-code-2019/internal/testutils"
	"testing"
)

func TestPart1(t *testing.T) {
	want := "731"

	if got := Part1(testutils.ReadFile(t, "input.txt")); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := "5672"

	if got := Part2(testutils.ReadFile(t, "input.txt")); got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func Test_shortestManhattenDistance(t *testing.T) {
	type args struct {
		serializedWires []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"case", args{[]string{"R8,U5,L5,D3", "U7,R6,D4,L4"}}, 6, false},
		{"case", args{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}}, 159, false},
		{"case", args{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}}, 135, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shortestManhattenDistance(tt.args.serializedWires)
			if (err != nil) != tt.wantErr {
				t.Errorf("shortestManhattenDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("shortestManhattenDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortestLatency(t *testing.T) {
	type args struct {
		serializedWires []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"case", args{[]string{"R8,U5,L5,D3", "U7,R6,D4,L4"}}, 30, false},
		{"case", args{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}}, 610, false},
		{"case", args{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}}, 410, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shortestLatency(tt.args.serializedWires)
			if (err != nil) != tt.wantErr {
				t.Errorf("shortestLatency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("shortestLatency() = %v, want %v", got, tt.want)
			}
		})
	}
}
