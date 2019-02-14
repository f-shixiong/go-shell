package lib

import (
	"go/ast"
)

func CompileMap(x *ast.MapType, r *RunNode) interface{} {
	//biebiwo
	switch x.Key.(*ast.Ident).Name {
	case "string":

	}

	return nil
}
