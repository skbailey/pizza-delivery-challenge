package reader

import (
	"io/ioutil"
	"pizza-delivery/internal/pkg/validation"
	"strings"
)

// Read reads input from a file
func Read(path string) (string, error) {
	err := validation.ValidateReadParams(path)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	input := string(bytes)
	return strings.TrimSpace(input), nil
}
