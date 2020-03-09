package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"strconv"
)

type FuncReturnParam struct {
	T abstraction.RefType
	K int
}

func (f FuncReturnParam) GetGVMTok() abstraction.TokType {
	return TokReturnParam
}

func (f FuncReturnParam) GetGVMType() abstraction.RefType {
	return f.T
}

func (f FuncReturnParam) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	return GetReturn(g, f.K), nil
}

func FuncReturnName(g *abstraction.ExecCtx, k int) string {
	return g.This["_gvm_return"+strconv.Itoa(k)].Unwrap().(string)
}

func GetReturn(g *abstraction.ExecCtx, k int) abstraction.Ref {
	return g.Parent[FuncReturnName(g, k)]
}
