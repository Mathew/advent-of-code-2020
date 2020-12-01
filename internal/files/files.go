package files

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Load(filePath string, separator string) ([]string, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, errors.WithMessagef(err, "can't locate directory: %s", filePath)
	}

	dat, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, errors.WithMessagef(err, "Can't read file: %s", dat)
	}

	if separator == "" {
		return []string{strings.TrimSpace(string(dat))}, nil
	}

	return strings.Split(string(dat), separator), nil
}
