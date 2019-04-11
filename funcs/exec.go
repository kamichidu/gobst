package funcs

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

type commandResult map[string]interface{}

func (v commandResult) String() string {
	s, _ := v["stdout"].(string)
	return s
}

func command(args ...interface{}) commandResult {
	if len(args) == 0 {
		panic(pkgName + ": command requires at least 1 argument")
	}
	strArgs := make([]string, len(args))
	for i := range args {
		strArgs[i] = fmt.Sprint(args[i])
	}
	cmd := exec.Command(strArgs[0], strArgs[1:]...)
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	out := map[string]interface{}{}
	out["stdout"] = stdout.String()
	out["stderr"] = stderr.String()
	if ee, ok := err.(*exec.ExitError); ok {
		out["error"] = ee.Error()
		if ws, ok := ee.Sys().(syscall.WaitStatus); ok {
			out["exitCode"] = ws.ExitStatus()
		} else {
			out["exitCode"] = -1
		}
	} else if err != nil {
		out["error"] = err.Error()
		out["exitCode"] = -1
	} else {
		out["exitCode"] = 0
	}
	return out
}

func init() {
	Register("command", command)
}
