package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
)

type StateVariable struct {
	Field string
}

func (s StateVariable) GetGVMTok() abstraction.TokType {
	panic("implement me")
}

func (s StateVariable) GetGVMType() abstraction.RefType {
	panic("implement me")
}

func (s StateVariable) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	panic("implement me")
}
