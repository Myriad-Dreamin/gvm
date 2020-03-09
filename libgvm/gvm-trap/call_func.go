package gvm_trap

import (
	"fmt"
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	"github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
)

type CallFunc struct {
	FN    string             `json:"fn"`
	Left  []string           `json:"left"`
	Right []abstraction.VTok `json:"right"`
}

func (c CallFunc) Error() string {
	return fmt.Sprintf("trap calling: %v", c.FN)
}

func (c CallFunc) Exec(g *abstraction.ExecCtx) error {
	g.PC++
	return c
}

func (c CallFunc) DoTrap(g *abstraction.ExecCtx) (err error) {
	var refs = make([]abstraction.Ref, len(c.Right))
	for l := range c.Right {
		refs[l], err = c.Right[l].Eval(g)
		if err != nil {
			return err
		}
	}
	err = libgvm.PushFrame(g, c.FN)
	if err != nil {
		return err
	}
	for l := range c.Right {
		gvm_type.AddParam(g, l, refs[l])
	}
	g.This["gvm_fp_cnt"] = gvm_type.Uint64(len(c.Right))

	for l := range c.Left {
		g.This[gvm_type.FuncReturnName(g, l)] = gvm_type.String(c.Left[l])
	}
	g.This["gvm_fr_cnt"] = gvm_type.Uint64(len(c.Left))
	return
}
