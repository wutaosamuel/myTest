package main

import (
	"reflect"
	"fmt"
	"errors"
)

var Input = []string {"name1", "na", "rsa", "name4", "wt"}

type Object struct {
	Name string
	Id   int
}

func main() {
	o := Object{}
	res := Ref(o)
	fmt.Println(res)
	for _, v := range res {
		fmt.Println(v)
	}

	fmt.Println("TestCreateNewPtr --> pass")
	TestCreateNewPtr()
}

func Ref(model interface{}) []interface{} {
	result := make([]interface{}, 5)
	m := reflect.TypeOf(model)
	if m.Kind() == reflect.Ptr {
		m = m.Elem()
	}
	if m.Kind() != reflect.Struct {
		return nil
	}
	for n, v := range Input {
		s := reflect.New(m).Elem()
		num := s.NumField()
		col := make([]interface{}, num)
		for i := 0; i < num; i++ {
			field := s.Field(i)
			col[i] = field.Addr().Interface()
		}
		s.Field(0).SetString(v)
		s.Field(1).SetInt(int64(n))
		result[n] = s.Addr().Interface()
	}

	return result
}

func CreateNewPtr(model interface{}) (interface{}, error) {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	if modelType.Kind() != reflect.Struct {
		return nil, errors.New("query model type is not struct")
	}

	object := reflect.New(modelType).Elem()
	return object.Addr().Interface(), nil
}

// TestCreateNewPtr
func TestCreateNewPtr() {
	type Object struct {
		Name string
		ID int
	}

	p, err := CreateNewPtr(&Object{})
	if err != nil {
		fmt.Println(err)
		return
	}
	ptr, ok := p.(*Object)
	if !ok {
		fmt.Println("not ok")
		return
	}
	ptr.Name = "ptr"
	ptr.ID = -1
	fmt.Println(ptr)
}
