package lib

import (
	"fmt"
	"go/ast"
	"reflect"
)

func CompileAssignStmt(stmt *ast.AssignStmt, r *RunNode) {
	if r.VarMap == nil {
		r.VarMap = make(map[string]interface{}, 0)
	}
	for i, k := range stmt.Lhs {
		switch k.(type) {
		case *ast.Ident:
			if len(stmt.Rhs) >= i {
				//*ast.BasicLit
				r.VarMap[k.(*ast.Ident).Name] = CompileExpr(stmt.Rhs[i], r)
			} else {
				r.VarMap[k.(*ast.Ident).Name] = CompileExpr(stmt.Rhs[0], r)

			}
		default:
			fmt.Println("o no")
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
		Debug("==> name : %+v, value : %#v \n", n.Name, stmt.Values[i])
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
		default:
			Debug("undefind type ->  %+v , value -> %+v\n", reflect.TypeOf(stmt), stmt)
		}
	}
}
