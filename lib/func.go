package lib

import (
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"github.com/f-shixiong/go-shell/lib/go/token"
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
					Error("T = %#v", T)
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
					Error("T = %#v", T)
				}
			default:
				Error("tp = %#v", tp)
			}
		}

	}
}

func CompileIncDec(stmt *ast.IncDecStmt, r *RunNode) {
	k := stmt.X.(*ast.Ident).Name
	v := r.GetValue(k)
	Debug("stmt = %#v", stmt)
	tok := stmt.Tok
	switch v := v.(type) {
	case int:
		if tok == token.INC {
			v++
		}
		if tok == token.DEC {
			v--
		}
		r.SetValue(k, v)
	case int64:
		if tok == token.INC {
			v++
		}
		if tok == token.DEC {
			v--
		}
		r.SetValue(k, v)
	case float32:
		if tok == token.INC {
			v++
		}
		if tok == token.DEC {
			v--
		}
		r.SetValue(k, v)
	case float64:
		if tok == token.INC {
			v++
		}
		if tok == token.DEC {
			v--
		}
		r.SetValue(k, v)
	default:
		Error("undefined type %#v", v)
	}
}

func CompileFuncDecl(d *ast.FuncDecl, r *RunNode) (ret []interface{}) {
	for _, stmt := range d.Body.List {
		switch stmt := stmt.(type) {
		case *ast.ReturnStmt:
			for _, result := range stmt.Results {
				ret = append(ret, CompileExpr(result, r))
			}
		default:
			hasR, ret := CompileStmt(stmt, r)
			if hasR {
				return ret
			}
		}
	}
	return
}
