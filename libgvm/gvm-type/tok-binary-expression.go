package gvm_type

import (
	"fmt"
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
)

type BinaryExpression struct {
	Type  abstraction.RefType `json:"type"`
	Sign  SignType            `json:"sign"`
	Left  abstraction.VTok            `json:"left"`
	Right abstraction.VTok            `json:"right"`
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
	switch b.Sign {
	case SignEQ:
		return EQ(l, r)
	case SignLE:
		return LE(l, r)
	case SignLT:
		return LT(l, r)
	case SignGE:
		return GE(l, r)
	case SignGT:
		return GT(l, r)
	case SignLAnd:
		return LAnd(l, r)
	case SignLOr:
		return LOr(l, r)
	case SignADD:
		return Add(l, r)
	case SignSUB:
		return Sub(l, r)
	case SignMUL:
		return Mul(l, r)
	case SignQUO:
		return Quo(l, r)
	case SignREM:
		return Rem(l, r)
	default:
		return nil, fmt.Errorf("unknown sign_type: %v", b.Sign)
	}
}

