package main

import (
	"github.com/mathew/advent-of-code-2020/internal/converters"
	"github.com/mathew/advent-of-code-2020/internal/files"
	"github.com/pkg/errors"
	"log"
)

func checkAddition(arr []int, num, answer int) (int, bool) {
	for _, a := range arr {
		if a + num == answer {
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

	x, y, err := findSumCombination(ints, 2020)
	if err != nil {
		return
	}

	log.Printf("%d + %d = %d", x, y, x*y)
}
