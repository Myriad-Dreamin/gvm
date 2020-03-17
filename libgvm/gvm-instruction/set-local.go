package gvm_instruction

import "github.com/Myriad-Dreamin/gvm/internal/abstraction"

type SetLocal struct {
	Target          string           `json:"target"`
	RightExpression abstraction.VTok `json:"expression"`
}

func (G SetLocal) Exec(g *abstraction.ExecCtx) error {
	k, err := G.RightExpression.Eval(g)
	if err != nil {
		return err
	}
	g.This[G.Target] = k
	g.PC++
	return nil
}

type ConditionSetLocal struct {
	SetLocal
	Condition abstraction.VTok `json:"condition"`
}

func (inst *ConditionSetLocal) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.SetLocal.Exec, incPC)
}
