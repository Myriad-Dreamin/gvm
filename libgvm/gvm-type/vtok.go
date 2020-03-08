package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"strconv"
)

type stateVariable interface {
	GetFieldGVM() string
}

type LocalStateVariable interface {
	abstraction.VTok
	stateVariable
}

type Constant = abstraction.Ref

type UnaryExpression interface {
	abstraction.VTok
	GetSign() SignType
	GetLeftTok() abstraction.VTok
}

type LocalVariable struct {
	Name string
	Type abstraction.RefType
}

func (l LocalVariable) GetGVMTok() abstraction.TokType {
	return TokLocalVariable
}

func (l LocalVariable) GetGVMType() abstraction.RefType {
	return l.Type
}

func (l LocalVariable) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	return g.This[l.Name], nil
}

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
	return g.This[strconv.Itoa(k)]
}

func FuncParamName(k int) string {
	return strconv.Itoa(k)
}

func FuncReturnName(g *abstraction.ExecCtx, k int) string {
	return g.This["_gvm_return"+strconv.Itoa(k)].Unwrap().(string)
}

func GetReturn(g *abstraction.ExecCtx, k int) abstraction.Ref {
	return g.Parent[g.This["_gvm_return"+strconv.Itoa(k)].Unwrap().(string)]
}
