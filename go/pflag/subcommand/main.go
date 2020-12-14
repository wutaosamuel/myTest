package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	var (
		helpFlag = pflag.BoolP("help", "h", false, "usage")
		cmd1    = pflag.NewFlagSet("cmd1", pflag.ExitOnError)
		cmd1Add = cmd1.BoolP("add", "a", false, "cmd1 add func")
		cmd2    = pflag.NewFlagSet("cmd2", pflag.ExitOnError)
		cmd2Add = cmd2.BoolP("add", "a", false, "cmd2")
	)

	pflag.Parse()
	cmd1.Parse()
	cmd2.Parse()
	if *helpFlag {
		pflag.Usage()
	}

	// parse
	switch os.Args[1] {
	case "cmd1":
		cmd1.Parse(os.Args[2:])
	case "cmd2":
		cmd2.Parse(os.Args[2:])
	default:
		pflag.Usage()
		cmd1.Usage()
		cmd2.Usage()
	}

	// need to custom usage
	fmt.Println(cmd2.FlagUsages())

	fmt.Println("cmd1:", *cmd1Add)
	fmt.Println("cmd2:", *cmd2Add)
}
