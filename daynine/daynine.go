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
	// leftmost := 100000
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
		// if x < leftmost {
		// 	leftmost = x
		// }
		y, err := strconv.Atoi(nStrings[1])
		if err != nil {
			return []Point{}, fmt.Errorf("Invalid point definition: %s", line)
		}
		result = append(result, Point{X: x, Y: y})
	}
	// fmt.Println("leftmost:", leftmost)
	return result, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getArea(p1, p2 Point) int {
	length, width := abs(1+p1.X-p2.X), abs(1+p1.Y-p2.Y)
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

type Boxes map[Box]int
type Edges map[Point]struct{}

func pIsInBox(box Box, p Point) bool {
	top, bottom, left, right := max(box.P1.Y, box.P2.Y), min(box.P1.Y, box.P2.Y), min(box.P1.X, box.P2.X), max(box.P1.X, box.P2.X)
	return p.Y <= top && p.Y >= bottom && p.X >= left && p.X <= right
}

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
		if p1.X == p2.X {
			for y := min(p1.Y, p2.Y); y <= max(p1.Y, p2.Y); y++ {
				edges[Point{X: p1.X, Y: y}] = struct{}{}
			}
		}
		for x := min(p1.X, p2.X); x <= max(p1.X, p2.X); x++ {
			edges[Point{X: x, Y: p1.Y}] = struct{}{}
		}
	}
	// fmt.Println("Boxes:", boxes)
	// fmt.Println("Edges:", edges)
	return boxes, edges
}

type Vector struct {
	P         Point
	Direction int
}

func walkEdge(vector Vector, boxes Boxes, edges Edges) Vector {
	nextVector := Vector{}
	for k := range boxes {
		if pIsInBox(k, vector.P) {
			delete(boxes, k)
		}
	}

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
		pToCheck := nextPoints[direction%4]
		_, exists := edges[pToCheck]
		// fmt.Println("Vector:", vector)
		// fmt.Println("Direction to check:", direction)
		// fmt.Println("Point to Check:", pToCheck)
		// fmt.Println("Exists:", exists)
		// fmt.Println()
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
	for _ = range 100000 {
		vector = walkEdge(vector, boxes, edges)
		// fmt.Println("Vector:", vector)
		if vector == initialVectorToWalk {
			fmt.Println("Reached original starting point:", vector)
			break
		}
	}

	result := 0
	for _, v := range boxes {
		if v > result {
			result = v
		}
	}
	return result
}
