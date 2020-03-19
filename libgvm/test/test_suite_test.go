package test

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
)

type doInst struct {
	g func(g *abstraction.ExecCtx) error
}

func (d doInst) Exec(g *abstraction.ExecCtx) error {
	return d.g(g)
}

func _runMemoryGVM(
	callback func(g *libgvm.GVMeX), instructions []abstraction.Instruction, wrap func(k *libgvm.GVMeX) *libgvm.GVMeX,
	handleErr ...func(g *libgvm.GVMeX, err error) bool) {
	var handle = func(g *libgvm.GVMeX, err error) bool {
		if err != nil && err != libgvm.StopUnderFlow {
			panic(err)
		}
		return true
	}
	if len(handleErr) > 0 {
		handle = handleErr[0]
	}
	g := wrap(sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX))
	sugar.HandlerError0(g.AddFunction("main", instructions))
	sugar.HandlerError0(g.AddFunction("setA", funcSetA()))
	sugar.HandlerError0(g.AddFunction("fib", funcFib()))
	var err error
	for err = g.Run("main"); err != libgvm.StopUnderFlow; err = g.Continue() {
		if !handle(g, err) {
			return
		}
	}
	if callback != nil {
		callback(g)
	}
}

func runMemoryGVM(callback func(g *libgvm.GVMeX), instructions []abstraction.Instruction) {
	_runMemoryGVM(callback, instructions, func(k *libgvm.GVMeX) *libgvm.GVMeX {
		return k
	})
}
