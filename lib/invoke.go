package lib

import (
	"fmt"
	"go/ast"
	_ "reflect"
)

func Invock(e expr, r *RunNode) (ret interface{}) {
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
			rs := InvockCos(f, r, e)
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

func InvockImport() {

}

func InvockCos(f *ast.FuncDecl, r *RunNode, e expr) (ret []interface{}) {
	sr := &RunNode{
		Father: r,
		VarMap: make(map[string]interface{}, 0),
	}
	if f.Type.Params != nil {
		j := 0
		for _, field := range f.Type.Params.List {
			for _, n := range field.Names {
				if j >= len(e.args) {
					Error("o this is crazzy")
				}
				sr.VarMap[n.Name] = e.args[j]
				j++
			}
		}
	}
	if len(f.Body.List) > 0 {
		Debug("f.body---> %#v  \n\n", f.Body.List[0])
	}
	ret = CompileFuncDecl(f, sr)
	Test("e -> %#v, ret-> %#v", e, ret)
	return
}
