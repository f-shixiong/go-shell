package lib

import (
	"fmt"
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
func %s %s (%s) %s {
	%s = %s(%s)
	return %s
}
`
var fmtter2 = `
func %s %s (%s) {
	%s(%s)
}

`
var genFmtter = `
%s %s %s
`

func OutFunc(fc FC) {
	if fc.RetStr == "" {
		fmt.Printf(fmtter2, fc.Rev, fc.Name, fc.Params, fc.Name, fc.ReqStr)
	} else {
		fmt.Printf(fmtter, fc.Rev, fc.Name, fc.Params, fc.Ret, fc.RetStr, fc.Name, fc.ReqStr, fc.RetStr)
	}
}

func OutGen(gen GEN) {
	switch gen.L1 {
	case "import":
		fmt.Printf(`
import (
	%s
)
`, gen.Name)
	default:
		fmt.Printf(genFmtter, gen.L1, gen.Name, gen.Name)
	}
}
