package lib

import (
	"github.com/f-shixiong/go-shell/lib/go/token"
)

func AssignIdent(tok token.Token, left string, r *RunNode, right interface{}) {
	var newV interface{}
	switch tok {
	case token.DEFINE, token.ASSIGN:
		newV = right
	case token.ADD_ASSIGN:
		newV = BAdd(r.GetValue(left), right)
	case token.SUB_ASSIGN:
		newV = BSub(r.GetValue(left), right)
	case token.MUL_ASSIGN:
		newV = BMul(r.GetValue(left), right)
	case token.QUO_ASSIGN:
		newV = BQuo(r.GetValue(left), right)
	case token.REM_ASSIGN:
		newV = BRem(r.GetValue(left), right)
	case token.AND_ASSIGN:
		newV = BAnd(r.GetValue(left), right)
	case token.OR_ASSIGN:
		newV = BOr(r.GetValue(left), right)
	case token.XOR_ASSIGN:
		newV = BXor(r.GetValue(left), right)
	case token.SHL_ASSIGN:
		newV = BShl(r.GetValue(left), right)
	case token.SHR_ASSIGN:
		newV = BShr(r.GetValue(left), right)
	default:
		Error(" todo %#v", tok.String())
	}
	r.VarMap[left] = newV
}
