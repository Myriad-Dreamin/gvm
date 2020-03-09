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
	v, ok := g.This[l.Name]
	if !ok {
		v = Undefined
	}

	if v.GetGVMType() != l.Type {
		return nil, expressionTypeError(l.Type, v.GetGVMType())
	}
	return v, nil
}
