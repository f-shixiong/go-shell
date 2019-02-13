package lib

import (
	"fmt"
	"go/ast"
	"go/token"
)

func CompileArgs(args []ast.Expr, r *RunNode) (retArgs []interface{}) {
	for _, a := range args {
		retArgs = append(retArgs, CompileExpr(a, r))
	}
	return
}

func CompileBinary(x *ast.BinaryExpr, r *RunNode) (ret interface{}) {
	switch x.Op {
	case token.ADD:
		return BAdd(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	default:
		fmt.Println("o it really happend ?")
	}
	return
}
