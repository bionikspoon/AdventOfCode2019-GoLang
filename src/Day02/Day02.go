package day02

import (
	"errors"
	"fmt"
	"strconv"
)

// Part1 finds the head value after running the program
func Part1(input string) string {
	programInput := programInput{12, 2}
	opcodes := deserialize(input)

	nextOpcodes, err := runWithInput(programInput, opcodes)
	check(err)

	return strconv.Itoa(nextOpcodes[0])
}

// Part2 performs a goal seek to find the correct nouns and verbs
func Part2(input string) string {
	opcodes := deserialize(input)

	solution, err := goalSeek(19690720, opcodes)
	check(err)

	return strconv.Itoa(100*solution.noun + solution.verb)
}

type programInput struct {
	noun, verb int
}

func goalSeek(target int, opcodes []int) (programInput, error) {
	ch := make(chan programInput)
	errCh := make(chan error)

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			candidate := programInput{noun, verb}

			go trial(ch, errCh, target, candidate, opcodes)
		}
	}

	select {
	case candidate := <-ch:
		return candidate, nil
	case err := <-errCh:
		return programInput{}, err
	default:
		return programInput{}, errors.New("goalSeek never finished")
	}
}

func trial(ch chan programInput, errCh chan error, target int, candidate programInput, opcodes []int) {
	nextOpcodes, err := runWithInput(candidate, dup(opcodes))
	if err != nil {
		errCh <- err
		return
	}
	if nextOpcodes[0] == target {
		ch <- candidate
	}
}

func runWithInput(input programInput, opcodes []int) ([]int, error) {
	opcodes[1] = input.noun
	opcodes[2] = input.verb

	return runProgram(opcodes)
}

func runProgram(opcodes []int) ([]int, error) {
	const maxIterations = 100

	instructionPointer := 0

	for j := 0; j <= maxIterations; j++ {
		opcode := opcodes[instructionPointer]

		switch opcode {
		case 1:
			opcodes[opcodes[instructionPointer+3]] = opcodes[opcodes[instructionPointer+1]] + opcodes[opcodes[instructionPointer+2]]
			instructionPointer += 4

		case 2:
			opcodes[opcodes[instructionPointer+3]] = opcodes[opcodes[instructionPointer+1]] * opcodes[opcodes[instructionPointer+2]]
			instructionPointer += 4

		case 99:
			return opcodes, nil

		default:
			return nil, fmt.Errorf("Unknown opcode %+v", opcode)
		}
	}

	return nil, errors.New("Program never finished")
}
