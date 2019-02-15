package lib

import (
	"go/ast"
)

func CompileVarSpec(stmt *ast.ValueSpec, r *RunNode) {
	for i, n := range stmt.Names {
		Debug("==> name : %+v, value : %#v \n", n.Name, stmt.Values)
		v := Var{}
		v.k = n.Name
		if len(stmt.Values) > i {
			v.v = stmt.Values[i]
		} else if len(stmt.Values) > 0 {
			v.v = stmt.Values[0]
		}
		r.Vars = append(r.Vars, v)
		if r.VarMap == nil {
			r.VarMap = make(map[string]interface{}, 0)
		}
		r.VarMap[n.Name] = CompileExpr(v.v, r)
	}
}

func CompileTypeSpec(stmt *ast.TypeSpec, r *RunNode) {
	if r.TypeMap == nil {
		r.TypeMap = make(map[string]ast.Expr, 0)
	}
	r.TypeMap[stmt.Name.Name] = stmt.Type
}

func CompileImportSpec(spe *ast.ImportSpec, r *RunNode) {
	//Debug(" =======>> %s\n", spe.Path.Value)
	//TODO
	// 1\ import so
	// 2\ import r

}
