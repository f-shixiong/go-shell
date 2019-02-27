package main

import (
	"fmt"
	"github.com/f-shixiong/go-shell/lib/internal/gen/lib"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

var testfile = "/home/shizuo/disk/go/src/strings/builder.go"
var testDir = "/home/shizuo/disk/go/src/math"

func main() {
	var dir string
	if len(os.Args) < 2 {
		dir = testDir
	} else {
		dir = os.Args[1]
	}
	var (
		funcMap = make(map[string]lib.FC, 0)
		genMap  = make(map[string]lib.GEN, 0)
		fset    = token.NewFileSet()
		pack    = strings.Split(dir, "/")[len(strings.Split(dir, "/"))-1]
	)
	fmt.Printf("init package = %+v\n", pack)
	fs, err := ioutil.ReadDir(testDir)
	if err != nil {
		panic(err)
	}
	for _, fModel := range fs {
		f := dir + "/" + fModel.Name()
		if strings.Contains(f, "_test.go") || f[len(f)-3:len(f)] != ".go" {
			continue
		}
		fmt.Println(f)
		file := getFile(fset, f)
		lib.LoadGen(file, genMap)
	}
	for _, fModel := range fs {
		f := dir + "/" + fModel.Name()
		if strings.Contains(f, "_test.go") || f[len(f)-3:len(f)] != ".go" {
			continue
		}
		fmt.Println(f)
		file := getFile(fset, f)
		lib.LoadFunc(file, funcMap, genMap, pack)
	}

	lib.Out(pack, funcMap, genMap)
	//ast.Print(fset, file)
}

func getFile(fset *token.FileSet, fileName string) *ast.File {
	data, err := ioutil.ReadFile(fileName)
	arr := strings.Split(string(data), "\n")
	begin := false
	cparr := make([]string, 0)
	for _, line := range arr {
		if len(line) > 7 && !begin {
			if line[0:7] == "package" {
				begin = true
			}
		}
		if begin {
			cparr = append(cparr, line)
		}
	}
	src := strings.Join(cparr, "\n")
	file, err := parser.ParseFile(fset, "", []byte(src), 0)
	if err != nil {
		fmt.Println(err)
	}
	return file
}
