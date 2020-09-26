package main

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Println(buildBuffer())
}

// buildBuffer build
func buildBuffer() string {
	buf := &bytes.Buffer{}
	// buf := bytes.NewBuffer(make([]byte, 0, size))
	buf.WriteString("This is a test")
	buf.WriteByte(':')
	buf.WriteByte('\n')
	buf.WriteByte('\t')
	buf.WriteString("Done")

	return buf.String()
}