package main

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// DeleteElementIndexes delete slice elements by indexes
func DeleteElementIndexes(slice interface{}, indexes []int) (interface{}, error) {
	// check slice variables is a slice
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return nil, errors.New("Slice must be slice")
	}
	// initialize
	value := reflect.ValueOf(slice)
	s := make([]interface{}, value.Len())
	for i := 0; i < value.Len(); i++ {
		s[i] = value.Index(i).Interface()
	}
	if len(indexes) == 0 || indexes == nil {
		return s, nil
	}

	// sort indexes slice
	sort.Ints(indexes)
	// remove duplicate && lower index && higher index
	deduplicate := -1
	index := make([]int, 0)
	for _, v := range indexes {
		if v < 0 {
			continue
		}
		if deduplicate == v {
			continue
		}
		if v > value.Len() {
			break
		}
		deduplicate = v
		index = append(index, v)
	}

	// get new slice
	length := value.Len() - len(index)
	result := make([]interface{}, length)
	for i, j := 0, 0; i < value.Len(); i++ {
		if j < len(index) && i == index[j]{
			j++
			continue
		}
		result[i-j] = s[i]
	}

	return result, nil
}

// obj object
type obj struct {
	Name string
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("Input: ", a)
	r, _ := DeleteElementIndexes(a, []int{2, 3})
	fmt.Println("Output: ", r)

	o := []*obj{&obj{"N1"}, &obj{"N2"}}
	fmt.Println("input: ", o)
	res, _ := DeleteElementIndexes(o, []int{0})
	for i := 0; i < len(res); i++ {
		fmt.Println("output: ", res[i])
	}
}
