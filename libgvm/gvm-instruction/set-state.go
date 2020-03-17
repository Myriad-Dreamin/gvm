package gvm_instruction

import "github.com/Myriad-Dreamin/gvm/internal/abstraction"

type SetState struct {
	Target          string           `json:"target"`
	RightExpression abstraction.VTok `json:"expression"`
}

func (G SetState) Exec(g *abstraction.ExecCtx) error {
	k, err := G.RightExpression.Eval(g)
	if err != nil {
		return err
	}
	err = g.Save(G.Target, k)
	if err != nil {
		return err
	}

	g.PC++
	return nil
}

type ConditionSetState struct {
	SetState
	Condition abstraction.VTok `json:"condition"`
}

func (inst *ConditionSetState) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.SetState.Exec, incPC)
}
