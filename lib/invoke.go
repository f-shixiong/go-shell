package lib

import (
	"fmt"
	_ "go/ast"
	_ "reflect"
)

func Invock(e expr, r *RunNode) {
	fmt.Println("------------------------------------")
	switch e.method {
	case "println":
		fmt.Println(e.args)
	default:
		fmt.Println("how it happend")

	}
}

func InvockSo() {

}

func InvockCos() {

}
