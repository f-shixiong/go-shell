package lib

import (
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"github.com/f-shixiong/go-shell/lib/go/token"
	"github.com/spf13/cast"
)

type funcArg struct {
	node *RunNode
	dep  int
	val  string
}

type expr struct {
	args     []interface{}
	t        int
	method   string
	instance interface{}
	left     interface{}
}

func CompileExpr(x ast.Expr, r *RunNode) (ret interface{}) {
	noret := false
	defer func() {
		if ret != nil || noret {
			Debug("compile_expr -> %#v , ret -> %+v ", x, ret)
		} else {
			Error("compile_expr = %#v , ret = %+v ", x, ret)
		}
	}()
	switch x := x.(type) {
	case *ast.BasicLit:
		switch x.Kind {
		case token.INT:
			ret = cast.ToInt(x.Value)
		case token.FLOAT:
			ret = cast.ToFloat64(x.Value)
		case token.STRING:
			ret = x.Value
		case token.CHAR:
			if len(x.Value) == 0 {
				Error("empty char %#v", x.Value)
			} else {
				return x.Value[0]
			}

		default:
			Error("type = %#v \n", x)
		}
	case *ast.CompositeLit:
		tp := r.GetType(x.Type.(*ast.Ident).Name)
		if tp == nil {
			Error("type = %s", x.Type.(*ast.Ident).Name)
		}
		switch tp := tp.(type) {
		case *ast.StructType:
			stu := Struct{
				vals: make(map[interface{}]interface{}, 0),
			}
			for _, el := range x.Elts {
				switch el := el.(type) {
				case *ast.KeyValueExpr:
					stu.vals[el.Key.(*ast.Ident).Name] = CompileExpr(el.Value, r)
				default:
					Error(" type = %#v", el)
				}
			}
			return stu
		case Struct:
			for _, el := range x.Elts {
				switch el := el.(type) {
				case *ast.KeyValueExpr:
					tp.vals[el.Key.(*ast.Ident).Name] = CompileExpr(el.Value, r)
				default:
					Error(" type = %#v", el)
				}
			}
			return tp
		default:
			Error("tp =  %#v", tp)
		}
	case *ast.FuncLit:
		//TODO
		Error(" todo = %#v \n", x)
	case *ast.ArrayType:
		//TODO
		return make([]interface{}, 0)
	case *ast.ChanType:
		Error(" todo = %#v \n", x)
	case *ast.Ellipsis:
		Error(" todo = %#v \n", x)
	case *ast.FuncType:
		Error(" todo = %#v \n", x)
	case *ast.InterfaceType:
		Error(" todo = %#v \n", x)
	case *ast.MapType:
		return make(map[interface{}]interface{}, 0)
	case *ast.BadExpr:
		//TODO
		Error(" todo = %#v \n", x)
	case *ast.BinaryExpr:
		ret = CompileBinary(x, r)
	case *ast.CallExpr:
		noret = true
		switch f := x.Fun.(type) {
		case *ast.Ident:
			e := expr{
				method: f.Name,
				args:   CompileArgs(x.Args, r),

				//TODO t
			}
			ret = Invock(e, r)
		case *ast.SelectorExpr:
			Debug("x = %#v, sel = %#v", f.X, f.Sel)
			e := expr{
				method: f.Sel.Name,
				args:   CompileArgs(x.Args, r),
				left:   CompileExpr(f.X, r),
			}
			ret = Invock(e, r)
		case *ast.ArrayType:
			if elt, ok := f.Elt.(*ast.Ident); ok {
				if elt.Name == "byte" {
					return []byte(cast.ToString(CompileArgs(x.Args, r)[0]))
				}
			}
			Error(" f = %#v, elt = %#v", f, f.Elt)
		default:
			Error("todo = %#v", f)
		}

	case *ast.Ident:
		noret = true
		_, ok := baseType[x.Name]
		if ok {
			return x.Name
		}
		if x.Name == "true" {
			return true
		}

		if x.Name == "false" {
			return false
		}
		ret = r.GetValue(x.Name)
	case *ast.IndexExpr:
		l := CompileExpr(x.X, r)
		switch l := l.(type) {
		case map[interface{}]interface{}:
			ret = l[CompileExpr(x.Index, r)]
		case []interface{}:
			ret = l[cast.ToInt(CompileExpr(x.Index, r))]
		default:
			Error(" todo = %#v", l)
		}
	case *ast.SliceExpr:
		//TODO
		Error(" todo = %#v \n", x)
	case *ast.KeyValueExpr:
		//TODO
		Error(" todo = %#v \n", x)
	case *ast.ParenExpr:
		ret = CompileExpr(x.X, r)
	case *ast.SelectorExpr:
		ret = CompileSelectorExpr(x, r)
	case *ast.StarExpr:
		//TODO
		rret := r.GetValue(x.X.(*ast.Ident).Name)
		switch rret := rret.(type) {
		case *int:
			ret = *rret
		case *interface{}:
			ret = *rret
		default:
			Error("todo = %#v , r = %#v ,rret = %#v", x, x.X, rret)
		}

	case *ast.StructType:
		//TODO field
		return Struct{
			vals:     make(map[interface{}]interface{}, 0),
			funcMap:  make(map[string]*ast.FuncDecl, 0),
			funcRMap: make(map[string]*RunNode, 0),
		}

	case *ast.TypeAssertExpr:
		//TODO
		Error(" todo = %#v \n", x)
	case *ast.UnaryExpr:
		ret = CompileUnaryExpr(x, r)
	case nil:
		noret = true
	case *ast.ShellStmt:
		ret = InvockShell(x, r)
	default:
		//TODO
		Error(" todo = %#v \n", x)
	}
	return
}
