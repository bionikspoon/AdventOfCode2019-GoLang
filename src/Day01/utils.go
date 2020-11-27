package day01

import (
	"strconv"
	"strings"
)

func withIO(input string, fn func(ints []int) int) string {
	ints := readInts(input)

	return strconv.Itoa(fn(ints))
}

func readInts(input string) (ints []int) {
	lines := strings.Split(input, "\n")

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
