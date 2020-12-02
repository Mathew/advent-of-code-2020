package main

import (
	"github.com/mathew/advent-of-code-2020/internal/converters"
	"github.com/mathew/advent-of-code-2020/internal/files"
	"log"
	"regexp"
)

type PolicyValidator func(Password) bool

type Password struct {
	password string
	min      int
	max      int
	letter   string
}

func (p Password) IsValid(validator PolicyValidator) bool {
	return validator(p)
}

func parsePassword(s string) (Password, error) {
	regex := `(?P<min>\d+)-(?P<max>\d+)\s(?P<target>\w):\s(?P<password>\w+)`
	re := regexp.MustCompile(regex)
	matches := re.FindStringSubmatch(s)

	is, err := converters.StringsToInts(matches[1], matches[2])
	if err != nil {
		return Password{}, err
	}

	return Password{
		matches[4],
		is[0],
		is[1],
		matches[3],
	}, nil
}

func CountValidPasswords(passwords []string, validator PolicyValidator) (int, error) {
	c := 0
	for _, p := range passwords {
		pass, err := parsePassword(p)
		if err != nil {
			return 0, err
		}

		if pass.IsValid(validator) {
			c += 1
		}
	}
	return c, nil
}

func singleLetterValidator(p Password) bool {
	letterCount := 0
	for _, l := range p.password {
		if string(l) == p.letter {
			letterCount += 1
		}
	}

	return letterCount >= p.min && letterCount <= p.max
}

func letterPositionValidator(p Password) bool {
	isLetterOne := string(p.password[p.min - 1]) == p.letter
	isLetterTwo := string(p.password[p.max - 1]) == p.letter

	if isLetterOne && isLetterTwo {
		return false
	}

	if isLetterOne || isLetterTwo {
		return true
	}

	return false
}

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Fatalf("%+v", err)
		}
	}()

	rawInput, err := files.Load("cmd/day2/input.txt", "\n")
	if err != nil {
		return
	}

	count, err := CountValidPasswords(rawInput, singleLetterValidator)
	if err != nil {
		return
	}
	log.Print("Part One:")
	log.Printf("Valid password count: %d", count)

	count, err = CountValidPasswords(rawInput, letterPositionValidator)
	if err != nil {
		return
	}
	log.Print("Part Two:")
	log.Printf("Valid password count: %d", count)
}
