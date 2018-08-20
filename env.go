package errors

import (
	"os"
	"strings"
)

func goEnv() []string {
	env := make([]string, 0, 2)
	// goroot
	goroot := os.Getenv("GOROOT")
	if len(goroot) > 0 {
		if strings.Contains(goroot, `\`) && strings.Contains(goroot, ":") { // win
			goroot = strings.Replace(goroot, `\`, "/", -1)
		}
		env = append(env, goroot)
	}
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		return env
	}
	if strings.Contains(gopath, `\`) && strings.Contains(gopath, ":") { // win
		gopath = strings.Replace(gopath, `\`, "/", -1)
		if strings.Contains(gopath, ";") {
			gopaths := strings.Split(gopath, ";")
			for _, item := range gopaths {
				env = append(env, strings.TrimSpace(item))
			}
		} else {
			env = append(env, strings.TrimSpace(gopath))
		}
	} else { // unix
		if strings.Contains(gopath, ":") {
			gopaths := strings.Split(gopath, ":")
			for _, item := range gopaths {
				env = append(env, strings.TrimSpace(item))
			}
		} else {
			env = append(env, strings.TrimSpace(gopath))
		}
	}
	return env
}
