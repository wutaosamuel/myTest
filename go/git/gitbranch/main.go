package main

import (
	"fmt"
	gogit "github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing"
)

func main() {
	var gitPath = "/home/pi/project/test"
	repo, err := gogit.PlainOpen(gitPath)
	b, e := repo.Branch("home")
	fmt.Println(b, e)
	panicError(err)
	branches, err := repo.Branches()
	panicError(err)
	//fmt.Println(branches)
	branches.ForEach(func(ref *plumbing.Reference) error {
		fmt.Println(ref)
		return nil
	})
}

// panicError panic error
func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

// branchGit by git command
func branchGit() {
	
}