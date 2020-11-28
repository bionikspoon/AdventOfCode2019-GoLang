package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type direction int

const (
	// U is Up
	U direction = iota
	// R is Right
	R
	// D is Down
	D
	// L is Left
	L
)

type instruction struct {
	direction direction
	distance  int
}

// Wire is a set of instructions
type Wire struct {
	ID           int
	instructions []instruction
}

func deserializeWires(serializedWires []string) ([]*Wire, error) {
	wires := make([]*Wire, len(serializedWires))
	for i, serializedWire := range serializedWires {
		serializedInstructions := strings.Split(serializedWire, ",")

		instructions := make([]instruction, len(serializedInstructions))

		for j, serializedInstruction := range serializedInstructions {
			instruction, err := deserializeInstruction(serializedInstruction)

			if err != nil {
				return nil, err
			}

			instructions[j] = instruction
		}

		wires[i] = &Wire{i, instructions}
	}
	return wires, nil
}

func deserializeInstruction(serializedInstruction string) (instruction, error) {
	re, err := regexp.Compile(`([URDL])(\d+)`)
	if err != nil {
		return instruction{}, err
	}

	groups := re.FindStringSubmatch(serializedInstruction)
	direction, err := deserializeDirection(groups[1])
	if err != nil {
		return instruction{}, err
	}
	distance, err := strconv.Atoi(groups[2])
	if err != nil {
		return instruction{}, err

	}

	return instruction{
		direction: direction,
		distance:  distance,
	}, nil
}

func deserializeDirection(serializedDirection string) (direction, error) {
	switch serializedDirection {
	case "U":
		return U, nil
	case "R":
		return R, nil
	case "D":
		return D, nil
	case "L":
		return L, nil
	default:
		return 0, fmt.Errorf("Unknown direction %v", serializedDirection)

	}
}
