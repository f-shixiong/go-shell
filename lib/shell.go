package lib

import (
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"io"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

func InvockShell2(stmt *ast.ShellStmt, r *RunNode) string {
	reg, _ := regexp.Compile("\\s{2,}")
	rcmd := stmt.Cmd[1 : len(stmt.Cmd)-1]
	str := string(reg.ReplaceAll([]byte(rcmd), []byte("\\s")))
	arr := strings.Split(str, " ")
	Debug("srt -> %s, arr %+v", str, arr)
	var brr []string
	for i, v := range arr {
		if i != 0 {
			brr = append(brr, v)
		}
	}
	cmd := exec.Command(arr[0], brr...)
	ret, err := cmd.Output()
	if err != nil {
		Error("err = %v", err)
	}
	return string(ret)
}

func InvockShell(stmt *ast.ShellStmt, r *RunNode) string {
	reg, _ := regexp.Compile("\\s{2,}")
	rcmd := stmt.Cmd[1 : len(stmt.Cmd)-1]
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
			cmd := exec.Command(arr[0], brr...)
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
