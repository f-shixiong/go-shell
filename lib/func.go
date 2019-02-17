package lib

import (
	"github.com/f-shixiong/go-shell/lib/go/ast"
)

func LoadFunc(f *ast.FuncDecl, r *RunNode) {
	Debug("load func = %#v, recv = %#v", f, f.Recv)
	r.FuncMap[f.Name.String()] = f
	if f.Recv != nil {
		Debug("recv  = %#v", f.Recv)
		for _, rec := range f.Recv.List {
			switch tp := rec.Type.(type) {
			case *ast.Ident:
				rc := r.Child()
				T := r.GetType(tp.Name)
				for _, n := range rec.Names {
					rc.VarMap[n.Name] = SELF
				}
				switch T := T.(type) {
				case Struct:
					T.funcMap[f.Name.Name] = f
					T.funcRMap[f.Name.Name] = rc
					r.TypeMap[tp.Name] = T
				default:
					Error("bu ke neng,T = %#v", T)
				}
			case *ast.StarExpr:
				//TODO
				rc := r.Child()
				Debug("%#v", tp)
				T := r.GetType(tp.X.(*ast.Ident).Name)
				Debug("t = %#v,x.name=%#v,", T, tp.X.(*ast.Ident).Name)
				for _, n := range rec.Names {
					rc.VarMap[n.Name] = SELF
				}
				switch T := T.(type) {
				case Struct:
					T.funcMap[f.Name.Name] = f
					T.funcRMap[f.Name.Name] = rc
					r.TypeMap[tp.X.(*ast.Ident).Name] = T
					Debug("T = %#v,k = %#v", T.funcRMap["SetName"], tp.X.(*ast.Ident).Name)
				default:
					Error("bu ke neng,T = %#v", T)
				}
			default:
				Error("bu ke neng ")
			}
		}

	}
}

func CompileFuncDecl(d *ast.FuncDecl, r *RunNode) (ret []interface{}) {
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
		case *ast.ReturnStmt:
			for _, result := range stmt.Results {
				ret = append(ret, CompileExpr(result, r))
			}
		case *ast.ShellStmt:
			CompileShell(stmt, r)
		case *ast.EmptyStmt:
		default:
			Error("undefind value -> %#v", stmt)
		}
	}
	return
}
