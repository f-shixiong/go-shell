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

	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			name := decl.Name.Name
			belongStr := ""
			if decl.Recv != nil {
				for _, r := range decl.Recv.List {
					rty := lib.GetRty(r.Type)
					name := fmt.Sprintf("%v", r.Names)
					name = name[1 : len(name)-1]
					belongStr += fmt.Sprintf("%s %s", name, rty)
				}
			}
			if belongStr != "" {
				belongStr = fmt.Sprintf("(%s)", belongStr)
			}
			tp := decl.Type
			paramStr := ""
			resultStr := ""
			reqArr := make([]string, 0)
			if tp.Params != nil {
				for _, fie := range tp.Params.List {
					name := fmt.Sprintf("%v", fie.Names)
					name = name[1 : len(name)-1]
					rty := lib.GetRty(fie.Type)
					paramStr += fmt.Sprintf("%s %s", name, rty)
					reqArr = append(reqArr, name)
				}
			}
			withK := false
			rsarr := make([]string, 0)
			rlist := make([]string, 0)
			if tp.Results != nil {
				for i, fie := range tp.Results.List {
					name := fmt.Sprintf("%v", fie.Names)
					name = name[1 : len(name)-1]
					rty := lib.GetRty(fie.Type)
					rsarr = append(rsarr, rty)
					rlist = append(rlist, fmt.Sprintf("r_%v", i))
				}
			}
			resultStr = strings.Join(rsarr, ",")
			if withK {
				resultStr = fmt.Sprintf("(%s)", resultStr)
			}
			retStr := strings.Join(rlist, ",")
			funcMap[name] = lib.FC{
				Name:   name,
				ReqStr: strings.Join(reqArr, ","),
				RetStr: retStr,
				Rev:    belongStr,
				Params: paramStr,
				Ret:    resultStr,
				Pack:   pack,
			}
			lib.OutFunc(funcMap[name])
		case *ast.GenDecl:
			gen := lib.GEN{}
			switch decl.Tok {
			case token.IMPORT:
				gen.L1 = "import"
				iarr := make([]string, 0)
				for _, sp := range decl.Specs {
					iarr = append(iarr, sp.(*ast.ImportSpec).Path.Value)
				}
				gen.Name = strings.Join(iarr, "\n")
			case token.VAR:
				gen.L1 = "var"
				for _, sp := range decl.Specs {
					//gen.Name += sp.(*ast.ValueSpec)
					fmt.Printf("sp -3 %#v", sp)
				}
			case token.TYPE:
				gen.L1 = "type"
				for _, sp := range decl.Specs {
					gen.Name += sp.(*ast.TypeSpec).Name.Name
				}
			default:
				fmt.Println("superrise0")
			}
			genMap[gen.Name] = gen
			lib.OutGen(genMap[gen.Name])

		default:
			fmt.Println("superrise")

		}
	}
	//ast.Print(fset, file)
}

func getName() {

}
