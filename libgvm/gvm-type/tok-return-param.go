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
	v := GetReturn(g, f.K)

	if v.GetGVMType() != f.T {
		return nil, expressionTypeError(f.T, v.GetGVMType())
	}
	return v, nil
}

func FuncReturnNameKey(k int) string {
	return "_gvm_return" + strconv.Itoa(k)
}

func FuncReturnName(g *abstraction.ExecCtx, k int) string {
	s := g.This[FuncReturnNameKey(k)]
	if s == nil {
		return ""
	}
	return s.Unwrap().(string)
}

func SetReturnField(g *abstraction.ExecCtx, k int, r string) {
	g.This[FuncReturnNameKey(k)] = String(r)
}

func GetReturn(g *abstraction.ExecCtx, k int) abstraction.Ref {
	l := g.Parent[FuncReturnName(g, k)]
	if l == nil {
		return Undefined
	}
	return l
}
