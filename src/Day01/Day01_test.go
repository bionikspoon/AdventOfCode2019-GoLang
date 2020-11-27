package day01

import (
	"bionikspoon/go-advent-of-code-2019/lib/testutils"
	"testing"
)

func TestPart1(t *testing.T) {
	want := "3497399"

	if got := Part1(testutils.ReadFile(t, "input.txt")); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := "5243207"

	if got := Part2(testutils.ReadFile(t, "input.txt")); got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}

}

func Test_fuelRequiredForMass(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case", args{12}, 2},
		{"case", args{14}, 2},
		{"case", args{1969}, 654},
		{"case", args{100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fuelRequiredForMass(tt.args.mass); got != tt.want {
				t.Errorf("fuelRequiredForMass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fuelRequiredForFuel(t *testing.T) {
	type args struct {
		fuelMass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case", args{14}, 2},
		{"case", args{1969}, 966},
		{"case", args{100756}, 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fuelRequiredForFuel(tt.args.fuelMass); got != tt.want {
				t.Errorf("fuelRequiredForFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
