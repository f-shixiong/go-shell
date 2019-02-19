package lib

import (
	"github.com/f-shixiong/go-shell/lib/go/ast"
	"github.com/f-shixiong/go-shell/lib/go/token"
	"github.com/spf13/cast"
)

func CompileStmt(stmt ast.Stmt, r *RunNode) (hasR bool, hasBreak bool, ret []interface{}) {
	switch stmt := stmt.(type) {
	case *ast.AssignStmt:
		CompileAssignStmt(stmt, r)
	case *ast.DeclStmt:
		CompileDeclStmt(stmt, r)
	case *ast.ExprStmt:
		CompileExpr(stmt.X, r)
	case *ast.RangeStmt:
		Debug("range-> 0 : %#v, 1 : %#v, 2 : %#v, 3: %#v", stmt.Key, stmt.Value, stmt.X, stmt.Body.List)
		hasR, ret = CompileRangeStmt(stmt, r)
		if hasR {
			return
		}
	case *ast.ShellStmt:
		CompileShell(stmt, r)
	case *ast.IncDecStmt:
		CompileIncDec(stmt, r)
	case *ast.ForStmt:
		hasR, ret = CompileForStmt(stmt, r)
		if hasR {
			return
		}
	case *ast.ReturnStmt:
		hasR = true
		for _, result := range stmt.Results {
			ret = append(ret, CompileExpr(result, r))
		}
	case *ast.SwitchStmt:
		hasR, hasBreak, ret = CompileSwitchStmt(stmt, r)
		if hasR || hasBreak {
			return
		}
	case *ast.IfStmt:
		rc := r.Child()
		hasR, hasBreak, ret = CompileIfStmt(stmt, rc)
		if hasR || hasBreak {
			return
		}
	case *ast.EmptyStmt:
	case nil:
	case *ast.BranchStmt:
		Debug("br = %#v,token=%#v", stmt, stmt.Tok.String())
		if stmt.Tok == token.BREAK {
			hasBreak = true
			return
		}
	default:
		Error("undefind value -> %#v", stmt)
	}
	return
}

func CompileIfStmt(stmt *ast.IfStmt, r *RunNode) (hasR bool, hasBreak bool, ret []interface{}) {
	Debug("stmt = %#v,else = %#v", stmt, stmt.Else)
	CompileStmt(stmt.Init, r)
	cond := cast.ToBool(CompileExpr(stmt.Cond, r))
	if cond {
		for _, sc := range stmt.Body.List {
			hasR, hasBreak, ret = CompileStmt(sc, r)
			if hasR || hasBreak {
				return
			}
		}
	} else {
		if el, ok := stmt.Else.(*ast.BlockStmt); ok {
			for _, sc := range el.List {
				hasR, _, ret = CompileStmt(sc, r)
				if hasR {
					return
				}
			}
		}
		if el, ok := stmt.Else.(*ast.IfStmt); ok {
			CompileIfStmt(el, r)
		}
	}
	return
}

func CompileForStmt(stmt *ast.ForStmt, r *RunNode) (hasR bool, ret []interface{}) {
	Debug("\ninit -> %#v\ncond -> %#v\npost -> %#v\nbody -> %#v", stmt.Init, stmt.Cond, stmt.Post, stmt.Body)
	rc := r.Child()
	switch init := stmt.Init.(type) {
	case *ast.AssignStmt:
		CompileAssignStmt(init, rc)
	case nil:
	default:
		Error("sha? %#v", init)
	}
	hasBreak := false
	for cast.ToBool(CompileExpr(stmt.Cond, rc)) {
		for _, stmt := range stmt.Body.List {
			hasR, hasBreak, ret = CompileStmt(stmt, rc)
			if hasR || hasBreak {
				return
			}
		}
		CompileStmt(stmt.Post, rc)
	}
	return
}

func CompileSwitchStmt(stmt *ast.SwitchStmt, r *RunNode) (hasR bool, hasBreak bool, ret []interface{}) {
	Debug("\ninit->%#v\ntag->%#v\nbody->%#v\n", stmt.Init, CompileExpr(stmt.Tag, r), stmt.Body.List)
	tag := CompileExpr(stmt.Tag, r)
	rc := r.Child()
	_ = tag
	var defaultCase *ast.CaseClause
	for _, sc := range stmt.Body.List {
		cc := sc.(*ast.CaseClause)
		if len(cc.List) == 0 {
			defaultCase = cc
			continue
		}
		//Test("colon->%#v,list0->%#v", cc.Body[0], cc.List[0])
		var match bool
		for _, l := range cc.List {
			if CompileExpr(l, rc) == tag {
				match = true
				break
			}
		}
		if match {
			for _, stmt := range cc.Body {
				hasR, hasBreak, ret = CompileStmt(stmt, rc)
				if hasR || hasBreak {
					return
				}
			}
			return
		}
	}
	if defaultCase != nil {
		for _, stmt := range defaultCase.Body {
			hasR, _, ret = CompileStmt(stmt, rc)
			if hasR {
				return
			}
		}
	}
	return
}
