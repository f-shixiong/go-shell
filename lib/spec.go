package lib

import (
	"go/ast"
	"plugin"
	"strings"
)

var pluginLib = "/home/shizuo/.go-shell/plugin-lib/"

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
		//r.Vars = append(r.Vars, v)
		if r.VarMap == nil {
			r.VarMap = make(map[string]interface{}, 0)
		}
		r.VarMap[n.Name] = CompileExpr(v.v, r)
	}
}

func CompileTypeSpec(stmt *ast.TypeSpec, r *RunNode) {
	if r.TypeMap == nil {
		r.TypeMap = make(map[string]interface{}, 0)
	}
	r.TypeMap[stmt.Name.Name] = CompileExpr(stmt.Type, r)
}

func CompileImportSpec(spe *ast.ImportSpec, r *RunNode) {
	//Debug(" =======>> %s\n", spe.Path.Value)
	//TODO can't find
	Debug("import what name : %#v, value :%s", spe.Name, spe.Path.Value)
	path := pluginLib + strings.Replace(spe.Path.Value, "\"", "", 2) + ".so"
	gdll, err := plugin.Open(path)
	Debug("gddl = %#v,err = %#v", gdll, err)
	if err != nil {
		Error("path = %#v,err = %#v", path, err)
	}
	if r.ImportMap == nil {
		r.ImportMap = make(map[string]plugin.Symbol, 0)
	}
	r.ImportMap[strings.Replace(spe.Path.Value, "\"", "", 2)] = gdll
}
