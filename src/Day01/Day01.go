package day01

import (
	"strconv"
	"strings"
)

// Part1 solves for the fuel requirement for modules
func Part1(input string) string {
	return withIO(input, sumWith(fuelRequiredForMass))
}

// Part2 solves for the fuel requirement for additional fuel
func Part2(input string) string {

	return withIO(input, sumWith(fuelRequiredForFuel))
}

func sumWith(fn func(int) int) func([]int) int {

	return func(masses []int) int {
		sum := 0
		for _, mass := range masses {
			sum += fn(mass)
		}

		return sum
	}
}

func fuelRequiredForMass(mass int) int {
	return int(mass/3) - 2
}

func fuelRequiredForFuel(mass int) int {
	fuelRequired := fuelRequiredForMass(mass)

	if fuelRequired >= 0 {
		return fuelRequired + fuelRequiredForFuel(fuelRequired)
	}

	return 0
}

func withIO(input string, fn func(ints []int) int) string {
	ints := readInts(input)

	return strconv.Itoa(fn(ints))
}

func readInts(input string) (ints []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		i, err := strconv.Atoi(line)
		check(err)
		ints = append(ints, i)
	}

	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
