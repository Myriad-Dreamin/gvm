package gvm_instruction

import "github.com/Myriad-Dreamin/gvm/internal/abstraction"

type Goto struct {
	Index uint64 `json:"goto"`
}

func (inst *Goto) Exec(g *abstraction.ExecCtx) error {
	g.PC = inst.Index
	return nil
}

type ConditionGoto struct {
	Goto
	Condition abstraction.VTok `json:"condition"`
}

func (inst *ConditionGoto) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.Goto.Exec, incPC)
}
