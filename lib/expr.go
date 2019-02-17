package lib

import (
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"github.com/spf13/cast"
	"go/token"
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
			Error("compile_expr -> %#v , ret -> %+v ", x, ret)
		}
	}()
	switch x := x.(type) {
	case *ast.BasicLit:
		switch x.Kind {
		case token.INT:
			ret = cast.ToInt(x.Value)
		case token.FLOAT:
			//TODO
		case token.STRING:
			ret = x.Value
		default:
			Error("there is should happend_1 %#v \n", x)
		}
	case *ast.CompositeLit:
		tp := r.GetType(x.Type.(*ast.Ident).Name)
		if tp == nil {
			Error("con't find type %s", x.Type.(*ast.Ident).Name)
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
					Error("tttttttt ")
				}
			}
			return stu
		case Struct:
			for _, el := range x.Elts {
				switch el := el.(type) {
				case *ast.KeyValueExpr:
					tp.vals[el.Key.(*ast.Ident).Name] = CompileExpr(el.Value, r)
				default:
					Error("tttttttt ")
				}
			}
			return tp
		default:
			Error("o no  unsupport %#v", tp)
		}
	case *ast.FuncLit:
		//TODO
		Error("there is should happend_3 %#v \n", x)
	case *ast.ArrayType:
		//TODO
		return make([]interface{}, 0)
	case *ast.ChanType:
		Error("there is should happend_4.1 %#v \n", x)
	case *ast.Ellipsis:
		Error("there is should happend_4.2 %#v \n", x)
	case *ast.FuncType:
		Error("there is should happend_4.3 %#v \n", x)
	case *ast.InterfaceType:
		Error("there is should happend_4.4 %#v \n", x)
	case *ast.MapType:
		return make(map[interface{}]interface{}, 0)
	case *ast.BadExpr:
		//TODO
		Error("there is should happend_5 %#v \n", x)
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
		default:
			Error("there is should not happd %#v", f)
		}

	case *ast.Ident:
		noret = true
		if _, ok := baseType[x.Name]; ok {
			ret = x.Name
		} else {
			ret = r.GetValue(x.Name)
		}
	case *ast.IndexExpr:
		l := CompileExpr(x.X, r)
		switch l := l.(type) {
		case map[interface{}]interface{}:
			ret = l[CompileExpr(x.Index, r)]
		case []interface{}:
			ret = l[cast.ToInt(CompileExpr(x.Index, r))]
		default:
			Error("o it dont  ha %#v", l)
		}
	case *ast.SliceExpr:
		//TODO
		Error("there is should happend_8 %#v \n", x)
	case *ast.KeyValueExpr:
		//TODO
		Error("there is should happend_9 %#v \n", x)
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
			Error("but int happend_12 %#v , r = %#v ,rret = %#v \n", x.X, rret)
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
		Error("there is should happend_14 %#v \n", x)
	case *ast.UnaryExpr:
		//TODO
		Error("there is should happend_15 %#v \n", x)
	case nil:
		noret = true
	default:
		//TODO
		Error("there is should happend_16 %#v \n", x)
	}
	return
}
