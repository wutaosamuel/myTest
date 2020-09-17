package main

import (
	"fmt"
	gogit "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"time"
)

func main() {
	dir := "./test"
	repo, err := gogit.PlainOpen(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	tree, err := repo.Worktree()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := tree.AddGlob(dir); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("add done")
	commit, err := tree.Commit("test1", &gogit.CommitOptions{
		Author: &object.Signature{
			Name: "name",
			Email: "email",
			When: time.Now(),
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	// verify commit job
	obj, err := repo.CommitObject(commit)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(obj)
}