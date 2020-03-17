package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"strconv"
)

type FuncParam struct {
	T abstraction.RefType
	K int
}

func (f FuncParam) GetGVMTok() abstraction.TokType {
	return TokFuncParam
}

func (f FuncParam) GetGVMType() abstraction.RefType {
	return f.T
}

func (f FuncParam) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	l := GetParam(g, f.K)

	if l.GetGVMType() != f.T {
		return nil, expressionTypeError(f.T, l.GetGVMType())
	}

	return l, nil
}

func GetParam(g *abstraction.ExecCtx, k int) abstraction.Ref {
	l, ok := g.This[FuncParamName(k)]
	if !ok {
		return Undefined
	}
	return l
}

func AddFuncParam(g *abstraction.ExecCtx, k int, r abstraction.Ref) {
	g.This[FuncParamName(k)] = r
}

func FuncParamName(k int) string {
	return strconv.Itoa(k)
}
