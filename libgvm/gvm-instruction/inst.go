package gvm_instruction

import (
	"fmt"
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
)

func incPC(g *abstraction.ExecCtx) error {
	g.PC++
	return nil
}

func cond(g *abstraction.ExecCtx, c abstraction.VTok,
	ifFunc func(g *abstraction.ExecCtx) error, elseFunc func(g *abstraction.ExecCtx) error) error {
	v, err := c.Eval(g)
	if err != nil {
		return err
	}
	if v.GetGVMType() != gvm_type.RefBool {
		return fmt.Errorf("type error: not bool value, is %v", v.GetGVMType())
	}
	if v.Unwrap().(bool) {
		if ifFunc != nil {
			return ifFunc(g)
		}
	} else if elseFunc != nil {
		return elseFunc(g)
	}

	return nil
}
