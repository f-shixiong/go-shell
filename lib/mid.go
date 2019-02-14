package lib

import (
	"fmt"
	"go/ast"
	"go/token"
)

func CompileArgs(args []ast.Expr, r *RunNode) (retArgs []interface{}) {
	for _, a := range args {
		retArgs = append(retArgs, CompileExpr(a, r))
	}
	return
}

func CompileBinary(x *ast.BinaryExpr, r *RunNode) (ret interface{}) {
	switch x.Op {
	case token.ADD: //+
		return BAdd(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.SUB:
		return BSub(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.MUL:
		return BMul(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.QUO:
		return BQuo(CompileExpr(x.X, r), CompileExpr(x.Y, r))
	case token.EQL: //==
		return CompileExpr(x.X, r) == CompileExpr(x.Y, r)
	//TODO
	default:
		fmt.Println("o it really happend ?")
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
