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

func walkPaths(graph Graph, currentNode string, currentPath []string, result *[][]string) {
	if currentNode == "out" {
		temp := make([]string, len(currentPath))
		copy(temp, currentPath)
		*result = append(*result, temp)
	}

	for _, nextNode := range graph[currentNode] {
		if slices.Contains(currentPath, nextNode) {
			continue
		}
		currentPath = append(currentPath, nextNode)
		walkPaths(graph, nextNode, currentPath, result)
		currentPath = currentPath[:len(currentPath)-1]
	}
}

func PartOne(graph Graph) int {
	result := [][]string{}
	walkPaths(graph, "you", []string{"you"}, &result)
	return len(result)
}

func walkPathsTwo(graph Graph, currentNode string, currentPath []string, result *[][]string) {
	if currentNode == "out" && slices.Contains(currentPath, "fft") && slices.Contains(currentPath, "dac") {
		fmt.Println("Valid Path:", currentPath)
		temp := make([]string, len(currentPath))
		copy(temp, currentPath)
		*result = append(*result, temp)
	}

	for _, nextNode := range graph[currentNode] {
		if slices.Contains(currentPath, nextNode) {
			continue
		}
		currentPath = append(currentPath, nextNode)
		walkPathsTwo(graph, nextNode, currentPath, result)
		currentPath = currentPath[:len(currentPath)-1]
	}
}

func PartTwo(graph Graph) int {
	result := [][]string{}
	walkPathsTwo(graph, "svr", []string{"svr"}, &result)
	return len(result)
}
