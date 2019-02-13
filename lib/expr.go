package lib

import (
	"fmt"
	"go/ast"
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
}

func CompileExpr(x ast.Expr, r *RunNode) interface{} {
	fmt.Printf("compile_expr -> %#v \n\n", x)
	switch x := x.(type) {
	case *ast.BasicLit:
		switch x.Kind {
		case token.INT:
			//TODO
		case token.FLOAT:
			//TODO
		case token.STRING:
			return x.Value
		default:
			fmt.Printf("there is should happend_1 %#v \n", x)
		}
	case *ast.CompositeLit:
		//TODO
		fmt.Printf("there is should happend_2 %#v \n", x)
	case *ast.FuncLit:
		//TODO
		fmt.Printf("there is should happend_3 %#v \n", x)
	case *ast.ArrayType:
		//TODO
		fmt.Printf("there is should happend_4 %#v \n", x)
	case *ast.ChanType:
		fmt.Printf("there is should happend_4.1 %#v \n", x)
	case *ast.Ellipsis:
		fmt.Printf("there is should happend_4.2 %#v \n", x)
	case *ast.FuncType:
		fmt.Printf("there is should happend_4.3 %#v \n", x)
	case *ast.InterfaceType:
		fmt.Printf("there is should happend_4.4 %#v \n", x)
	case *ast.MapType:
		fmt.Printf("there is should happend_4.5 %#v \n", x)
	case *ast.BadExpr:
		//TODO
		fmt.Printf("there is should happend_5 %#v \n", x)
	case *ast.BinaryExpr:
		return CompileBinary(x, r)
	case *ast.CallExpr:
		//now TODO
		switch f := x.Fun.(type) {
		case *ast.Ident:
			e := expr{
				method: f.Name,
				args:   CompileArgs(x.Args, r),
				//TODO t
			}
			Invock(e, r)
		default:
			fmt.Println("there is should not happd")
		}

	case *ast.Ident:
		return r.VarMap[x.Name]
	case *ast.IndexExpr:
		//TODO
		fmt.Printf("there is should happend_7 %#v \n", x)
	case *ast.SliceExpr:
		//TODO
		fmt.Printf("there is should happend_8 %#v \n", x)
	case *ast.KeyValueExpr:
		//TODO
		fmt.Printf("there is should happend_9 %#v \n", x)
	case *ast.ParenExpr:
		//TODO
		fmt.Printf("there is should happend_10 %#v \n", x)
	case *ast.SelectorExpr:
		//TODO
		fmt.Printf("there is should happend_11 %#v \n", x)
	case *ast.StarExpr:
		//TODO
		fmt.Printf("there is should happend_12 %#v \n", x)
	case *ast.StructType:
		//TODO
		fmt.Printf("there is should happend_13 %#v \n", x)
	case *ast.TypeAssertExpr:
		//TODO
		fmt.Printf("there is should happend_14 %#v \n", x)
	case *ast.UnaryExpr:
		//TODO
		fmt.Printf("there is should happend_15 %#v \n", x)
	default:
		//TODO
		fmt.Printf("there is should happend_16 %#v \n", x)
	}
	return nil
}
