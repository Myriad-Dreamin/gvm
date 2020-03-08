package libgvm

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
)

//noinspection GoUnusedConst
const (
	TokConstant abstraction.TokType = iota
	TokStateVariable
	TokLocalStateVariable
	TokBinaryExpression
	TokUnaryExpression

	TokFuncParam
	TokLocalVariable
)
