package day01

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type module struct{ mass int }

func (m module) fuelRequired() int {
	return intDiv(m.mass, 3) - 2
}

func intDiv(numerator, denominator int) int {
	return int(math.Floor(float64(numerator) / float64(denominator)))
}

func Part1(input string) string {

	temp := strings.Split(input, "\n")

	sum := 0

	for _, line := range temp {
		lineValue, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		sum += module{lineValue}.fuelRequired()

	}

	fmt.Println(sum)

	return strconv.Itoa(sum)
}

func Part2(input string) string {
	return input
}
