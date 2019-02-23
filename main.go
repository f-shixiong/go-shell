package main

import (
	"flag"
	"fmt"
	"github.com/f-shixiong/go-shell/lib"
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"github.com/f-shixiong/go-shell/lib/go/parser"
	"github.com/f-shixiong/go-shell/lib/go/token"
	"io/ioutil"
	"os"
	"strings"
)

const (
	ONLYONEFILE  = "default"
	CONTEXT      = "package main\n%s"
	CONTEXT_MAIN = "package main\nfunc main(){ \n %s \n}"
)

var runNode = &lib.RunNode{}

var (
	mRepl = flag.Bool("r", false, "repl mode")
	mFile = flag.String("f", "", "file mode")
	mText = flag.String("t", "", "text mode")
)

func main() {
	fmt.Printf("-----goshell 1.0 by dy.com------\n\n")
	flag.Parse()
	var (
		data []byte
		err  error
		fset = token.NewFileSet()
	)

	if *mFile != "" {
		data, err = ioutil.ReadFile(*mFile)
		if err != nil {
			panic(err)
		}
	} else if *mText != "" {
		data = []byte(*mText)
	} else {
		ReplMode(fset, nil, runNode)
	}
	DataMode(fset, data, runNode)
}

func ReplMode(fset *token.FileSet, data []byte, r *lib.RunNode) {
	var (
		status = 0
		let    = ""
		let2   = ""
		err    error
		files  []*ast.File
	)
	_ = status
	_ = err
	_ = files
	for {
		fmt.Scanln(&let, &let2)
		fmt.Printf("let = %s, let2 = %s\n", let, let2)
		if let == "exit" || let2 == "exit" {

			os.Exit(0)
		}

	}
	//ExprFile(f, runNode)
}

func DataMode(fset *token.FileSet, data []byte, r *lib.RunNode) {
	var (
		status = 0
		let    = ""
		err    error
		files  []*ast.File
	)
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
		fmt.Printf("compile fail > \n%s,err=%+v \n", let, err)
		os.Exit(0)
	}
	for _, f := range files {
		ExprFile(f, runNode)
	}
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
