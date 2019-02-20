package main

import (
	"fmt"
	"github.com/f-shixiong/go-shell/lib/internal/gen/lib"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

var testfile = "/home/shizuo/disk/go/src/strings/builder.go"

func main() {
	var f string
	if testfile != "" {
		f = testfile
	} else if len(os.Args) < 2 {
		os.Exit(0)
	} else {
		f = os.Args[1]
	}
	fmt.Printf("f = %+v\n", f)
	data, err := ioutil.ReadFile(f)
	arr := strings.Split(string(data), "\n")
	begin := false
	cparr := make([]string, 0)
	pack := ""
	for _, line := range arr {
		if len(line) > 7 && !begin {
			if line[0:7] == "package" {
				parr := strings.Split(line, " ")
				pack = parr[1]
				begin = true
			}
		}
		if begin {
			cparr = append(cparr, line)
		}
	}
	src := strings.Join(cparr, "\n")
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "", []byte(src), 0)
	if err != nil {
		fmt.Println(err)
	}

	funcMap := make(map[string]lib.FC, 0)
	genMap := make(map[string]lib.GEN, 0)

	lib.Load(file, funcMap, genMap)
	lib.Out(pack, funcMap, genMap)
	//ast.Print(fset, file)
}

func getName() {

}
