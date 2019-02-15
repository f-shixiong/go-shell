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
	TypeMap map[string]ast.Expr
	childs  []*RunNode //TODO set father
	FuncMap map[string]*ast.FuncDecl
}

func (r *RunNode) Get(scopes []string) Var {

	return Var{}
}

func (r *RunNode) GetValue(k string) interface{} {
	//TODO father
	return r.VarMap[k]
}

func (r *RunNode) GetType(k string) ast.Expr {
	//TODO father
	return r.TypeMap[k]
}

func (r *RunNode) GetFunc(k string) *ast.FuncDecl {
	//TODO father
	return r.FuncMap[k]
}
