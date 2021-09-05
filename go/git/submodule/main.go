package main

import (
	"fmt"
	//"os"

	gogit "github.com/go-git/go-git"
	//gogit "github.com/go-git/go-git/v5"
)

func main() {
	tmpDir := "/dev/shm/power"
	//power := "https://github.com/microsoft/PowerToys.git"
	//r, err := gogit.PlainClone(tmpDir, false, &gogit.CloneOptions{
	//URL:      power,
	//Progress: os.Stdout,
	//})
	//if err != nil {
	//fmt.Println(err)
	//return
	//}
	r, err := gogit.PlainOpen(tmpDir)
	if err != nil {
		fmt.Println("1", err)
		return
	}
	w, err := r.Worktree()
	if err != nil {
		fmt.Println("2", err)
		return
	}
	subs, err := w.Submodules()
	if err != nil {
		fmt.Println("3", err)
		return
	}
	if err := subs.Update(&gogit.SubmoduleUpdateOptions{}); err != nil {
		fmt.Println("4", err)
		return
	}
}
