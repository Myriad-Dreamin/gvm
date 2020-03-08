package gvm_type

import "github.com/Myriad-Dreamin/gvm/internal/abstraction"

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
