package lib

import (
	"fmt"
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"github.com/spf13/cast"
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
			Error(" type = %+v ", e.args[0])

		}
	case "delete":
		if len(e.args) != 2 {
			Error("fail call delete %+v", e)
		}
		left, ok := e.args[0].(map[interface{}]interface{})
		if !ok {
			Error("fail call delete %+v", e)
		}
		delete(left, e.args[1])

	case "int64":
		return cast.ToInt64(e.args[0])
	case "string":
		return cast.ToString(e.args[0])
	case "[]byte":
		return cast.ToString(e.args[0])
	case "int8":
		return cast.ToInt8(e.args[0])
	case "int16":
		return cast.ToInt8(e.args[0])
	case "int32":
		return cast.ToInt8(e.args[0])
	case "uint64":
		return cast.ToUint64(e.args[0])
	case "uint32":
		return cast.ToUint32(e.args[0])
	case "uint16":
		return cast.ToUint16(e.args[0])
	case "uint8":
		return cast.ToUint8(e.args[0])
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
		Error(" method = %s , entity =%#v", e.method, e)
	}
	return
}

func InvockInternal(e expr, r *RunNode) (ret interface{}) {
	//删除internal func
	//build so
	//import internal func
	return nil
}

func Invock(e expr, r *RunNode) (ret interface{}) {
	if e.left == nil {
		return InvockConst(e, r)
	}
	switch l := e.left.(type) {
	case *plugin.Plugin:
		method, err := l.Lookup(e.method)
		if err != nil {
			Error("method = %+v", e.method)
		}
		ins := make([]reflect.Value, 0)
		for _, i := range e.args {
			if i != nil {
				ins = append(ins, reflect.ValueOf(i))
			} else {
				ev := getEmpty()
				ins = append(ins, reflect.ValueOf(ev))
				Debug("kind = %#v,nil? = %#v", reflect.ValueOf(ev).Kind(), ev == nil)
			}
		}
		rm := reflect.ValueOf(method)
		if !rm.IsValid() {
			Error("method =  %#v", method)
		}
		//if rm.Type().NumIn() > len(ins) {
		//	Error("invalid method call in = %v oin = %v ", rm.Type().NumIn(), len(ins))
		//}
		rret := rm.Call(ins)
		rets := make([]interface{}, 0)
		for _, rr := range rret {
			rets = append(rets, rr.Interface())
		}
		return rets
	case Struct:
		f, ok := l.funcMap[e.method]
		if !ok {
			Error("method = %v , l = %#v", e.method, l)
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
		Error("left = %#v", l)
	}

	return
}

func InvockCos(f *ast.FuncDecl, rc *RunNode, e expr) (ret []interface{}) {
	if f.Type.Params != nil {
		j := 0
		for _, field := range f.Type.Params.List {
			for _, n := range field.Names {
				if j >= len(e.args) {
					Error("len error")
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
	Debug("ret = %#v", ret)
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
		Error("err = %v", err)
	}
	return string(ret)
}
