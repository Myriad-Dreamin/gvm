package test

import (
	"fmt"
	"github.com/Myriad-Dreamin/gvm"
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	"github.com/Myriad-Dreamin/gvm/libgvm/gvm-instruction"
	"github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setStateTestCase() []abstraction.Instruction {
	return []abstraction.Instruction{
		gvm_instruction.SetState{
			Target:          "a",
			RightExpression: gvm_type.Bool(true),
		},
		gvm_instruction.SetState{
			Target:          "b",
			RightExpression: gvm_type.Bool(false),
		},
		gvm_instruction.SetState{
			Target: "c",
			RightExpression: gvm_type.BinaryExpression{
				Type:  gvm_type.RefBool,
				Sign:  gvm_type.SignLAnd,
				Left:  gvm_type.Bool(false),
				Right: gvm_type.Bool(true),
			},
		},
		gvm_instruction.SetState{
			Target: "d",
			RightExpression: gvm_type.BinaryExpression{
				Type:  gvm_type.RefBool,
				Sign:  gvm_type.SignLOr,
				Left:  gvm_type.Bool(false),
				Right: gvm_type.Bool(true),
			},
		},
	}
}

func funcSetA() []abstraction.Instruction {
	return []abstraction.Instruction{
		gvm_instruction.SetState{
			Target:          "a",
			RightExpression: gvm_type.Bool(true),
		},
	}
}

func callSetBoolFuncTestCase() []abstraction.Instruction {
	return []abstraction.Instruction{
		gvm_instruction.CallFunc{FN: "setA"},
	}
}

func BenchmarkBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runMemoryGVM(func(g *libgvm.GVMeX) {}, setStateTestCase())
	}
}

func BenchmarkPureBase(b *testing.B) {
	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	sugar.HandlerError0(g.AddFunction("main", setStateTestCase()))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Run("main")
	}
}

func BenchmarkPureSetStatus(b *testing.B) {
	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	sugar.HandlerError0(g.AddFunction("main", []abstraction.Instruction{
		gvm_instruction.SetState{
			Target:          "a",
			RightExpression: gvm_type.Bool(true),
		},
	}))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Run("main")
	}
}

//noinspection SpellCheckingInspection
type donothing struct {
}

func (d donothing) Exec(g *abstraction.ExecCtx) error {
	g.PC++
	return nil
}

func BenchmarkPureDoNothing(b *testing.B) {
	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	sugar.HandlerError0(g.AddFunction("main", []abstraction.Instruction{
		donothing{},
	}))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Run("main")
	}
}

func BenchmarkPureNilInstruction(b *testing.B) {
	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	sugar.HandlerError0(g.AddFunction("main", nil))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Run("main")
	}
}

func TestBase(t *testing.T) {
	t.Run("set state", func(t *testing.T) {
		runMemoryGVM(func(g *libgvm.GVMeX) {
			fmt.Println(g.Machine.(*libgvm.Mem).Context)
			assert.EqualValues(t, true, g.Machine.(*libgvm.Mem).Context["a"].Unwrap())
			assert.EqualValues(t, false, g.Machine.(*libgvm.Mem).Context["b"].Unwrap())
			assert.EqualValues(t, false, g.Machine.(*libgvm.Mem).Context["c"].Unwrap())
			assert.EqualValues(t, true, g.Machine.(*libgvm.Mem).Context["d"].Unwrap())
		}, setStateTestCase())
	})
	t.Run("branch condition", func(t *testing.T) {
		runMemoryGVM(func(g *libgvm.GVMeX) {

		}, nil)
	})
	t.Run("get local state", func(t *testing.T) {
		runMemoryGVM(func(g *libgvm.GVMeX) {

		}, nil)
	})
	t.Run("call function", func(t *testing.T) {
		runMemoryGVM(func(g *libgvm.GVMeX) {
			assert.EqualValues(t, true, g.Machine.(*libgvm.Mem).Context["a"].Unwrap())
		}, callSetBoolFuncTestCase())
	})
}

func funcFib() []gvm.Instruction {
	// func fib(n int64) (r int64)
	return []gvm.Instruction{
		// q := 0
		gvm_instruction.SetLocal{Target: "q", RightExpression: gvm_type.Int64(0)},
		// if n > 0 { q = fib(n - 1); }
		gvm_instruction.ConditionCallFunc{
			CallFunc: gvm_instruction.CallFunc{
				FN: "fib", Left: []string{"q"}, Right: []gvm.VTok{gvm_type.BinaryExpression{
					Type: gvm_type.RefInt64, Sign: gvm_type.SignSub, Left: gvm_type.FuncParam{T: gvm_type.RefInt64, K: 0}, Right: gvm_type.Int64(1),
				}},
			},
			Condition: gvm_type.BinaryExpression{
				Type: gvm_type.RefBool, Sign: gvm_type.SignGT, Left: gvm_type.FuncParam{T: gvm_type.RefInt64, K: 0}, Right: gvm_type.Int64(0),
			},
		},
		// r = n + q; return r
		gvm_instruction.SetFuncReturn{Target: 0, RightExpression: gvm_type.BinaryExpression{
			Type: gvm_type.RefInt64, Sign: gvm_type.SignAdd, Left: gvm_type.LocalVariable{Name: "q", Type: gvm_type.RefInt64}, Right: gvm_type.FuncParam{T: gvm_type.RefInt64, K: 0},
		}},
	}
}

//noinspection SpellCheckingInspection
func TestFibonacci(t *testing.T) {
	runMemoryGVM(func(g *libgvm.GVMeX) {
		//fmt.Println(g.Machine.(*libgvm.Mem).Context)
	}, []abstraction.Instruction{
		gvm_instruction.CallFunc{FN: "fib", Left: []string{"res"}, Right: []abstraction.VTok{gvm_type.Int64(4)}},
		doInst{g: func(g *abstraction.ExecCtx) error {
			fmt.Println("fib(4) =", g.This["res"])
			g.PC++
			return nil
		}},
	})
}

//noinspection SpellCheckingInspection
func BenchmarkFibnacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runMemoryGVM(func(g *libgvm.GVMeX) {
		}, []abstraction.Instruction{
			gvm_instruction.CallFunc{FN: "fib", Left: []string{"res"}, Right: []abstraction.VTok{
				gvm_type.Int64(3)}},
		})
	}
}
