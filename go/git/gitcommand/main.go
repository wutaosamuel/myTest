package main

import (
	"fmt"
	"regexp"
	"os"
	"os/exec"
	"strings"
)

func main() {
	gitPath := "/home/pi/project/test"
	command := "git remote show all"

	// replace ${var} or $var in shell
	commandToRun := os.ExpandEnv(command)
	// separate command to args in terms of spaces
	args := strings.Fields(strings.TrimSpace(commandToRun))
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = gitPath
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(GetBranchStatus(string(output)))
}

// GetBranchStatus get branch status against remote branch
func GetBranchStatus(message string) []string {
	var result []string
	cutMark := "  Local refs configured for 'git push':"
	split := strings.Split(message, "\n")
	canAdd := false
	rgx := regexp.MustCompile(`\((.*?)\)`)
	for k, v := range split {
		if canAdd {
			res := rgx.FindString(v)
			result = append(result, res)
		}
		if v == cutMark {
			canAdd = true
			fmt.Println(k)
			fmt.Println(v)
		}
	}
	return result
}