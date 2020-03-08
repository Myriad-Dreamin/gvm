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
	return GetParam(g, f.K), nil
}

func GetParam(g *abstraction.ExecCtx, k int) abstraction.Ref {
	return g.This[FuncParamName(k)]
}

func FuncParamName(k int) string {
	return strconv.Itoa(k)
}

func FuncReturnName(g *abstraction.ExecCtx, k int) string {
	return g.This["_gvm_return"+strconv.Itoa(k)].Unwrap().(string)
}

func GetReturn(g *abstraction.ExecCtx, k int) abstraction.Ref {
	return g.Parent[FuncReturnName(g, k)]
}
