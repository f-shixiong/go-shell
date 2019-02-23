package lib

import (
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"io"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

type shellRunTime struct {
	path string
}

func InvockShell(stmt *ast.ShellStmt, r *RunNode) string {
	parallelArr := strings.Split(stmt.Cmd, "||")
	for i, pc := range parallelArr {
		go InvockSerialShell(pc, r)
		if i == len(parallelArr)-1 {

			return InvockSerialShell(pc, r)
		}
	}
	return ""
}

func InvockSerialShell(cmd string, r *RunNode) string {
	serialArr := strings.Split(cmd, "&&")
	sr := &shellRunTime{}
	for i, sc := range serialArr {
		InvockSingleShell(sc, r, sr)
		if i == len(serialArr)-1 {
			return InvockSingleShell(sc, r, sr)
		}
	}
	return ""
}

func InvockSingleShell(cmdr string, r *RunNode, sr *shellRunTime) string {
	reg, _ := regexp.Compile("\\s{2,}")
	rcmd := cmdr[1 : len(cmdr)-1]
	str := string(reg.ReplaceAll([]byte(rcmd), []byte("\\s")))
	Debug("srt -> %s", str)
	pipArr := strings.Split(str, "|")

	var outer io.Reader
	_ = outer
	var ret []byte
	var err error
	var l = len(pipArr)
	for i, estr := range pipArr {
		if i == l-1 {
			arr := strings.Split(estr, " ")
			var brr []string
			for i, v := range arr {
				if i != 0 {
					brr = append(brr, v)
				}
			}
			cmd := exec.Command(arr[0], brr...)
			setPathEnv(cmd, sr)
			if l != 1 && outer != nil {
				initCmd(cmd, outer)
			}
			ret, err = cmd.Output()
			if err != nil {
				Error("err = %v", err)
			}
		} else {
			arr := strings.Split(estr, " ")
			var brr []string
			for i, v := range arr {
				if i != 0 {
					brr = append(brr, v)
				}
			}
			if arr[0] == "cd" && len(brr) > 0 {
				sr.path = brr[0]
				continue
			}
			cmd := exec.Command(arr[0], brr...)
			setPathEnv(cmd, sr)
			if outer != nil {
				initCmd(cmd, outer)
			}
			outer, err = cmd.StdoutPipe()
			if err != nil {
				Error("err = %v", err)
			}

			err = cmd.Start()
			if err != nil {
				Error("err = %v", err)
			}

		}

	}

	return string(ret)
}

func initCmd(cmd *exec.Cmd, outer io.Reader) {
	inter, e0 := cmd.StdinPipe()
	if e0 != nil {
		Error("e0 = %v", e0)
	}
	dt, e1 := ioutil.ReadAll(outer)
	if e1 != nil {
		Error("e1 = %v", e1)
	}
	_, e2 := inter.Write(dt)
	inter.Close()
	if e2 != nil {
		Error("e2 = %v", e2)
	}

}

func setPathEnv(cmd *exec.Cmd, sr *shellRunTime) {
	if sr.path != "" {
		cmd.Path = sr.path
	}
}
