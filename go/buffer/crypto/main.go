package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func main() {
	obj := &Obj{
		name:    "Object",
		amount:  -0.5,
		time:    time.Now(),
		current: 100,
		past:    80.5}
	fmt.Println("1\t" + obj.objString())
	fmt.Println("\t" + obj.objString())
}

// Obj object
type Obj struct {
	name    string
	amount  float64
	time    time.Time
	current float64
	past    float64
}

// objString string of obj
func (o *Obj) objString() string {
	buf := &bytes.Buffer{}
	t := o.time.Format("2006-01-02")
	a := strconv.FormatFloat(o.amount, 'f', 6, 64)
	c := strconv.FormatFloat(o.current, 'f', 6, 64)
	p := strconv.FormatFloat(o.past, 'f', 6, 64)
	// a1 :=strconv.FormatFloat(o.amount, 'E', 8, 64)
	// fmt.Println("a1: ", a1)

	buf.WriteString(o.name)
	buf.WriteByte('\t')
	buf.WriteString(t)
	buf.WriteByte('\t')
	buf.WriteString(a)
	buf.WriteByte('\t')
	buf.WriteString(c)
	buf.WriteByte('\t')
	buf.WriteString(p)

	return buf.String()
}
