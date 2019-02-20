package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	var f string
	if len(os.Args) < 2 {
		os.Exit(0)
	} else {
		f = os.Args[1]
	}
	fmt.Printf("f %+v\n", f)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", f, 0)
	if err != nil {
		fmt.Println(err)
	}

	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			fmt.Printf("%#v\n", decl)
		case *ast.GenDecl:
			fmt.Printf("%#v\n", decl)
		default:
			fmt.Println("superrise")

		}

	}

}
