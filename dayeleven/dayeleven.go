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

func traverseGraph(graph Graph, order []string, start string) map[string]int {
	pathsToNode := map[string]int{}
	startIndex := slices.Index(order, start)
	pathsToNode[start] = 1
	for _, node := range order[startIndex:] {
		for _, neighbor := range graph[node] {
			pathsToNode[neighbor] += pathsToNode[node]
		}
	}
	return pathsToNode
}

func PartOne(graph Graph) int {
	order := TopologicallySortNodes(graph)
	paths := traverseGraph(graph, order, "you")
	return paths["out"]
}

func PartTwo(graph Graph) int {
	paths := traverseGraph(graph, TopologicallySortNodes(graph), "svr")
	return paths["out"]
}
