package lib

import (
	"go/ast"
)

type Var struct {
	k interface{}
	v ast.Expr
}

func (Var) String() string {
	return "" //TODO
}

type RunNode struct {
	Father  *RunNode
	Vars    []Var
	VarMap  map[string]interface{}
	childs  []*RunNode //TODO set father
	FuncMap map[string]*ast.FuncDecl
}

func (r *RunNode) Get(scopes []string) Var {

	return Var{}
}
