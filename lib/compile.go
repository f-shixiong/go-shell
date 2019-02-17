package lib

import (
	"fmt"
	"github.com/f-shixiong/go-shell/lib/go/ast"
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
			o := r.GetValue(k.X.(*ast.Ident).Name)
			if o == nil {
				Error("o no so sad ")
				return
			}
			mp := o.(map[interface{}]interface{})
			mp[CompileExpr(k.Index, r)] = CompileExpr(right, r)
			r.VarMap[k.X.(*ast.Ident).Name] = mp
			Debug("$%+v", CompileExpr(right, r))
		case *ast.StarExpr:
			Debug("test--> key = %#v,value = %#v", k.X, CompileExpr(right, r))
			realLeft := CompileExpr(right, r)
			r.VarMap[k.X.(*ast.Ident).Name] = &realLeft

		case *ast.SelectorExpr:
			Debug("x = %#v,sel = %#v,vars= %#v", k.X, k.Sel, r.VarMap)
			t := CompileExpr(k.X, r)
			switch t := t.(type) {
			case Struct:
				Debug("k = %#v, v = %#v, r = %#v", k.Sel.Name, CompileExpr(right, r), right)
				t.vals[k.Sel.Name] = CompileExpr(right, r)
			default:
				Error("should not %#v", t)
			}
			//Test("kk = %#v, vv = %#v", k.X.(*ast.Ident).Name, t)
			r.VarMap[k.X.(*ast.Ident).Name] = t
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

func CompileGenDecl(d *ast.GenDecl, r *RunNode) {
	for _, spe := range d.Specs {
		switch spe := spe.(type) {
		case *ast.ImportSpec:
			CompileImportSpec(spe, r)
		case *ast.ValueSpec:
			CompileVarSpec(spe, r)
		case *ast.TypeSpec:
			CompileTypeSpec(spe, r)
		default:
			Error("o no spe= %#v", spe)
		}
	}
}

func CompileRangeStmt(stmt *ast.RangeStmt, r *RunNode) {
	xs := CompileExpr(stmt.X, r)
	switch xs := xs.(type) {
	case map[interface{}]interface{}:
		for k, v := range xs {
			rc := r.Child()
			rc.VarMap[stmt.Key.(*ast.Ident).Name] = k
			rc.VarMap[stmt.Value.(*ast.Ident).Name] = v
			f := &ast.FuncDecl{
				Body: stmt.Body,
			}
			CompileFuncDecl(f, rc)
		}
	default:
		Error("o dont")

	}
}

func CompileShell(stmt *ast.ShellStmt, r *RunNode) {
	ret := InvockShell(stmt, r)
	if ret != "" {
		fmt.Println(ret)
	}
}
