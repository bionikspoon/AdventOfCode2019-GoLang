package day03

// Coordinates describes a location within a Grid
type Coordinates struct {
	x, y int
}

// GridWire describes a Wire's relationship with a Grid
type GridWire struct {
	steps int
	wire  *Wire
}

// An Intersection includes all of the information about a grid cell
type Intersection struct {
	coordinates Coordinates
	gridWires   GridWires
}

// GridWires are a collection of GridWire
type GridWires []GridWire

// A Grid creates a 2d representation of all the wires.
type Grid map[Coordinates]GridWires

// NewGrid creates new Grid
func NewGrid() *Grid {
	grid := make(Grid)

	return &grid
}

// AddWires walks through an array of Wires
func (grid *Grid) AddWires(wires []*Wire) *Grid {

	for _, wire := range wires {
		grid.addWire(wire)
	}

	return grid

}

func (grid Grid) addWire(wire *Wire) {
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
				grid[position] = GridWires{}
			}

			if grid[position].containsWire(wire) {
				continue
			}

			grid[position] = append(grid[position], GridWire{steps, wire})
		}
	}
}

// Intersections are all of the Grid cells with multiple wires.
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

func (gridWires GridWires) containsWire(wire *Wire) bool {
	for _, candidate := range gridWires {
		if candidate.wire == wire {
			return true
		}
	}

	return false
}

func (intersection *Intersection) manhattanDistanceFrom(start Coordinates) int {
	return manhattanDistance(intersection.coordinates, start)
}

func (intersection Intersection) totalSteps() int {
	steps := 0

	for _, gridWire := range intersection.gridWires {
		steps += gridWire.steps
	}

	return steps
}

func manhattanDistance(a, b Coordinates) int {
	return intAbs(a.x-b.x) + intAbs(a.y-b.y)
}
