package dayeight

import (
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
	Z int
}

func ParseInput(path string) ([]Point, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []Point{}, err
	}

	result := []Point{}
	for line := range strings.SplitSeq(string(data), "\n") {
		if line == "" {
			continue
		}
		coordinates := []int{}
		for coord := range strings.SplitSeq(line, ",") {
			if coord == "" {
				continue
			}
			n, err := strconv.Atoi(coord)
			if err != nil {
				return []Point{}, err
			}
			coordinates = append(coordinates, n)
		}
		point := Point{
			X: coordinates[0],
			Y: coordinates[1],
			Z: coordinates[2],
		}
		result = append(result, point)
	}
	return result, nil
}

type Line struct {
	Dist   float64
	Point1 Point
	Point2 Point
}

// func getTop1000Lines(points []Point) []Line
//
// type Circuits map[int][]Point
//
// type PointMembers map[Point]int
//
// func joinPoints(line Line, circuits Circuits, pointMembers PointMembers)
//
// func getSum(circuits Circuits) int
//
// func PartOne(input []Point) int
