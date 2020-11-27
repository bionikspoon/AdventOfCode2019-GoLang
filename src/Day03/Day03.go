package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) string {
	fields := strings.Fields(input)

	distance, err := shortestManhattenDistance(fields)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(distance)
}

func Part2(input string) string {
	fields := strings.Fields(input)

	distance, err := shortestLatency(fields)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(distance)
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

type GridWire struct {
	steps int
	wire  *Wire
}

type Grid map[Coordinates][]GridWire

func NewGrid(serializedWires []string) (*Grid, error) {
	grid := make(Grid)
	wires := make([]*Wire, len(serializedWires))
	for i, serializedWire := range serializedWires {
		serializedInstructions := strings.Split(serializedWire, ",")

		instructions := make([]Instruction, len(serializedInstructions))

		for j, serializedInstruction := range serializedInstructions {
			instruction, err := deserializeInstruction(serializedInstruction)

			if err != nil {
				return &grid, err
			}

			instructions[j] = instruction
		}

		wires[i] = &Wire{i, instructions}
	}

	for _, wire := range wires {
		steps := 0
		position := Coordinates{0, 0}

		for _, instruction := range wire.instructions {
			for i := 0; i < instruction.distance; i++ {
				steps++

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
					grid[position] = []GridWire{}
				}

				if positionContains(grid[position], wire) {
					continue
				}

				grid[position] = append(grid[position], GridWire{steps, wire})
			}
		}
	}

	return &grid, nil
}

type Intersection struct {
	coordinates Coordinates
	gridWires   []GridWire
}

func (grid *Grid) Intersections() []Intersection {
	intersections := []Intersection{}

	for coordinates, wires := range *grid {
		if len(wires) == 1 {
			continue
		}

		intersections = append(intersections, Intersection{coordinates, wires})
	}

	return intersections
}

func shortestManhattenDistance(serializedWires []string) (int, error) {
	grid, err := NewGrid(serializedWires)
	if err != nil {
		return 0, err
	}
	intersections := grid.Intersections()

	intersectionDistances := make([]int, len(intersections))
	centralPort := Coordinates{0, 0}

	for i, intersection := range intersections {
		intersectionDistances[i] = manhattanDistance(intersection.coordinates, centralPort)
	}

	return minInt(intersectionDistances), nil
}

func shortestLatency(serializedWires []string) (int, error) {
	grid, err := NewGrid(serializedWires)
	if err != nil {
		return 0, err
	}
	intersections := grid.Intersections()

	intersectionDistances := make([]int, len(intersections))

	for i, intersection := range intersections {
		intersectionDistances[i] = intersectionSteps(intersection)
	}

	return minInt(intersectionDistances), nil
}

func intersectionSteps(intersection Intersection) int {
	steps := 0

	for _, gridWire := range intersection.gridWires {
		steps += gridWire.steps
	}

	return steps
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

func positionContains(gridPosition []GridWire, wire *Wire) bool {
	for _, candidate := range gridPosition {
		if candidate.wire == wire {
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
