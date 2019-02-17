package lib

import (
	"fmt"
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"os/exec"
	"plugin"
	"reflect"
	"regexp"
	"strings"
)

func InvockConst(e expr, r *RunNode) (ret interface{}) {
	Debug("------------------------------------")
	switch e.method {
	case "println":
		fmt.Println(e.args)
	case "make":
		//TODO already make
		ret = e.args[0]
	case "append":
		ret = append(e.args[0].([]interface{}), e.args[1])
	case "new":
		switch e.args[0] {
		case "int":
			ret = new(int)
		default:
			Error("unsupport type %+v ", e.args[0])

		}
	default:
		if f := r.GetFunc(e.method); f != nil {
			rc := r.Child()
			rs := InvockCos(f, rc, e)
			if len(rs) == 1 {
				ret = rs[0]
			} else {
				ret = rs
			}
			return
		}
		Error("how it happend, method = %s , entity =%#v", e.method, e)
	}
	return
}

func Invock(e expr, r *RunNode) (ret interface{}) {
	if e.left == nil {
		return InvockConst(e, r)
	}
	switch l := e.left.(type) {
	case *plugin.Plugin:
		method, err := l.Lookup(e.method)
		if err != nil {
			Error("not found method %+v", e.method)
		}
		ins := make([]reflect.Value, 0)
		for _, i := range e.args {
			ins = append(ins, reflect.ValueOf(i))
		}
		rm := reflect.ValueOf(method)
		if !rm.IsValid() {
			Error("method fail, method =  %#v", method)
		}
		rret := rm.Call(ins)
		rets := make([]interface{}, 0)
		for _, rr := range rret {
			rets = append(rets, rr.Interface())
		}
		return rets
	case Struct:
		f, ok := l.funcMap[e.method]
		if !ok {
			Error("method not found %v , l = %#v", e.method, l)
			return
		}
		rc, ok := l.funcRMap[e.method]
		if ok {
			for k, v := range rc.VarMap {
				if v == SELF {
					rc.VarMap[k] = l
				}
			}
		}
		Debug("rc.VarMap = %#v", rc.VarMap)
		ret = InvockCos(f, rc, e)
	default:
		Error("unsupport left %#v", l)
	}

	return
}

func InvockCos(f *ast.FuncDecl, rc *RunNode, e expr) (ret []interface{}) {
	if f.Type.Params != nil {
		j := 0
		for _, field := range f.Type.Params.List {
			for _, n := range field.Names {
				if j >= len(e.args) {
					Error("o this is crazzy")
				}
				rc.VarMap[n.Name] = e.args[j]
				j++
			}
		}
	}
	if len(f.Body.List) > 0 {
		Debug("f.body---> %#v  \n\n", f.Body.List[0])
	}
	ret = CompileFuncDecl(f, rc)
	return
}

func InvockShell(stmt *ast.ShellStmt, r *RunNode) string {
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
		Error("err -> %v", err)
	}
	return string(ret)
}
