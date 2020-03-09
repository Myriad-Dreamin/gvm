package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
)

type BinaryExpression struct {
	Type  abstraction.RefType `json:"type"`
	Sign  SignType            `json:"sign"`
	Left  abstraction.VTok    `json:"left"`
	Right abstraction.VTok    `json:"right"`
}

func (b BinaryExpression) GetGVMTok() abstraction.TokType {
	return TokBinaryExpression
}

func (b BinaryExpression) GetGVMType() abstraction.RefType {
	return b.Type
}

func (b BinaryExpression) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	l, err := b.Left.Eval(g)
	if err != nil {
		return nil, err
	}
	r, err := b.Right.Eval(g)
	if err != nil {
		return nil, err
	}
	l, err = BiCalc(l, r, b.Sign)
	if err != nil {
		return nil, err
	}
	if l.GetGVMType() != b.Type {
		return nil, expressionTypeError(b.Type, l.GetGVMType())
	}
	return l, nil
}
