package lib

import (
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"github.com/f-shixiong/go-shell/lib/go/token"
)

func CompileArgs(args []ast.Expr, r *RunNode) (retArgs []interface{}) {
	for _, a := range args {
		retArgs = append(retArgs, CompileExpr(a, r))
	}
	return
}

func CompileBinary(x *ast.BinaryExpr, r *RunNode) (ret interface{}) {
	switch x.Op {
	case token.ADD: // +
		return BAdd(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.SUB: // -
		return BSub(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.MUL: // *
		return BMul(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.QUO: // /
		return BQuo(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.REM: // %
		return BRem(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.AND: // &
		return BAnd(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.OR: // |
		return BOr(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.XOR: // ^
		return BXor(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.SHL: // <<
		return BShl(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.SHR: // >>
		return BShr(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.EQL: // ==
		return CompileExpr(x.X, r) == CompileExpr(x.Y, r)
	case token.LSS: // <
		return BLss(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.NEQ:
		return CompileExpr(x.X, r) != CompileExpr(x.Y, r)
	case token.GTR: // >
		return BGtr(CompileExpr(x.X, r), CompileExpr(x.Y, r))

	//TODO
	default:
		Error("o it really happend ?,%+v", x.Op.String())
	}
	return
}

/*
	ADD // +
        SUB // -
        MUL // *
        QUO // /
        REM // %

        AND     // &
        OR      // |
        XOR     // ^
        SHL     // <<
        SHR     // >>
        AND_NOT // &^

        ADD_ASSIGN // +=
        SUB_ASSIGN // -=
        MUL_ASSIGN // *=
        QUO_ASSIGN // /=
        REM_ASSIGN // %=

        AND_ASSIGN     // &=
        OR_ASSIGN      // |=
        XOR_ASSIGN     // ^=
        SHL_ASSIGN     // <<=
        SHR_ASSIGN     // >>=
        AND_NOT_ASSIGN // &^=

        LAND  // &&
        LOR   // ||
        ARROW // <-
        INC   // ++
        DEC   // --

        EQL    // ==
        LSS    // <
        GTR    // >
        ASSIGN // =
	NOT    // !

        NEQ      // !=
        LEQ      // <=
        GEQ      // >=
        DEFINE   // :=
        ELLIPSIS // ...
*/

func CompileSelectorExpr(x *ast.SelectorExpr, r *RunNode) (ret interface{}) {
	Debug(">>>>> to_debg %#v ", x)
	Debug(">>>>> to_debg %#v ", CompileExpr(x.X, r))
	Debug(">>>>> to_debg %#v ", x.Sel)
	l := CompileExpr(x.X, r)
	switch l := l.(type) {
	case map[interface{}]interface{}:
		return l[x.Sel.Name]
	case Struct:
		return l.vals[x.Sel.Name]
	default:
		Error("what we should do l %#v", l)
	}
	return
}

func CompileUnaryExpr(x *ast.UnaryExpr, r *RunNode) (ret interface{}) {
	Debug("there x = %#v x.Op = %#v \n", x, x.Op.String())

	switch x.Op {
	case token.AND:
		r0 := CompileExpr(x.X, r)
		ret = &r0

	default:
		Error("how it ha %#v", x)
	}

	return
}
