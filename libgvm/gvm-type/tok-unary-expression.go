package gvm_type

import (
	"fmt"
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
)

type UnaryExpression struct {
	Type abstraction.RefType `json:"type"`
	Sign SignType            `json:"sign"`
	Left abstraction.VTok    `json:"left"`
}

func (u UnaryExpression) GetGVMTok() abstraction.TokType {
	return TokUnaryExpression
}

func (u UnaryExpression) GetGVMType() abstraction.RefType {
	return u.Type
}

func (u UnaryExpression) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	l, err := u.Left.Eval(g)
	if err != nil {
		return nil, err
	}
	switch u.Sign {
	case SignLNot:
		return LNot(l)
	default:
		return nil, fmt.Errorf("unknown sign_type: %v", u.Sign)
	}
}
