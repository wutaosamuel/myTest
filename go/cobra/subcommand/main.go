package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var printCmd = &cobra.Command {
		Use: "print [string to print]",
		Short: "print anything to the screen",
		Long: "print is for printing anything back to screen.",
		// Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}
	
	var echoCmd = &cobra.Command{
    Use:   "echo [string to echo]",
    Short: "Echo anything to the screen",
    Long: `echo is for echoing anything back.
			Echo works a lot like print, except it has a child command.`,
    //Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Echo: " + strings.Join(args, " "))
    },
	}

	var rootCmd = &cobra.Command{
		Use: "app",
		Version: "0.0.1",
		// TODO: required, if without subcommand
		// Run: func(cmd *cobra.Command, args []string) {},
	}

	// global flag
	var global string
	rootCmd.PersistentFlags().StringVarP(&global, "global", "g", "", "global variable")

	// local flag
	var local string
	echoCmd.Flags().StringVarP(&local, "local", "l", "", "local variable")
	
	//rootCmd.AddCommand(printCmd, echoCmd)
	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()

	fmt.Println(global)
	fmt.Println(local)
	// fmt.Print(rootCmd.UsageString())
}