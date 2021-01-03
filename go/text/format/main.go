package main

import (
	"fmt"
	"strconv"
	"bytes"
	"text/tabwriter"
	"os"
)

func main() {
	names := []string{
		"BTC",
		"ETH",
		"This is a test string",
	}
	buffer := bytes.NewBuffer(make([]byte, 0, 0))
	buffer.WriteString("Formater string\n")
	for k, v := range names {
		n := strconv.Itoa(k)
		buffer.WriteString(n+"\t")
		buffer.WriteString(v)
		buffer.WriteString("\t"+n)
		buffer.WriteByte('\n')
	}
	fmt.Println(buffer.String())

	fmt.Println("fmt.sprintf: ")
	for k, v := range names {
		fmt.Println(fmt.Sprintf("%-20s\t%d", v, k))
	}

	fmt.Println("tab writer: ")
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	w.Write(buffer.Bytes())
	w.Flush()

}