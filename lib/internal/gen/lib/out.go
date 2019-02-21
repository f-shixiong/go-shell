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
	T    string
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
var tmapFmtter = `
var StuMap = map[string]interface{}{
	%s
}
`

var fmapFmtter = `
var FucMap = map[string]interface{}{
	%s
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

func OutType(tmap map[string]GEN, pack string) {
	tpArr := []string{}
	for k, v := range tmap {
		if skiPrivate(k) {
			continue
		}
		if v.T != "" {
			continue
		}
		tpArr = append(tpArr, fmt.Sprintf(`"%s":%s`, k, pack+"."+k+"{},\n	"))
	}
	fmt.Printf(tmapFmtter, strings.Join(tpArr, ""))
}

func OutFuc2(fArr []string, pack string) {
	cfArr := []string{}
	for _, k := range fArr {
		if skiPrivate(k) {
			continue
		}
		cfArr = append(cfArr, fmt.Sprintf(`"%s":%s`, k, pack+"."+k+",\n        "))
	}
	fmt.Printf(fmapFmtter, strings.Join(cfArr, ""))

}

func Out(pack string, funcMap map[string]FC, genMap map[string]GEN) {
	fmt.Printf("package %s\n", pack)
	fmt.Printf(`import "%s"`, pack)
	imap := map[string]string{}
	tmap := map[string]GEN{}
	for _, v := range genMap {
		if v.L1 == "type" {
			tmap[v.Name] = v
			continue
		}
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
	_ = iGen
	//OutGen(iGen, pack)
	OutType(tmap, pack)
	for k, v := range genMap {
		if v.L1 == "import" || v.L1 == "type" || skiPrivate(k) || k == "" {
			continue
		}
		OutGen(v, pack)
	}
	fArr := []string{}
	for k, v := range funcMap {
		if skiPrivate(k) || v.Rev != "" {
			continue
		}
		fArr = append(fArr, k)
		_ = v
		//OutFunc(v, pack)
	}
	OutFuc2(fArr, pack)

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
