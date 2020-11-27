package day02

import (
	"errors"
	"fmt"
	"strconv"
)

type programKnobs struct {
	noun, verb int
}
type memory []int

type trialResults struct {
	candidate programKnobs
	memory    memory
	err       error
}

// Part1 finds the head value after running the program
func Part1(input string) string {
	programInput := programKnobs{12, 2}
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

func goalSeek(target int, memory memory) (programKnobs, error) {
	ch := make(chan trialResults)

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			candidate := programKnobs{noun, verb}

			go trial(ch, candidate, memory)
		}
	}

	for i := 0; i < 100*100; i++ {
		trialResults := <-ch

		if trialResults.err != nil {
			return trialResults.candidate, trialResults.err
		}

		if trialResults.memory[0] == target {
			return trialResults.candidate, trialResults.err
		}
	}

	return programKnobs{}, errors.New("goalSeek never finished")
}

func trial(ch chan trialResults, candidate programKnobs, memory memory) {
	nextMemory, err := runWithInput(candidate, dup(memory))

	ch <- trialResults{candidate, nextMemory, err}
}

func runWithInput(input programKnobs, memory memory) (memory, error) {
	memory[1] = input.noun
	memory[2] = input.verb

	return runProgram(memory)
}

func runProgram(memory memory) (memory, error) {
	const maxIterations = 100

	instructionPointer := 0

	for j := 0; j <= maxIterations; j++ {
		opcode := memory[instructionPointer]

		switch opcode {
		case 1:
			memory[memory[instructionPointer+3]] = memory[memory[instructionPointer+1]] + memory[memory[instructionPointer+2]]
			instructionPointer += 4

		case 2:
			memory[memory[instructionPointer+3]] = memory[memory[instructionPointer+1]] * memory[memory[instructionPointer+2]]
			instructionPointer += 4

		case 99:
			return memory, nil

		default:
			return nil, fmt.Errorf("Unknown opcode %+v", opcode)
		}
	}

	return nil, errors.New("Program never finished")
}
