package main

import (
	"testing"
)

func TestToDockerFile(t *testing.T) {
	df := NewDockerFile()
	version := NewEnvKV()
	version.Set("VERSION", "0.0.1")
	golang := NewEnvKV()
	golang.Set("GOLANG_VERSION", "1.16.6")
	app := NewEnvKV()
	app.Set("APP", "DockerInGit")
	df.Env = []*EnvKV{version, golang, app}
	df.Run = []string{"apt update -y", "apt upgrade -y", "wget -v"}
	df.Ports = []string{"22/tcp", "80/tcp", "443"}
	df.Volumes = []string{"/mnt/user", "/mnt/git"}
	df.AppRun = []string{"bash /app/pre_install.sh", "bash /app/install.sh"}
	df.Cmd = []string{"/usr/local/go run /app/cmd/app", "command 2"}
	t.Log(df.ToDockerFile())

	// test single
	ds := NewDockerFile()
	ds.Env = []*EnvKV{app}
	ds.Run = []string{"apt update -y"}
	ds.Ports = []string{"22/tcp"}
	ds.Volumes = []string{"/mnt/user"}
	ds.AppRun = []string{"bash /app/install.sh"}
	ds.Cmd = []string{"/usr/local/go run /app/cmd/app"}
	//t.Log(ds.ToDockerFile())
}
