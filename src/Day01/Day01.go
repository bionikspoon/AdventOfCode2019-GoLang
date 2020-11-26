package day01

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
