package main

import (
	"flag"
	"fmt"
	"github.com/f-shixiong/go-shell/lib"
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"github.com/f-shixiong/go-shell/lib/go/parser"
	"github.com/f-shixiong/go-shell/lib/go/token"
	"github.com/peterh/liner"
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
	mRepl        = flag.Bool("r", false, "repl mode")
	mFile        = flag.String("f", "", "file mode")
	mText        = flag.String("t", "", "text mode")
	bFlushImport = flag.Bool("flush_import", false, "flush package")
)

func main() {
	fmt.Printf("-----go-shell version:beta ------\n\n")
	flag.Parse()

	var (
		data []byte
		err  error
		fset = token.NewFileSet()
	)
	if bFlushImport != nil && *bFlushImport {
		lib.FlushImport = true
	}
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
	lib.Mode = 1
	var (
		code  = ""
		step  = 0
		err   error
		files []*ast.File
		lineR = liner.NewLiner()
		begin = true
	)
	defer lineR.Close()
	for {
		if begin {
			fmt.Printf("\033[32m%s [%d]↘\033[0m\n", "In", step)
		}
		let, _ := lineR.Prompt("")
		code += let
		if let == "exit" || let == ":q" || let == "quit" || let == "exit;" || let == "quit;" {
			lineR.Close()
			os.Exit(0)
		}
		if len(let) > 0 && let[len(let)-1] == ';' {
			files, err = getFiles(fset, []byte(code), r)
			if err != nil {
				fmt.Printf("\033[31m%s\033[0m\n", "-------------------------------------------------")
				fmt.Println(err)
			} else {
				fmt.Printf("\033[34m%s\033[0m\n", "-------------------------------------------------")
				for _, f := range files {
					ExprFile(f, r)
				}
			}
			step++
			code = ""
			begin = true
		} else {
			begin = false
		}
	}
	//ExprFile(f, runNode)
}

func DataMode(fset *token.FileSet, data []byte, r *lib.RunNode) {
	files, err := getFiles(fset, data, r)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		ExprFile(f, runNode)
	}

}

func getFiles(fset *token.FileSet, data []byte, r *lib.RunNode) ([]*ast.File, error) {
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
