package gvm_instruction

import (
	"fmt"
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
)

type SetFuncReturn struct {
	Target          int              `json:"target"`
	RightExpression abstraction.VTok `json:"expression"`
}

func (G SetFuncReturn) Exec(g *abstraction.ExecCtx) error {
	k, err := G.RightExpression.Eval(g)
	if err != nil {
		return err
	}
	rName := gvm_type.FuncReturnName(g, G.Target)
	if len(rName) == 0 {
		return fmt.Errorf("rName not found")
	}
	g.Parent[rName] = k
	g.PC++
	return nil
}

type ConditionSetFuncReturn struct {
	SetFuncReturn
	Condition abstraction.VTok `json:"condition"`
}

func (inst *ConditionSetFuncReturn) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.SetFuncReturn.Exec, incPC)
}
