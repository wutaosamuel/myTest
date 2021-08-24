package main

import (
	"fmt"
	"strconv"
)

// Object
type Object struct {
	Name  string
	ID    int
	Value float64
}

// NewObject
func NewObject() *Object {
	return &Object{
		Name:  "",
		ID:    -1,
		Value: -1,
	}
}

// InsertTo
func (o *Object) InsertTo() (string, string) {
	key := "name, id, value"
	value := fmt.Sprintf("'%s', %d, %s", o.Name, o.ID, strconv.FormatFloat(o.Value, 'f', -1, 64))
	return key, value
}

// SelectTo
func (o *Object) SelectTo() []interface{} {
	result := make([]interface{}, 3)
	result[0] = &o.Name
	result[1] = &o.ID
	result[2] = &o.Value
	return result
}

// UpdateTo
func (o *Object) UpdateTo() string {
	result := fmt.Sprintf("name = '%s', id = %d, value = %f", o.Name, o.ID, o.Value)
	return result
}
