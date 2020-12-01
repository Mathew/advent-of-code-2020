package main

import (
	"github.com/mathew/advent-of-code-2020/internal/converters"
	"github.com/mathew/advent-of-code-2020/internal/files"
	"github.com/pkg/errors"
	"log"
)

func sum(is ...int) int {
	total := 0
	for _, i := range is {
		total += i
	}

	return total
}

func multiply(is ...int) int {
	total := is[0]
	for _, i := range is[1:] {
		total = total * i
	}

	return total
}

func checkAddition(arr []int, num, answer int) (int, bool) {
	for _, a := range arr {
		if a+num == answer {
			return a, true
		}
	}

	return 0, false
}

func findSumCombination(numbers []int, answer int) (int, int, error) {
	for i, num := range numbers {
		if match, ok := checkAddition(numbers[i+1:], num, answer); ok {
			return num, match, nil
		}
	}

	return 0, 0, errors.New("No match found.")
}

func findComponentSumCombination(numbers []int, answer, numComponents int, path []int) ([]int, bool) {
	if numComponents == 2 {
		for i, a := range numbers {
			if match, ok := checkAddition(numbers[i+1:], sum(path...)+a, answer); ok {
				path = append(path, a, match)
				return path, true
			}
		}
	} else {
		for i, a := range numbers {
			cp := path
			cp = append(cp, a)
			if result, ok := findComponentSumCombination(numbers[i+1:], answer, numComponents-1, cp); ok {
				return result, true
			}
		}
	}

	return nil, false
}

func main() {
	var err error

	defer func() {
		if err != nil {
			log.Fatalf("%+v", err)
		}
	}()

	rawInput, err := files.Load("cmd/day1/input.txt", "\n")
	if err != nil {
		return
	}

	ints, err := converters.StringsToInts(rawInput...)
	if err != nil {
		return
	}

	// Part 1.
	x, y, err := findSumCombination(ints, 2020)
	if err != nil {
		return
	}
	log.Printf("Part One:\n%d + %d = %d \n", x, y, x*y)

	// Part 2.
	components, ok := findComponentSumCombination(ints, 2020, 3, nil)
	log.Print("Part Two: ")
	if !ok {
		log.Printf("not found")
	} else {
		log.Printf("%v multiplied together = %v", components, multiply(components...))
	}
}
