package converters

import (
	"github.com/pkg/errors"
	"strconv"
)

func StringToInt(s string) (int, error) {
	r, err := strconv.Atoi(s)

	if err != nil {
		return 0, errors.WithMessagef(err, "Cannot convert to int: %s", s)
	}

	return r, nil
}

func StringsToInts(ss ...string) ([]int, error) {
	var is []int
	for _, s := range ss {
		r, err := StringToInt(s)
		if err != nil {
			return []int{}, err
		}

		is = append(is, r)
	}

	return is, nil
}
