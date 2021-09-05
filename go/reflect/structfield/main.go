package main

import (
	"fmt"
	"reflect"
)

type Object struct {
	Name string
	Id   int
}

func main() {
	o := &Object{Name: "n", Id: 1}
	m := FieldMap(o)
	fmt.Println(m)
}

func FieldMap(model interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	m := reflect.TypeOf(model)
	v := reflect.ValueOf(model)
	if m.Kind() == reflect.Ptr {
		m = m.Elem()
		v = v.Elem()
	}
	if m.Kind() != reflect.Struct {
		return result
	}
	num := m.NumField()
	for i := 0; i < num; i++ {
		field := m.Field(i)
		value := v.Field(i)
		result[field.Name] = value.Interface()
		fmt.Println("type:", field.Type.Kind())
		fmt.Println("String:", fmt.Sprint(value.Interface()))
	}
	return result
}
