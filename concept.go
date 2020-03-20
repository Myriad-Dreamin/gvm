package gvm

import "github.com/Myriad-Dreamin/gvm/internal/abstraction"

type (
	TokType = abstraction.TokType
	RefType = abstraction.RefType

	Ref  = abstraction.Ref
	VTok = abstraction.VTok

	Function    = abstraction.Function
	ExecCtx     = abstraction.ExecCtx
	Machine     = abstraction.Machine
	Trap = abstraction.Trap

	Locals abstraction.Locals
)
