package main

import (
	"fmt"
	gogit "github.com/go-git/go-git"
	//"github.com/go-git/go-git/plumbing"
)

func main() {
	var gitPath = "/home/pi/project/test"
	repo, err := gogit.PlainOpen(gitPath)
	panicError(err)
	remotes, err := repo.Remotes()
	panicError(err)
	//fmt.Println(remotes)
	for _, remote := range remotes {
		remoteConfig := remote.Config()
		fmt.Println(remoteConfig.Name)
		for _, url := range remoteConfig.URLs {
			fmt.Println(url)
		}
	}
}

// panicError panic error
func panicError(err error) {
	if err != nil {
		panic(err)
	}
}