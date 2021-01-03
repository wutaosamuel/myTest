package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(buildBuffer())
}

// buildBuffer build
func buildBuffer() string {
	buf := &bytes.Buffer{}
	// buf := bytes.NewBuffer(make([]byte, 0, size))
	amount := strconv.FormatFloat(0.56287426, 'f', 6, 64)
	past := strconv.FormatFloat(0.48587426, 'f', 6, 64)
	current := strconv.FormatFloat(0.85687426, 'f', 6, 64)
	buf.WriteString("This is a test")
	buf.WriteByte(':')
	buf.WriteByte('\n')
	// main
	buf.WriteString("Test")
	buf.WriteByte('\t')
	buf.WriteString(amount)
	buf.WriteByte('\t')
	buf.WriteString(past)
	buf.WriteByte('\t')
	buf.WriteString(current)
	buf.WriteByte('\n')
	// underscore
	buf.WriteString("---------\n")
	// main2
	buf.WriteString("Test")
	buf.WriteByte('\t')
	buf.WriteString(amount)
	buf.WriteByte('\t')
	buf.WriteString(past)
	buf.WriteByte('\t')
	buf.WriteString(current)
	buf.WriteByte('\n')
	// underscore
	buf.WriteString("---------")
	buf.WriteByte('\n')
	buf.Write(main3String())
	buf.WriteString(main4String())

	return buf.String()
}

// main3
func main3String() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 25))
	buf.WriteString("BufferString")
	buf.WriteByte('\t')
	buf.WriteString("0.065")
	buf.WriteByte('\t')
	buf.WriteString("20")
	buf.WriteByte('\t')
	buf.WriteString("30")
	buf.WriteByte('\n')
	
	return buf.Bytes()
}

// main4
func main4String() string {
	buf := bytes.NewBuffer(make([]byte, 0, 18))
	buf.WriteString("bufferString")
	buf.WriteByte('\t')
	buf.WriteString("0.065")
	buf.WriteByte('\t')
	buf.WriteString("20")
	buf.WriteByte('\t')
	buf.WriteString("30")
	buf.WriteByte('\n')
	
	return buf.String()
}