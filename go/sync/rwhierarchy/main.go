package main

import (
	"fmt"
	"sync"
)

// Obj test obj
type Obj struct {
	Val	int
	rw *sync.RWMutex
}

// NewObj new
func NewObj() *Obj {
	return &Obj{
		Val: -1,
		rw: &sync.RWMutex{}}
}

// Write write val
func (obj *Obj)Write(val int) {
	obj.rw.Lock()
	defer obj.rw.Unlock()
	obj.Val = val

	return
}

// Read read val
func (obj *Obj)Read() int {
	obj.rw.RLock()
	defer obj.rw.RUnlock()

	return obj.Val
}

// Objects contain obj
type Objects struct {
	Name string
	O []*Obj
	rw *sync.RWMutex
}

// NewObjects new
func NewObjects() *Objects {
	return &Objects{
		Name: "Name0",
		O: make([]*Obj, 0),
		rw: &sync.RWMutex{}}
}

// Write write obj
func (o *Objects)Write(val []int) {
	o.rw.Lock()
	defer o.rw.Unlock()
	for _, v := range val {
		obj := NewObj()
		obj.Write(v)
		o.O = append(o.O, obj)
	}

	return
}

// Read read obj
func (o *Objects)Read() {
	o.rw.RLock()
	defer o.rw.RUnlock()
	for _, v := range o.O {
		fmt.Println(v.Read())
	}

	return
}

func main() {
	objects := NewObjects()
	vals := []int{1,2,3,4,5,6,7,8,9}
	var wg sync.WaitGroup

	fmt.Println("Start to test")
	wg.Add(1)
	go func() {
		defer wg.Done()
		objects.Write(vals)
	}()
	wg.Wait()
	fmt.Println("Write done")
	fmt.Println("Read Start")
	wg.Add(2)
	for i:=0; i <2; i++ {
		go func() {
			defer wg.Done()
			objects.Read()
		}()
	}
	wg.Wait()
	fmt.Println("Read done")
}