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
		switch elt := ft.Elt.(type) {
		case *ast.Ident:
			l := ""
			if ft.Len != nil {
				l = fmt.Sprintf("%v", ft.Len)
			}
			rty = fmt.Sprintf("[%s]%s", l, elt.Name)
		default:
			rty = "undefine"
			fmt.Printf("sp -1 %#v\n", ft)
		}
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
	default:
		fmt.Printf("sp -1 %#v\n", ft)

	}
	return rty
}
