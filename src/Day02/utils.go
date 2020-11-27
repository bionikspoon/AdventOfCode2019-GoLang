package day02

import (
	"strconv"
	"strings"
)

func dup(ints []int) []int {
	next := make([]int, len(ints))
	copy(next, ints)
	return next
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func deserialize(input string) []int {
	opcodes := []int{}

	for _, s := range strings.Split(input, ",") {
		opcode, err := strconv.Atoi(s)
		check(err)

		opcodes = append(opcodes, opcode)
	}

	return opcodes
}

func serialize(opcodes []int) string {
	codes := make([]string, len(opcodes))

	for index, opcode := range opcodes {
		codes[index] = strconv.Itoa(opcode)
	}

	return strings.Join(codes, ",")
}
