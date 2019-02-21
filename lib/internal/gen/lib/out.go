package lib

import (
	"fmt"
	"strings"
)

type FC struct {
	Pack   string
	Name   string
	Rev    string
	Params string
	Ret    string
	ReqStr string
	RetStr string
}

type GEN struct {
	Name string
	L1   string
}

var fmtter = `
func %s %s(%s) %s {
	%s := %s.%s(%s)
	return %s
}
`
var fmtter2 = `
func %s %s(%s) {
	%s.%s(%s)
}

`

var flit1 = `func(%s)`
var flit2 = `func(%s)%s`
var genFmtter = `
%s %s %s
`
var importFmtter = `
import (
        %s
        "%s"
)
`
var typeFmtter = `
type %s struct {
	inside %s
}
`

func getFuncLit(fc FC) string {
	if fc.Rev != "" {
		panic("sup")
	}
	if fc.RetStr == "" {
		return fmt.Sprintf(flit1, fc.Params)
	} else {
		return fmt.Sprintf(flit2, fc.Params, fc.Ret)
	}

}

func OutFunc(fc FC, pack string) {
	if fc.Rev != "" {
		return
	}
	if fc.RetStr == "" {
		fmt.Printf(fmtter2, fc.Rev, fc.Name, fc.Params, pack, fc.Name, fc.ReqStr)
	} else {
		fmt.Printf(fmtter, fc.Rev, fc.Name, fc.Params, fc.Ret, fc.RetStr, pack, fc.Name, fc.ReqStr, fc.RetStr)
	}
}

func OutGen(gen GEN, pack string) {
	switch gen.L1 {
	case "import":
		fmt.Printf(importFmtter, gen.Name, pack)
	//case "type":
	//	fmt.Printf(typeFmtter, gen.Name, pack+"."+gen.Name)
	default:
		fmt.Printf(genFmtter, gen.L1, gen.Name, pack+"."+gen.Name)
	}
}

func Out(pack string, funcMap map[string]FC, genMap map[string]GEN) {
	fmt.Printf("package %s\n", pack)
	imap := map[string]string{}
	for _, v := range genMap {
		if v.L1 != "import" {
			continue
		}
		iarr := strings.Split(v.Name, "\n")
		for _, i := range iarr {
			imap[i] = ""
		}
	}
	iGen := GEN{}
	iGen.L1 = "import"
	iiarr := []string{}
	for k, _ := range imap {
		iiarr = append(iiarr, k)
	}
	iGen.Name = strings.Join(iiarr, "\n	")
	OutGen(iGen, pack)

	for k, v := range genMap {
		if v.L1 == "import" || skiPrivate(k) || k == "" {
			continue
		}
		OutGen(v, pack)
	}

	for k, v := range funcMap {
		if skiPrivate(k) {
			continue
		}
		OutFunc(v, pack)
	}

}

func skiPrivate(k string) bool {
	if len(k) == 0 {
		return false
	}
	b := []byte(k[0:1])[0]
	if b <= 90 && b >= 65 {
		return false
	}
	return true

}
