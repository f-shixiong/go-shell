package tools

import (
	"fmt"
	"git.xiaojukeji.com/fanshixiong/nodejson"
	"github.com/f-shixiong/go-shell/lib"
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"github.com/f-shixiong/go-shell/lib/go/parser"
	"github.com/f-shixiong/go-shell/lib/go/token"
	"strings"
)

const (
	ONLYONEFILE  = "default"
	CONTEXT      = "package main\n%s"
	CONTEXT_MAIN = "package main\nfunc main(){ \n %s \n}"
)

var runNode = &lib.RunNode{}

type Lambda struct {
	runNode *lib.RunNode
	Lambda  string
}

func (l Lambda) ExecDiDi(request nodejson.Node) (response nodejson.Node, err error) {
	var (
		fset = token.NewFileSet()
	)
	if l.runNode == nil {
		l.runNode = &lib.RunNode{}
	}
	l.runNode.SetValue("req", request)
	l.runNode.SetValue("resp", &response)
	l.runNode.SetValue("err", &err)

	LambdaMode(fset, l.Lambda, l.runNode)
	return
}

func (l Lambda) Exec() {
	var (
		fset = token.NewFileSet()
	)
	if l.runNode == nil {
		l.runNode = &lib.RunNode{}
	}

	LambdaMode(fset, l.Lambda, l.runNode)
}

func LambdaMode(fset *token.FileSet, lambda string, r *lib.RunNode) {
	files, err := getFiles(fset, lambda, r)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		ExprFile(f, runNode)
	}

}

func getFiles(fset *token.FileSet, lambda string, r *lib.RunNode) ([]*ast.File, error) {
	var (
		status = 0
		let    = ""
		err    error
		files  []*ast.File
	)
	for _, l := range strings.Split(lambda, "\n") {
		if l == "" {
			continue
		}
		let += (l + "\n")
		file, e1 := parser.ParseFile(fset, "", fmt.Sprintf(CONTEXT, let), 0)
		if e1 != nil {
			err = e1
			status = -1
			file, e2 := parser.ParseFile(fset, "", fmt.Sprintf(CONTEXT_MAIN, let), 0)
			if e2 != nil {
				err = e2
				continue
			}
			lib.Debug(let)
			let = ""
			status = 0
			files = append(files, file)

		} else {
			lib.Debug(let)
			files = append(files, file)
			let = ""
			status = 0
		}
	}
	if status != 0 {
		return nil, fmt.Errorf("compile fail > \n%s,err=%+v \n", let, err)
	}
	return files, nil
}

func ExprFile(file *ast.File, r *lib.RunNode) {
	if len(file.Decls) == 0 {
		return
	}
	d := file.Decls[0]
	switch d := d.(type) {
	case *ast.GenDecl:
		lib.CompileGenDecl(d, runNode)
	case *ast.FuncDecl:
		if d.Name.String() == "main" {
			lib.CompileFuncDecl(d, runNode)
		} else {
			if runNode.FuncMap == nil {
				runNode.FuncMap = make(map[string]*ast.FuncDecl, 0)
			}
			lib.LoadFunc(d, runNode)
		}
	default:
		lib.Error("d = %#v", d)
	}
}
