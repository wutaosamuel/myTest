package main

import (
	"errors"
	"fmt"
)

// ErrStr is a implementation of transfering between error && string
/* ErrStr
 *  - error <-> string
 *  -
 */
type ErrStr struct {
	str string
	err error
}

// NewErrStr create new error
func (e *ErrStr) NewErrStr() *ErrStr {
	return &ErrStr{}
}

// Error return string in error
func (e *ErrStr) Error() string {
	return e.str
}

// Errors return string && error
func (e *ErrStr) Errors() string {
	return e.str + " \\\n" + e.err.Error()
}

// Err is custom error custommed from string
func Err(text string) error {
	return errors.New(text)
}

// Errs is to merge a err and a string
func Errs(text string, err error) error {
	e := &ErrStr{text, err}
	return errors.New(e.Errors())
}

// CheckPanic check err != nil panic
/* CheckPanic
 * panic && print error only
 */
func CheckPanic(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
