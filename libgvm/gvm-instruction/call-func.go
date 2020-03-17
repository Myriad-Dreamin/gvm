package gvm_instruction

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm/gvm-trap"
)

type CallFunc = gvm_trap.CallFunc

type ConditionCallFunc struct {
	CallFunc
	Condition abstraction.VTok `json:"condition"`
}

func (inst ConditionCallFunc) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.CallFunc.Exec, incPC)
}
