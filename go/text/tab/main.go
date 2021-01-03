package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

func main() {
	var w *tabwriter.Writer
	w = &tabwriter.Writer{}
	//w := new(tabwriter.Writer)

	names := []string{
		"BTC",
		"ETH",
		"This is a test string",
	}

	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	for k, v := range names {
		fmt.Fprintln(w, v, "\t", k)
	}
	w.Flush()

	fmt.Println()
	buffer := bytes.NewBuffer(make([]byte, 0, 0))
	tw := new(tabwriter.Writer)
	tw.Init(os.Stdout, 0, 8, 0, '\t', 0)
	for k, v := range names {
		n := strconv.Itoa(k)
		buffer.WriteString(n + "\t")
		buffer.WriteString(v)
		buffer.WriteString("\t" + n)
		buffer.WriteByte('\n')
	}
	fmt.Fprintln(tw, buffer.String())
	tw.Flush()
}
