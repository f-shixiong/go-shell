package main

import (
	"flag"
	"fmt"
	"github.com/f-shixiong/go-shell/lib"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

//var file = flag.String("p", "", "import path")
const (
	ONLYONEFILE  = "default"
	CONTEXT      = "package main\n%s"
	CONTEXT_MAIN = "package main\nfunc main(){ \n %s \n}"
)

var runNode = &lib.RunNode{}

func main() {
	fmt.Printf("-----goshell 1.0 by dy.com------\n\n")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		os.Exit(0)
	}
	data, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	let := ""
	fset := token.NewFileSet()
	files := make([]*ast.File, 0)
	status := 0
	for _, l := range strings.Split(string(data), "\n") {
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
		fmt.Printf("paser not work lt = \n%s,err=%+v \n", let, err)
		os.Exit(0)
	}
	for _, f := range files {
		if len(f.Decls) == 0 {
			continue
		}
		d := f.Decls[0]
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
				runNode.FuncMap[d.Name.String()] = d
			}
		default:
			lib.Error("no----0v0----no")
		}
	}

}
