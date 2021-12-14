package main

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"

	gogit "gopkg.in/src-d/go-git.v4"
	gitClient "gopkg.in/src-d/go-git.v4/plumbing/transport/client"
	gitHttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func ProxyClone(dir, gitUrl, proxyUrl string) {
	gitName := filepath.Base(gitUrl) // need to check .git
	// if fileExt := filepath.Ext(gitName) // check .git extension
	gitDir := filepath.Join(dir, gitName[:len(gitName)-4])
	fmt.Println(gitDir)
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		fmt.Println("url parse fail,", err)
		return
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}
	client := &http.Client{
		Transport: transport,
	}
	gitClient.InstallProtocol("https", gitHttp.NewClient(client))
	_, err = gogit.PlainClone(gitDir, false, &gogit.CloneOptions{URL: gitUrl})
	if err != nil {
		fmt.Println("gogit plain clone fail,", err)
	}
}
