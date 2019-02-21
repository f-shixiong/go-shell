package lib

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func LoadFunc(file *ast.File, funcMap map[string]FC, genMap map[string]GEN, pack string) {
	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			fc := getFunc(decl, genMap, pack)
			funcMap[fc.Name] = fc
			//fmt.Println(funcMap[name])
		case *ast.GenDecl:
		default:
			fmt.Println("superrise")

		}
	}

}

func LoadGen(file *ast.File, genMap map[string]GEN) {
	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
		case *ast.GenDecl:
			gen := getGen(decl)
			genMap[gen.Name] = gen

		default:
			fmt.Println("superrise")

		}
	}

}

func getFunc(decl *ast.FuncDecl, genMap map[string]GEN, pack string) FC {
	name := decl.Name.Name
	belongStr := ""
	if decl.Recv != nil {
		for _, r := range decl.Recv.List {
			rty := GetRty(r.Type)
			name := fmt.Sprintf("%v", r.Names)
			name = name[1 : len(name)-1]
			name = strings.Replace(name, " ", ",", 20)
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
	paramArr := make([]string, 0)
	if tp.Params != nil {
		for _, fie := range tp.Params.List {
			name := fmt.Sprintf("%v", fie.Names)
			name = name[1 : len(name)-1]
			name = strings.Replace(name, " ", ",", 20)
			rty := GetRty(fie.Type)
			paramArr = append(paramArr, fmt.Sprintf("%s %s", name, rty))
			reqArr = append(reqArr, name)
		}
	}
	paramStr = strings.Join(paramArr, ",")
	rsarr := make([]string, 0)
	rlist := make([]string, 0)
	if tp.Results != nil {
		for i, fie := range tp.Results.List {
			name := fmt.Sprintf("%v", fie.Names)
			name = name[1 : len(name)-1]
			name = strings.Replace(name, " ", ",", 20)
			rty := GetRty(fie.Type)
			rty = coverRty(pack, genMap, rty)
			rsarr = append(rsarr, rty)
			rlist = append(rlist, fmt.Sprintf("r_%v", i))
		}
	}
	resultStr = strings.Join(rsarr, ",")
	//if withK {
	resultStr = fmt.Sprintf("(%s)", resultStr)
	//}
	retStr := strings.Join(rlist, ",")
	return FC{
		Name:   name,
		ReqStr: strings.Join(reqArr, ","),
		RetStr: retStr,
		Rev:    belongStr,
		Params: paramStr,
		Ret:    resultStr,
	}

}

func convertFuncType(tp *ast.FuncType) FC {
	paramStr := ""
	resultStr := ""
	reqArr := make([]string, 0)
	paramArr := make([]string, 0)
	if tp.Params != nil {
		for _, fie := range tp.Params.List {
			name := fmt.Sprintf("%v", fie.Names)
			name = name[1 : len(name)-1]
			name = strings.Replace(name, " ", ",", 20)
			rty := GetRty(fie.Type)
			paramArr = append(paramArr, fmt.Sprintf("%s %s", name, rty))
			reqArr = append(reqArr, name)
		}
	}
	paramStr = strings.Join(paramArr, ",")
	rsarr := make([]string, 0)
	rlist := make([]string, 0)
	if tp.Results != nil {
		for i, fie := range tp.Results.List {
			name := fmt.Sprintf("%v", fie.Names)
			name = name[1 : len(name)-1]
			name = strings.Replace(name, " ", ",", 20)
			rty := GetRty(fie.Type)
			rsarr = append(rsarr, rty)
			rlist = append(rlist, fmt.Sprintf("r_%v", i))
		}
	}
	resultStr = strings.Join(rsarr, ",")
	//if withK {
	resultStr = fmt.Sprintf("(%s)", resultStr)
	//}
	retStr := strings.Join(rlist, ",")
	return FC{
		ReqStr: strings.Join(reqArr, ","),
		RetStr: retStr,
		Params: paramStr,
		Ret:    resultStr,
	}
}

func getGen(decl *ast.GenDecl) GEN {
	gen := GEN{}
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
			vp := sp.(*ast.ValueSpec)
			name := fmt.Sprintf("%v", vp.Names)
			if skiPrivate(name) || name == "" {
				continue
			} else {
				fmt.Printf("sp -5 %#v\n", vp)
			}
			gen.Name += name

		}
	case token.TYPE:
		gen.L1 = "type"
		for _, sp := range decl.Specs {
			gen.Name += sp.(*ast.TypeSpec).Name.Name
		}
	case token.CONST:
		gen.L1 = "const"
		for _, sp := range decl.Specs {
			//gen.Name += sp.(*ast.ValueSpec)
			vp := sp.(*ast.ValueSpec)
			name := fmt.Sprintf("%v", vp.Names)
			if skiPrivate(name) {
				continue
			} else {
				fmt.Printf("sp -5 %#v\n", vp)
			}

		}
	default:
		fmt.Printf("superrise0, %+v\n", decl.Tok.String())
	}
	return gen
}
