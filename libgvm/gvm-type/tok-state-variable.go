package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
)

type StateVariable struct {
	Type  abstraction.RefType
	Field string
}

func (s StateVariable) GetGVMTok() abstraction.TokType {
	return TokStateVariable
}

func (s StateVariable) GetGVMType() abstraction.RefType {
	return s.Type
}

func (s StateVariable) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	return g.Load(s.Field, s.Type)
}
