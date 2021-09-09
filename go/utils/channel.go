package utils

import (
	"errors"
)

// IsChannelError
func IsChannelError(err chan error) error {
	if e, ok := <-err; ok {
		return e
	} else {
		return errors.New("unexpected error channel closed")
	}
}
