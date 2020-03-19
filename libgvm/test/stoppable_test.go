package test

import (
	"errors"
	"fmt"
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	gvm_instruction "github.com/Myriad-Dreamin/gvm/libgvm/gvm-instruction"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"testing"
)

type MyMachine struct {
	abstraction.Machine
	SomeCondition bool
}

type StopToCheckIfTrue struct {
}

var StopInterruption = errors.New("stop")

func (s StopToCheckIfTrue) Exec(g *abstraction.ExecCtx) error {
	if (g.Machine.(*MyMachine)).SomeCondition {
		g.PC++
		return nil
	} else {
		return StopInterruption
	}
}

func TestStopOnce(t *testing.T) {
	_runMemoryGVM(nil, []abstraction.Instruction{
		gvm_instruction.SetState{
			Target:          "a",
			RightExpression: gvm_type.Bool(true),
		},
		gvm_instruction.SetLocal{
			Target:          "b",
			RightExpression: gvm_type.Bool(true),
		},
		doInst{g: func(g *abstraction.ExecCtx) error {
			if sugar.HandlerError(g.Load("a", gvm_type.RefBool)).(abstraction.Ref).Unwrap().(bool) != true {
				panic("bad value")
			}
			g.PC++
			return nil
		}},
		doInst{g: func(g *abstraction.ExecCtx) error {
			if g.This["b"].Unwrap().(bool) != true {
				panic("bad value")
			}
			g.PC++
			return nil
		}},
		StopToCheckIfTrue{},
		doInst{g: func(g *abstraction.ExecCtx) error {
			if sugar.HandlerError(g.Load("a", gvm_type.RefBool)).(abstraction.Ref).Unwrap().(bool) != true {
				panic("bad value")
			}
			g.PC++
			return nil
		}},
		doInst{g: func(g *abstraction.ExecCtx) error {
			if g.This["b"].Unwrap().(bool) != true {
				panic("bad value")
			}
			g.PC++
			return nil
		}},
	}, func(k *libgvm.GVMeX) *libgvm.GVMeX {
		k.Machine = &MyMachine{
			Machine:       k.Machine,
			SomeCondition: false,
		}
		return k
	}, func(g *libgvm.GVMeX, err error) bool {
		if err == StopInterruption {
			g.Machine.(*MyMachine).SomeCondition = true
			return true
		}
		fmt.Println(err)
		return false
	})
}
