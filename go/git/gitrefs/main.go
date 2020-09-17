package main

import (
	"fmt"
	"time"
	gogit "github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing"
	"github.com/go-git/go-git/plumbing/object"
)

func main() {
	var gitPath = "/home/pi/project/test"
	repo, err := gogit.PlainOpen(gitPath)
	if err != nil {
		panic(err)
	}
	var gitRefs = []*plumbing.Reference{}
	refs, _ := repo.References()
	refs.ForEach(func (ref *plumbing.Reference) error {
		gitRefs = append(gitRefs, ref)
		return nil
	})
	for _, r := range gitRefs {
		fmt.Println(r)
	}

	var gitCommits = []*object.Commit{}
	since := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	until := time.Date(2020, 9, 11, 0, 0, 0, 0, time.UTC)
	commits, _ := repo.Log(
		&gogit.LogOptions{
			From: gitRefs[2].Hash(),
			Since: &since,
			Until: &until})
	commits.ForEach(func (commit *object.Commit) error {
		gitCommits = append(gitCommits, commit)
		return nil
	})
}