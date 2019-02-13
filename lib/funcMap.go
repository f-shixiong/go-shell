package lib

import (
	"go/ast"
)

var FuncMap = make(map[string]*ast.FuncDecl, 0)
