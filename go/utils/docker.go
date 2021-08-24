package main

import (
	"fmt"
	"encoding/json"
)

// cp && git submodule update --init --recursive
// remove submodule, if possible

// EnvKV
type EnvKV struct {
	Key   string
	Value string
}

// NewEnvKV
func NewEnvKV() *EnvKV {
	return &EnvKV{
		Key:   "",
		Value: "",
	}
}

// Set
func (e *EnvKV) Set(key, value string) {
	e.Key = key
	e.Value = value
}

// DockerFile
type DockerFile struct {
	TemporaryDir string
	Dir          string // dockerfile and volume location

	// dockerfile
	Version int
	Env     []*EnvKV
	Run     []string
	Ports   []string
	Volumes []string
	AppRun  []string
	Cmd     []string
}

// NewDockerFile
func NewDockerFile() *DockerFile {
	return &DockerFile{
		TemporaryDir: "/dev/shm",
		Dir:          "/mnt/user",

		Version: 0,
		Env:     make([]*EnvKV, 0),
		Run:     make([]string, 0),
		Ports:   make([]string, 0),
		Volumes: make([]string, 0),
		AppRun:  make([]string, 0),
		Cmd:     make([]string, 0),
	}
}

// Make

// Remove

// ToDockerFile
func (df *DockerFile) ToDockerFile() string {
	dockerFile := "FROM ubuntu:latest\n\n"
	dockerFile += df.toVersion()
	dockerFile += df.toEnv()
	dockerFile += df.toRun()
	dockerFile += df.toPorts()
	dockerFile += df.toVolumes()
	dockerFile += df.toAppRun()
	dockerFile += df.toCmd()

	return dockerFile
}

// JsonMarshal
func (df *DockerFile) JsonMarshal() ([]byte, error) {
	return json.Marshal(df)
}

// JsonUnmarshal
func (df *DockerFile) JsonUnmarshal(buffer []byte) error {
	return json.Unmarshal(buffer, df)
}

// toVersion
func (df *DockerFile) toVersion() string {
	return fmt.Sprintf("ARG VER_SION=0.0.%d\n\n", df.Version+1)
}

// toEnv
func (df *DockerFile) toEnv() string {
	env := "ENV "
	if len(df.Env) == 0 {
		return ""
	}
	for i, e := range df.Env {
		if i == len(df.Env)-1 {
			env += e.Key + "=" + e.Value
			continue
		}
		env += e.Key + "=" + e.Value + " \\\n"
	}
	env += "\n\n"

	return env
}

// toRun
func (df *DockerFile) toRun() string {
	run := "RUN "
	if len(df.Run) == 0 {
		return ""
	}
	if len(df.Run) == 1 {
		return run + df.Run[0] + "\n\n"
	}
	for i, r := range df.Run {
		if i == 0 {
			run += r + " \\\n"
			continue
		}
		if i == len(df.Run)-1 {
			run += "&& " + r
			continue
		}
		run += "&& " + r + " \\\n"
	}
	run += "\n\n"

	return run
}

// toPorts
func (df *DockerFile) toPorts() string {
	ports := "EXPOSE"
	if len(df.Ports) == 0 {
		return ""
	}
	for _, p := range df.Ports {
		ports += " " + p
	}
	ports += "\n\n"

	return ports
}

// toVolumes
func (df *DockerFile) toVolumes() string {
	volumes := "VOLUME ["
	if len(df.Volumes) == 0 {
		return ""
	}
	for i, p := range df.Volumes {
		if i == len(df.Volumes)-1 {
			volumes += "\"" + p + "\""
			continue
		}
		volumes += "\"" + p + "\", "
	}
	volumes += "]\n\n"

	return volumes
}

// toAppRun
func (df *DockerFile) toAppRun() string {
	run := "RUN "
	if len(df.AppRun) == 0 {
		return ""
	}
	if len(df.AppRun) == 1 {
		return run + df.AppRun[0] + "\n\n"
	}
	for i, r := range df.AppRun {
		if i == 0 {
			run += r + " \\\n"
			continue
		}
		if i == len(df.AppRun)-1 {
			run += "&& " + r
			continue
		}
		run += "&& " + r + " \\\n"
	}
	run += "\n\n"

	return run
}

// toCmd
func (df *DockerFile) toCmd() string {
	cmd := "CMD "
	if len(df.Cmd) == 0 {
		return ""
	}
	if len(df.Cmd) == 1 {
		return cmd + df.Cmd[0]
	}
	for i, c := range df.Cmd {
		if i == 0 {
			cmd += c + " \\\n"
			continue
		}
		if i == len(df.Cmd)-1 {
			cmd += "&& " + c
			continue
		}
		cmd += "&& " + c + " \\\n"
	}
	cmd += "\n\n"

	return cmd
}
