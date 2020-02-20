package utils

import (
	"io/ioutil"
)

// ParseFile get string from a file
func ParseFile(path string) (string, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
