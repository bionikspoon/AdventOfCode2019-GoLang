package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) string {
	fields := strings.Fields(input)

	distance, err := NewDiagram(fields)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(distance)
}

func Part2(input string) string {
	return input
}

type Direction int

const (
	U Direction = iota
	R
	D
	L
)

type Instruction struct {
	direction Direction
	distance  int
}

type Wire struct {
	index        int
	instructions []Instruction
}

type Coordinates struct {
	x, y int
}

func NewDiagram(serializedWires []string) (int, error) {
	wires := make([]*Wire, len(serializedWires))
	for i, serializedWire := range serializedWires {
		serializedInstructions := strings.Split(serializedWire, ",")

		instructions := make([]Instruction, len(serializedInstructions))

		for j, serializedInstruction := range serializedInstructions {
			instruction, err := deserializeInstruction(serializedInstruction)

			if err != nil {
				return 0, err
			}

			instructions[j] = instruction
		}

		wires[i] = &Wire{i, instructions}
	}

	grid := make(map[Coordinates][]*Wire)

	for _, wire := range wires {
		position := Coordinates{0, 0}

		for _, instruction := range wire.instructions {
			for i := 0; i < instruction.distance; i++ {
				switch instruction.direction {
				case U:
					position.y++
				case R:
					position.x++
				case D:
					position.y--
				case L:
					position.x--
				}

				_, ok := grid[position]

				if !ok {
					grid[position] = []*Wire{}
				}

				if positionContains(grid[position], wire) {
					continue
				}

				grid[position] = append(grid[position], wire)
			}
		}
	}

	intersections := []Coordinates{}

	for coordinates, wires := range grid {
		if len(wires) == 1 {
			continue
		}

		intersections = append(intersections, coordinates)
	}

	fmt.Println(intersections)

	intersectionDistances := make([]int, len(intersections))
	centralPort := Coordinates{0, 0}

	for i, intersection := range intersections {
		intersectionDistances[i] = manhattanDistance(intersection, centralPort)
	}

	fmt.Println(intersectionDistances)

	return minInt(intersectionDistances), nil
}

func manhattanDistance(a, b Coordinates) int {
	return intAbs(a.x-b.x) + intAbs(a.y-b.y)
}

func minInt(slice []int) int {
	var min int
	for index, element := range slice {
		if index == 0 || element < min {
			min = element
		}
	}
	return min
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func positionContains(gridPosition []*Wire, wire *Wire) bool {
	for _, candidate := range gridPosition {
		if candidate == wire {
			return true
		}
	}

	return false
}

func deserializeInstruction(serializedInstruction string) (Instruction, error) {
	re, err := regexp.Compile(`([URDL])(\d+)`)
	if err != nil {
		return Instruction{}, err
	}

	groups := re.FindStringSubmatch(serializedInstruction)
	direction, err := deserializeDirection(groups[1])
	if err != nil {
		return Instruction{}, err
	}
	distance, err := strconv.Atoi(groups[2])
	if err != nil {
		return Instruction{}, err

	}

	return Instruction{
		direction: direction,
		distance:  distance,
	}, nil
}

func deserializeDirection(serializedDirection string) (Direction, error) {
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
