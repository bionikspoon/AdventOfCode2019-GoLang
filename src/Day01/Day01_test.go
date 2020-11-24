package day01

import (
	"io/ioutil"
	"testing"
)

func TestPart1(t *testing.T) {

	want := "3497399"

	if got := Part1(readFile(t, "input.txt")); got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.SkipNow()

	want := "5243207"

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

func Test_module_fuelRequired(t *testing.T) {
	type fields struct {
		mass int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"case", fields{12}, 2},
		{"case", fields{14}, 2},
		{"case", fields{1969}, 654},
		{"case", fields{100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := module{
				mass: tt.fields.mass,
			}
			if got := m.fuelRequired(); got != tt.want {
				t.Errorf("module.fuelRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
