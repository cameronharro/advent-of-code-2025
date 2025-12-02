package daytwo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(path string) ([][]int, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return [][]int{}, err
	}

	ranges := strings.Split(strings.Trim(string(bytes), "\n"), ",")
	result := make([][]int, len(ranges))
	for i, r := range ranges {
		strs := strings.Split(r, "-")

		if len(strs) != 2 {
			return [][]int{}, fmt.Errorf("Invalid range: %v", strs)
		}
		bottom, err := strconv.Atoi(strs[0])
		if err != nil {
			return [][]int{}, fmt.Errorf("Invalid range bound: %s - %v", strs[0], err.Error())
		}
		top, err := strconv.Atoi(strs[1])
		if err != nil {
			return [][]int{}, fmt.Errorf("Invalid range bound: %s - %v", strs[1], err.Error())
		}
		thisResult := []int{bottom, top}
		result[i] = thisResult
	}
	return result, nil
}

func sumRange(r []int, invalidTest func(n int) bool) (int, error) {
	if len(r) != 2 {
		return 0, fmt.Errorf("Invalid range: %v", r)
	}

	if !(r[0] < r[1]) {
		return 0, fmt.Errorf("Invalid range: %v", r)
	}

	result := 0
	for i := r[0]; i <= r[1]; i++ {
		if invalidTest(i) {
			result += i
		}
	}
	return result, nil
}

func Part(invalidTest func(n int) bool) func(input [][]int) (int, error) {
	return func(input [][]int) (int, error) {
		result := 0
		for _, r := range input {
			sum, err := sumRange(r, invalidTest)
			if err != nil {
				return 0, err
			}
			result += sum
		}

		return result, nil
	}
}

func PartOneValidCheck(n int) bool {
	str := strconv.Itoa(n)
	if len(str)%2 != 0 {
		return false
	}
	if str[:len(str)/2] == str[len(str)/2:] {
		return true
	}
	return false
}

func PartTwoValidCheck(n int) bool {
	str := strconv.Itoa(n)
	for i := 0; i < len(str)/2; i++ {
		matched := true
		fieldLength := i + 1
		if len(str)%fieldLength != 0 {
			continue
		}
		fields := len(str) / fieldLength
		currentField := str[0:fieldLength]
		for j := 1; j < fields; j++ {
			if currentField != str[j*fieldLength:(j+1)*fieldLength] {
				matched = false
			}
		}
		if matched {
			return true
		}
	}
	return false
}
