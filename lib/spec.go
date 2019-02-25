package lib

import (
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"plugin"
	"strings"
)

var pluginLib = "/home/shizuo/.go-shell/plugin-lib/"

func CompileVarSpec(stmt *ast.ValueSpec, r *RunNode) {
	for i, n := range stmt.Names {
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
		Debug("==> value : %#v \n", v.v)
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
	pluginName := strings.Replace(strings.Replace(spe.Path.Value, "\"", "", 2), "/", "@@@@@@", 20) + ".so"
	path := ""
	if isExist(defaultPluginLib + "/" + pluginName) {
		path = defaultPluginLib + "/" + pluginName
	} else {
		path = installPath + "/" + pluginName
		if !isExist(installPath+"/"+pluginName) || FlushImport {
			autoImport(strings.Replace(spe.Path.Value, "\"", "", 2))
		}
	}
	gdll, err := plugin.Open(path)
	Debug("gddl = %#v,err = %#v", gdll, err)
	if err != nil {
		Error("path = %#v,err = %#v", path, err)
	}
	arr := strings.Split(strings.Replace(spe.Path.Value, "\"", "", 2), "/")

	importMap[arr[len(arr)-1]] = gdll
}
