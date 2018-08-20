package errors

import (
	"path"
	"runtime"
	"strings"
	"time"
)

func fileName(src string) (goPath string, file string) {
	file = src
	goHomes := goEnv()
	if goHomes == nil {
		return
	}
	for _, goHome := range goHomes {
		if strings.Contains(file, goHome) {
			goPath = goHome
			file = strings.Replace(file, path.Join(goHome, "src"), "", 1)[1:]
			return
		}
	}
	return
}

func timeNow() time.Time {
	return time.Now().In(_cfg.loc)
}

func callers() []uintptr {
	pcs := make([]uintptr, Configure().depth)
	n := runtime.Callers(Configure().skip, pcs[:])
	return pcs[0:n]
}

func callersByAssigned(depth int, skip int) []uintptr {
	pcs := make([]uintptr, depth)
	n := runtime.Callers(skip, pcs[:])
	return pcs[0:n]
}
