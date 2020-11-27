package day02

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) string {
	opcodes := deserialize(input)

	opcodes[1] = 12
	opcodes[2] = 2

	nextOpcodes, err := runProgram(opcodes)
	check(err)

	return strconv.Itoa(nextOpcodes[0])
}

func Part2(input string) string {

	return input
}

func run(input string) string {
	opcodes := deserialize(input)

	nextOpcodes, err := runProgram(opcodes)
	check(err)

	return serialize(nextOpcodes)
}

func runProgram(opcodes []int) ([]int, error) {
	const maxIterations = 100

	i := 0

	for j := 0; j <= maxIterations; j++ {
		opcode := opcodes[i]

		switch opcode {
		case 1:
			opcodes[opcodes[i+3]] = opcodes[opcodes[i+1]] + opcodes[opcodes[i+2]]
			i += 4

		case 2:
			opcodes[opcodes[i+3]] = opcodes[opcodes[i+1]] * opcodes[opcodes[i+2]]
			i += 4

		case 99:
			return opcodes, nil

		default:
			return nil, fmt.Errorf("Unknown opcode %+v", opcode)
		}
	}

	return nil, errors.New("Program never finished")
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
