package day03

import (
	"strconv"
	"strings"
)

// Part1 finds the interesection with shortest manhatten distance
func Part1(input string) string {
	fields := strings.Fields(input)

	distance, err := shortestManhattenDistance(fields)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(distance)
}

// Part2 finds the interesection with lowest latency (distance)
func Part2(input string) string {
	fields := strings.Fields(input)

	distance, err := shortestLatency(fields)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(distance)
}

func shortestManhattenDistance(serializedWires []string) (int, error) {
	return withIntersections(serializedWires, func(intersections []Intersection) int {
		intersectionDistances := make([]int, len(intersections))
		centralPort := Coordinates{0, 0}

		for i, intersection := range intersections {
			intersectionDistances[i] = intersection.manhattanDistanceFrom(centralPort)
		}

		return minInt(intersectionDistances)
	})
}

func shortestLatency(serializedWires []string) (int, error) {
	return withIntersections(serializedWires, func(intersections []Intersection) int {
		intersectionDistances := make([]int, len(intersections))

		for i, intersection := range intersections {
			intersectionDistances[i] = intersection.totalSteps()
		}

		return minInt(intersectionDistances)
	})

}

func withIntersections(serializedWires []string, fn func(intersections []Intersection) int) (int, error) {
	wires, err := deserializeWires(serializedWires)

	if err != nil {
		return 0, err
	}

	intersections := NewGrid().AddWires(wires).Intersections()
	return fn(intersections), nil

}
