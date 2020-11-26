package day01

import (
	"strconv"
	"strings"
)

// Part1 solves for the fuel requirement for modules
func Part1(input string) string {

	moduleMasses := readInts(input)
	fuelRequired := sumWith(fuelRequiredForMass, moduleMasses)

	return strconv.Itoa(fuelRequired)
}

// Part2 solves for the fuel requirement for additional fuel
func Part2(input string) string {
	moduleMasses := readInts(input)
	fuelRequired := sumWith(fuelRequiredForFuel, moduleMasses)

	return strconv.Itoa(fuelRequired)
}

func sumWith(fn func(int) int, masses []int) int {
	sum := 0
	for _, mass := range masses {
		sum += fn(mass)
	}

	return sum
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
