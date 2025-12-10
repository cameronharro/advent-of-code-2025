package dayten

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// true=on, false=off
type LightSet []bool
type ButtonSet []int
type JoltageSpecs []int
type Machine struct {
	Lights  LightSet
	Buttons []ButtonSet
	Joltage JoltageSpecs
}

func Map[T, V any](slice []T, mapFunc func(E T) V) []V {
	result := make([]V, len(slice))
	for i, element := range slice {
		result[i] = mapFunc(element)
	}
	return result
}

func ParseInput(path string) ([]Machine, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	result := []Machine{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			continue
		}

		// Parse Lights
		lightRe := regexp.MustCompile(`\[([.#]*)\]`)
		lightMatch := lightRe.FindStringSubmatch(str)
		if len(lightMatch) < 2 {
			return nil, fmt.Errorf("Failed to parse light: %s\n", str)
		}

		lights := Map(strings.Split(lightMatch[1], ""), func(e string) bool {
			if e == "#" {
				return true
			}
			return false
		})

		machine := Machine{Lights: lights}

		// Parse Buttons
		buttonRe := regexp.MustCompile(`\(([\d,]*)\)`)
		buttonMatches := buttonRe.FindAllStringSubmatch(str, -1)
		if len(buttonMatches) == 0 {
			return nil, fmt.Errorf("No buttons found: %s\n", str)
		}
		buttons := make([]ButtonSet, len(buttonMatches))
		for i, buttonMatch := range buttonMatches {
			match := buttonMatch[1]
			button := make([]int, len(lights))
			for nStr := range strings.SplitSeq(match, ",") {
				n, err := strconv.Atoi(nStr)
				if err != nil {
					return nil, err
				}
				button[n] = 1
			}
			buttons[i] = button
		}
		machine.Buttons = buttons

		// Parse Joltage
		joltageRe := regexp.MustCompile(`\{([\d,]*)\}`)
		joltageMatch := joltageRe.FindStringSubmatch(str)
		if len(joltageMatch) < 2 {
			return nil, fmt.Errorf("Failed to parse joltage: %s\n", str)
		}

		joltage := Map(strings.Split(joltageMatch[1], ","), func(e string) int {
			n, _ := strconv.Atoi(e)
			return n
		})
		machine.Joltage = joltage

		result = append(result, machine)
	}

	return result, nil
}

// generateCombinations is a recursive helper function to build combinations.
func generateCombinations(arr []ButtonSet, start int, currentCombination []ButtonSet, result *[][]ButtonSet) {
	// Add the current combination to the result (make a copy to avoid modification issues)
	temp := make([]ButtonSet, len(currentCombination))
	copy(temp, currentCombination)
	*result = append(*result, temp)

	// Iterate through the remaining elements to build new combinations
	for i := start; i < len(arr); i++ {
		// Include the current element in the combination
		currentCombination = append(currentCombination, arr[i])
		// Recursively call with the next element
		generateCombinations(arr, i+1, currentCombination, result)
		// Backtrack: remove the current element to explore other combinations
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}

func combinationSolvesLights(buttons []ButtonSet, lights LightSet) bool {
	lightState := make([]int, len(lights))
	for _, button := range buttons {
		for i, signal := range button {
			lightState[i] += signal
		}
	}
	// fmt.Println("Light State:", lightState)
	// fmt.Println("Desired state:", lights)

	for i, led := range lights {
		if led != (lightState[i]%2 != 0) {
			// fmt.Println("Did not match")
			// fmt.Println()
			return false
		}
	}
	// fmt.Println("Matched")
	// fmt.Println()
	return true
}

func solveMachineLights(machine Machine) int {
	buttonCombinations := [][]ButtonSet{}
	generateCombinations(machine.Buttons, 0, []ButtonSet{}, &buttonCombinations)
	result := 1000
	for _, combination := range buttonCombinations {
		if combinationSolvesLights(combination, machine.Lights) {
			if len(combination) < result {
				result = len(combination)
			}
		}
	}
	return result
}

func PartOne(machines []Machine) int {
	result := 0
	for _, machine := range machines {
		result += solveMachineLights(machine)
	}
	return result
}

func combinationSolvesJoltage(buttons []ButtonSet, joltage JoltageSpecs) bool {
	joltagestate := make([]int, len(joltage))
	for _, button := range buttons {
		for i, signal := range button {
			joltagestate[i] += signal
		}
	}

	return slices.Equal(joltagestate, joltage)
}

func solveMachineJoltage(machine Machine) int {
	var buttonCombinations [][]ButtonSet
	result := 1000
	for _, combination := range buttonCombinations {
		if combinationSolvesJoltage(combination, machine.Joltage) {
			if len(combination) < result {
				result = len(combination)
			}
		}
	}
	return result
}

func PartTwo(machines []Machine) int {
	result := 0
	for _, machine := range machines {
		result += solveMachineJoltage(machine)
	}
	return result
}
