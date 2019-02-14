package lib

import (
	"go/ast"
	"reflect"
)

func CompileAssignStmt(stmt *ast.AssignStmt, r *RunNode) {
	if r.VarMap == nil {
		r.VarMap = make(map[string]interface{}, 0)
	}
	for i, k := range stmt.Lhs {
		right := stmt.Rhs[0]
		if len(stmt.Rhs) >= i {
			right = stmt.Rhs[i]
		}
		switch k := k.(type) {
		case *ast.Ident:
			if len(stmt.Rhs) >= i {
				//*ast.BasicLit
				r.VarMap[k.Name] = CompileExpr(right, r)
			} else {
				r.VarMap[k.Name] = CompileExpr(right, r)

			}
		case *ast.IndexExpr:
			Debug("test--> %#v,%#v", k.X, k.Index)
			Debug("test if com --> %#v", CompileExpr(k.Index, r))
			//TODO 1l
			o, ok := r.VarMap[k.X.(*ast.Ident).Name]
			if !ok {
				Error("o no so sad ")
				return
			}
			mp := o.(map[interface{}]interface{})
			mp[CompileExpr(k.Index, r)] = CompileExpr(right, r)
			r.VarMap[k.X.(*ast.Ident).Name] = mp
			Debug("$%+v", CompileExpr(right, r))

		default:
			Error("o no , this is what %#v ", k)
		}
	}
}

func CompileDeclStmt(stmt *ast.DeclStmt, r *RunNode) {
	switch stmt.Decl.(type) {
	case *ast.GenDecl:
		CompileGenDecl(stmt.Decl.(*ast.GenDecl), r)
	default:
		Error("o no = %+v, type =%+v \n", stmt.Decl, reflect.TypeOf(stmt.Decl))
	}
}

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

func CompileGenDecl(d *ast.GenDecl, r *RunNode) {
	for _, spe := range d.Specs {
		switch spe := spe.(type) {
		case *ast.ImportSpec:
			CompileImportSpec(spe, r)
		case *ast.ValueSpec:
			CompileVarSpec(spe, r)
		default:
			Error("o no")
		}
	}
}

func CompileImportSpec(spe *ast.ImportSpec, r *RunNode) {
	//Debug(" =======>> %s\n", spe.Path.Value)
	//TODO
	// 1\ import so
	// 2\ import r

}

func CompileFuncDecl(d *ast.FuncDecl, r *RunNode) {
	for _, stmt := range d.Body.List {
		switch stmt := stmt.(type) {
		case *ast.AssignStmt:
			CompileAssignStmt(stmt, r)
		case *ast.DeclStmt:
			CompileDeclStmt(stmt, r)
		case *ast.ExprStmt:
			CompileExpr(stmt.X, r)
		case *ast.RangeStmt:
			Debug("range-> 0 : %#v, 1 : %#v, 2 : %#v, 3: %#v", stmt.Key, stmt.Value, stmt.X, stmt.Body.List)
			CompileRangeStmt(stmt, r)
		default:
			Error("undefind value -> %#v", stmt)
		}
	}
}

func CompileRangeStmt(stmt *ast.RangeStmt, r *RunNode) {
	xs := CompileExpr(stmt.X, r)
	switch xs := xs.(type) {
	case map[interface{}]interface{}:
		for k, v := range xs {
			rs := &RunNode{
				Father: r,
				VarMap: make(map[string]interface{}, 0),
			}
			rs.VarMap[stmt.Key.(*ast.Ident).Name] = k
			rs.VarMap[stmt.Value.(*ast.Ident).Name] = v
			f := &ast.FuncDecl{
				Body: stmt.Body,
			}
			CompileFuncDecl(f, rs)
		}
	default:
		Error("o dont")

	}
}
