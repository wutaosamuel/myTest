package utils

import (
	"errors"
)

// ErrorString is a  trivial implementation of error
type ErrorString struct {
	s   string
	err error
}

// NewErr for new error
func NewErr() *ErrorString {
	return &ErrorString{}
}

// Error for an error
func (e *ErrorString) Error() string {
	return e.s
}

// Errors for all error
func (e *ErrorString) Errors() string {
	return e.s + e.err.Error()
}

// Err is custom error custommed from string
func Err(text string) error {
	return errors.New(text)
}

// Errs is to merge a err and a string
func Errs(text string, err error) error {
	errs := &ErrorString{text, err}
	return errors.New(errs.Errors())
}
