package lib

import (
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"plugin"
)

type Var struct {
	k interface{}
	v ast.Expr
}

type Struct struct {
	vals     map[interface{}]interface{}
	funcMap  map[string]*ast.FuncDecl
	funcRMap map[string]*RunNode
	Name     string
}

func (Var) String() string {
	return "" //TODO
}

type RunNode struct {
	Father *RunNode
	//Vars      []Var
	VarMap    map[string]interface{}
	TypeMap   map[string]interface{}
	childs    []*RunNode //TODO set father
	FuncMap   map[string]*ast.FuncDecl
	ImportMap map[string]plugin.Symbol
}

func (r *RunNode) Child() *RunNode {
	return &RunNode{
		Father: r,
		VarMap: make(map[string]interface{}, 0),
	}
}

func (r *RunNode) Get(scopes []string) Var {

	return Var{}
}

func (r *RunNode) GetValue(k string) interface{} {
	if val, ok := r.VarMap[k]; ok {
		return val
	} else if r.Father != nil {
		return r.Father.GetValue(k)
	}
	return r.ImportMap[k]
}

func (r *RunNode) SetValue(k string, v interface{}) {
	if _, ok := r.VarMap[k]; ok {
		r.VarMap[k] = v
	} else if r.Father != nil {
		r.Father.SetValue(k, v)
	} else {
		Error("what k =%v,v=%v", k, v)
	}

}

func (r *RunNode) GetType(k string) interface{} {
	if tp, ok := r.TypeMap[k]; ok {
		return tp
	}
	return r.Father.TypeMap[k]
}

func (r *RunNode) GetFunc(k string) *ast.FuncDecl {
	//TODO father
	return r.FuncMap[k]
}
