package lib

import (
	"fmt"
	"go/ast"
)

func GetRty(tp ast.Expr) string {
	var rty string
	switch ft := tp.(type) {
	case *ast.Ident:
		rty = ft.Name
	case *ast.ArrayType:
		elt := ft.Elt.(*ast.Ident).Name
		l := ""
		if ft.Len != nil {
			l = fmt.Sprintf("%v", ft.Len)
		}
		rty = fmt.Sprintf("[%s]%s", l, elt)
	case *ast.StarExpr:
		switch x := ft.X.(type) {
		case *ast.Ident:
			rty = "*" + x.Name
		default:
			fmt.Printf("sp -2 %#v\n", x)

		}
	case *ast.SelectorExpr:
		rty = ft.X.(*ast.Ident).Name + "." + ft.Sel.Name
	case *ast.Ellipsis:
		rty = fmt.Sprintf("...%v", GetRty(ft.Elt))
	case *ast.FuncType:
		rty = getFuncLit(convertFuncType(ft))
		fmt.Printf("func = %#v\n", "todo")
	default:
		fmt.Printf("sp -1 %#v\n", ft)

	}
	return rty
}
