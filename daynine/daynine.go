package daynine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func ParseInput(path string) ([]Point, error) {
	file, err := os.Open(path)
	if err != nil {
		return []Point{}, err
	}

	scanner := bufio.NewScanner(file)
	result := []Point{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		nStrings := strings.Split(line, ",")
		if len(nStrings) != 2 {
			return []Point{}, fmt.Errorf("Invalid point definition: %s", line)
		}
		x, err := strconv.Atoi(nStrings[0])
		if err != nil {
			return []Point{}, fmt.Errorf("Invalid point definition: %s", line)
		}
		y, err := strconv.Atoi(nStrings[1])
		if err != nil {
			return []Point{}, fmt.Errorf("Invalid point definition: %s", line)
		}
		result = append(result, Point{X: x, Y: y})
	}
	return result, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getArea(p1, p2 Point) int {
	length, width := 1+abs(p1.X-p2.X), 1+abs(p1.Y-p2.Y)
	return length * width
}

func PartOne(points []Point) int {
	result := 0
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			area := getArea(p1, p2)
			if area > result {
				result = area
			}
		}
	}
	return result
}

type Box struct {
	P1, P2 Point
}

func boxesIntersect(box1, box2 Box) bool {
	xMin1 := min(box1.P1.X, box1.P2.X)
	xMax1 := max(box1.P1.X, box1.P2.X)
	yMin1 := min(box1.P1.Y, box1.P2.Y)
	yMax1 := max(box1.P1.Y, box1.P2.Y)
	xMin2 := min(box2.P1.X, box2.P2.X)
	xMax2 := max(box2.P1.X, box2.P2.X)
	yMin2 := min(box2.P1.Y, box2.P2.Y)
	yMax2 := max(box2.P1.Y, box2.P2.Y)
	return (xMin1 <= xMax2 && xMin2 <= xMax1) && (yMin1 <= yMax2 && yMin2 <= yMax1)
}

type Boxes map[Box]int
type Edges map[Point]struct{}

func createBoxes(points []Point) (Boxes, Edges) {
	boxes := make(Boxes)
	edges := make(Edges)
	for i, p1 := range points {
		// Make boxes
		for _, p2 := range points[i+1:] {
			area := getArea(p1, p2)
			boxes[Box{P1: p1, P2: p2}] = area
		}
		//Make Edges
		var p2 Point
		if i == len(points)-1 {
			p2 = points[0]
		} else {
			p2 = points[i+1]
		}
		for y := min(p1.Y, p2.Y); y <= max(p1.Y, p2.Y); y++ {
			edges[Point{X: p1.X, Y: y}] = struct{}{}
		}
		for x := min(p1.X, p2.X); x <= max(p1.X, p2.X); x++ {
			edges[Point{X: x, Y: p1.Y}] = struct{}{}
		}
	}
	return boxes, edges
}

type Vector struct {
	P         Point
	Direction int
}

func walkEdge(vector Vector, edges Edges) Vector {
	nextVector := Vector{}

	nextPoints := []Point{
		{X: vector.P.X + 1, Y: vector.P.Y},
		{X: vector.P.X, Y: vector.P.Y - 1},
		{X: vector.P.X - 1, Y: vector.P.Y},
		{X: vector.P.X, Y: vector.P.Y + 1},
	}

	for i := vector.Direction - 1; i < vector.Direction+3; i++ {
		direction := i
		if direction < 0 {
			direction += 4
		}
		direction = direction % 4
		pToCheck := nextPoints[direction]
		_, exists := edges[pToCheck]
		if exists {
			continue
		}
		nextVector.P = pToCheck
		nextVector.Direction = direction
		break
	}

	return nextVector
}

func PartTwo(points []Point, initialVectorToWalk Vector) int {
	boxes, edges := createBoxes(points)
	vector := initialVectorToWalk
	lastCorner := vector
	for {
		nextVector := walkEdge(vector, edges)
		if nextVector.Direction != vector.Direction {
			lineToCheck := Box{P1: vector.P, P2: lastCorner.P}
			for k := range boxes {
				if boxesIntersect(k, lineToCheck) {
					delete(boxes, k)
				}
			}
			lastCorner = nextVector
		}
		if nextVector == initialVectorToWalk {
			break
		}
		vector = nextVector
	}

	result := 0
	for _, v := range boxes {
		if v > result {
			result = v
		}
	}
	return result
}
