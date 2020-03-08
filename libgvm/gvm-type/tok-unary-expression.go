package gvm_type

import "github.com/Myriad-Dreamin/gvm/internal/abstraction"

type UnaryExpression struct {
	Type abstraction.RefType `json:"type"`
	Sign SignType            `json:"sign"`
	Left abstraction.VTok    `json:"left"`
}

func (u UnaryExpression) GetGVMTok() abstraction.TokType {
	panic("implement me")
}

func (u UnaryExpression) GetGVMType() abstraction.RefType {
	panic("implement me")
}

func (u UnaryExpression) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	panic("implement me")
}
