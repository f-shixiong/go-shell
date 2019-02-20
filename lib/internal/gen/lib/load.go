package lib

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func Load(file *ast.File, funcMap map[string]FC, genMap map[string]GEN) {
	for _, decl := range file.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			name := decl.Name.Name
			belongStr := ""
			if decl.Recv != nil {
				for _, r := range decl.Recv.List {
					rty := GetRty(r.Type)
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
					rty := GetRty(fie.Type)
					paramStr += fmt.Sprintf("%s %s", name, rty)
					reqArr = append(reqArr, name)
				}
			}
			//withK := false
			rsarr := make([]string, 0)
			rlist := make([]string, 0)
			if tp.Results != nil {
				for i, fie := range tp.Results.List {
					name := fmt.Sprintf("%v", fie.Names)
					name = name[1 : len(name)-1]
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
			funcMap[name] = FC{
				Name:   name,
				ReqStr: strings.Join(reqArr, ","),
				RetStr: retStr,
				Rev:    belongStr,
				Params: paramStr,
				Ret:    resultStr,
			}
		case *ast.GenDecl:
			gen := GEN{}
			switch decl.Tok {
			case token.IMPORT:
				gen.L1 = "import"
				iarr := make([]string, 0)
				for _, sp := range decl.Specs {
					iarr = append(iarr, sp.(*ast.ImportSpec).Path.Value)
				}
				gen.Name = strings.Join(iarr, "\n    	")
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

		default:
			fmt.Println("superrise")

		}
	}

}
