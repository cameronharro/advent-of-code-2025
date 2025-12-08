package dayeight

import (
	"maps"
	"math"
	"os"
	"slices"
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

func getDistance(point1, point2 Point) float64 {
	return math.Sqrt(math.Pow(float64(point1.X-point2.X), 2) + math.Pow(float64(point1.Y-point2.Y), 2) + math.Pow(float64(point1.Z-point2.Z), 2))
}

func getShortestLines(points []Point) []Line {
	result := make([]Line, (len(points)-1)*len(points)/2)
	marker := 0
	for i, point1 := range points {
		for _, point2 := range points[i+1:] {
			result[marker] = Line{
				Point1: point1,
				Point2: point2,
				Dist:   getDistance(point1, point2),
			}
			marker++
		}
	}
	slices.SortFunc(result, func(a, b Line) int {
		diff := a.Dist - b.Dist
		if diff > 0 {
			return 1
		} else if diff < 0 {
			return -1
		}
		return 0
	})
	// fmt.Println("len() shortest lines:", len(result))
	// fmt.Println("shortest lines:", result)
	return result
}

type Circuits map[int][]Point

type PointMembers map[Point]int

func joinPoints(line Line, circuits Circuits, pointMembers PointMembers, nextCircuit int) (int, int) {
	currentPoint1Member, currentPoint2Member := pointMembers[line.Point1], pointMembers[line.Point2]
	newCircuitLength := 0
	if currentPoint1Member == 0 && currentPoint2Member == 0 {
		circuits[nextCircuit] = []Point{line.Point1, line.Point2}
		pointMembers[line.Point1] = nextCircuit
		pointMembers[line.Point2] = nextCircuit
		return nextCircuit + 1, 2
	} else if currentPoint1Member == currentPoint2Member {
		newCircuitLength = len(circuits[currentPoint1Member])
	} else if currentPoint1Member == 0 {
		pointMembers[line.Point1] = currentPoint2Member
		circuits[currentPoint2Member] = append(circuits[currentPoint2Member], line.Point1)
		newCircuitLength = len(circuits[currentPoint2Member])
	} else if currentPoint2Member == 0 {
		pointMembers[line.Point2] = currentPoint1Member
		circuits[currentPoint1Member] = append(circuits[currentPoint1Member], line.Point2)
		newCircuitLength = len(circuits[currentPoint1Member])
	} else {
		circuits[currentPoint1Member] = append(circuits[currentPoint1Member], circuits[currentPoint2Member]...)
		newCircuitLength = len(circuits[currentPoint1Member])
		for _, point := range circuits[currentPoint2Member] {
			pointMembers[point] = currentPoint1Member
		}
		delete(circuits, currentPoint2Member)
	}
	return nextCircuit, newCircuitLength
}

func getSum(circuits Circuits) int {
	circuitsSlice := slices.Collect(maps.Values(circuits))
	if len(circuitsSlice) < 3 {
		return 0
	}
	slices.SortFunc(circuitsSlice, func(E1, E2 []Point) int {
		return len(E2) - len(E1)
	})
	// fmt.Println("sorted circuits:", circuitsSlice)
	return len(circuitsSlice[0]) * len(circuitsSlice[1]) * len(circuitsSlice[2])
}

func PartOne(input []Point, numConnections int) int {
	shortestLines := getShortestLines(input)

	circuits := Circuits{}
	PointMembers := PointMembers{}
	nextCircuit := 0
	for _, line := range shortestLines[:min(len(shortestLines), numConnections)] {
		nextCircuit, _ = joinPoints(line, circuits, PointMembers, nextCircuit)
		// fmt.Println("circuits:", circuits)
	}

	return getSum(circuits)
}

func PartTwo(input []Point) int {
	shortestLines := getShortestLines(input)
	targetLength := len(input)
	circuits := Circuits{}
	PointMembers := PointMembers{}
	nextCircuit := 0
	for _, line := range shortestLines {
		innerNextCircuit, newCircuitLength := joinPoints(line, circuits, PointMembers, nextCircuit)
		nextCircuit = innerNextCircuit
		// fmt.Printf("Connection: %v\n", line)
		// fmt.Printf("New Circuit Length: %d, target length: %d\n", newCircuitLength, targetLength)
		// fmt.Println()
		if newCircuitLength >= targetLength {
			return line.Point1.X * line.Point2.X
		}
	}
	return 0
}
