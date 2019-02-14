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
	default:
		if f, ok := r.FuncMap[e.method]; ok {
			InvockCos(f, r, e)
			return
		}
		Error("how it happend, method = %s , entity =%#v", e.method, e)
	}
	return
}

func InvockImport() {

}

func InvockCos(f *ast.FuncDecl, r *RunNode, e expr) {
	sr := &RunNode{
		Father: r,
		VarMap: make(map[string]interface{}, 0),
	}
	_ = sr
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

	Debug("f.body---> %#v  \n\n", f.Body.List[0].(*ast.ExprStmt).X.(*ast.CallExpr).Args[0])
	CompileFuncDecl(f, sr)
}
