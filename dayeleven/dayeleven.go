package dayeleven

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

type Graph map[string][]string

func ParseInput(path string) (Graph, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	result := make(Graph)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		re := regexp.MustCompile(`(\w*):([\w\s]*)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			return nil, fmt.Errorf("Invalid line: %s\n", line)
		}
		result[matches[1]] = strings.Fields(matches[2])
	}
	return result, nil
}

func TopologicallySortNodes(graph Graph) []string {
	queue := []string{}
	indegrees := map[string]int{}
	for thisNode, neighbors := range graph {
		if _, exists := indegrees[thisNode]; !exists {
			indegrees[thisNode] = 0
		}
		for _, neighbor := range neighbors {
			if _, exists := indegrees[neighbor]; !exists {
				indegrees[neighbor] = 0
			}
			indegrees[neighbor] += 1
		}
	}

	for node, indegree := range indegrees {
		if indegree == 0 {
			queue = append(queue, node)
			delete(indegrees, node)
		}
	}

	result := []string{}
	for i := 0; i < len(queue); i++ {
		currentNode := queue[i]
		for _, neighbor := range graph[currentNode] {
			if _, exists := indegrees[neighbor]; !exists {
				result = append(result, neighbor)
			}
			indegrees[neighbor] -= 1
			if indegrees[neighbor] == 0 {
				queue = append(queue, neighbor)
				delete(indegrees, neighbor)
			}
		}
		result = append(result, currentNode)
	}
	return result
}

type MaskCounts [4]int

const (
	MaskNone = 0
	MaskFFT  = 1 << 0
	MaskDAC  = 1 << 1
	MaskBoth = MaskFFT | MaskDAC
)

func traverseGraph(graph Graph, order []string, start string) map[string]MaskCounts {
	pathsToNode := make(map[string]MaskCounts)
	startIndex := slices.Index(order, start)
	pathsToNode[start] = MaskCounts{1, 0, 0, 0}
	for _, node := range order[startIndex:] {
		transitionMask := MaskNone
		if node == "fft" {
			transitionMask |= MaskFFT
		}
		if node == "dac" {
			transitionMask |= MaskDAC
		}

		counts := pathsToNode[node]
		if transitionMask != MaskNone {
			newCounts := [4]int{}
			for mask := range 4 {
				newMask := mask | transitionMask
				newCounts[newMask] += counts[mask]
			}
			counts = newCounts
			pathsToNode[node] = counts
		}

		for _, neighbor := range graph[node] {
			neighborCounts := pathsToNode[neighbor]
			for i := range 4 {
				neighborCounts[i] += counts[i]
			}
			pathsToNode[neighbor] = neighborCounts
		}
	}
	return pathsToNode
}

func PartOne(graph Graph) int {
	paths := traverseGraph(graph, TopologicallySortNodes(graph), "you")
	result := 0
	for _, v := range paths["out"] {
		result += v
	}
	return result
}

func PartTwo(graph Graph) int {
	paths := traverseGraph(graph, TopologicallySortNodes(graph), "svr")
	return paths["out"][3]
}
