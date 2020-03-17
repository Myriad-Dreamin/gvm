package gvm_instruction

import "github.com/Myriad-Dreamin/gvm/internal/abstraction"

type SetParentLocal struct {
	Target          string           `json:"target"`
	RightExpression abstraction.VTok `json:"expression"`
}

func (G SetParentLocal) Exec(g *abstraction.ExecCtx) error {
	k, err := G.RightExpression.Eval(g)
	if err != nil {
		return err
	}
	g.Parent[G.Target] = k
	g.PC++
	return nil
}

type ConditionSetParentLocal struct {
	SetParentLocal
	Condition abstraction.VTok `json:"condition"`
}

func (inst *ConditionSetParentLocal) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.SetParentLocal.Exec, incPC)
}
