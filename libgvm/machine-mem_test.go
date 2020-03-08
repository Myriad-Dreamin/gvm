package libgvm_test

import (
	"fmt"
	"github.com/Myriad-Dreamin/gvm"
	"github.com/Myriad-Dreamin/gvm/gvm-instruction"
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type doInst struct {
	g func(g *abstraction.ExecCtx) error
}

func (d doInst) Exec(g *abstraction.ExecCtx) error {
	return d.g(g)
}

func runMemoryGVM(callback func(g *libgvm.GVMeX), instructions []abstraction.Instruction) {
	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	sugar.HandlerError0(g.AddFunction("main", instructions))
	sugar.HandlerError0(g.AddFunction("setA", funcSetA()))
	sugar.HandlerError0(g.AddFunction("fib", funcFib()))
	var err error
	for err = g.Run("main"); err == nil; {
		err = g.Run("main")
		time.Sleep(time.Second)
	}
	callback(g)
}

type BinaryExpression struct {
	Type  abstraction.RefType `json:"type"`
	Sign  libgvm.SignType     `json:"sign"`
	Left  gvm.VTok            `json:"left"`
	Right gvm.VTok            `json:"right"`
}

func (b BinaryExpression) GetGVMTok() abstraction.TokType {
	return libgvm.TokBinaryExpression
}

func (b BinaryExpression) GetGVMType() abstraction.RefType {
	return b.Type
}

func (b BinaryExpression) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	l, err := b.Left.Eval(g)
	if err != nil {
		return nil, err
	}
	r, err := b.Right.Eval(g)
	if err != nil {
		return nil, err
	}
	switch b.Sign {
	case libgvm.SignEQ:
		return libgvm.EQ(l, r)
	case libgvm.SignLE:
		return libgvm.LE(l, r)
	case libgvm.SignLT:
		return libgvm.LT(l, r)
	case libgvm.SignGE:
		return libgvm.GE(l, r)
	case libgvm.SignGT:
		return libgvm.GT(l, r)
	case libgvm.SignLAnd:
		return libgvm.LAnd(l, r)
	case libgvm.SignLOr:
		return libgvm.LOr(l, r)
	case libgvm.SignADD:
		return libgvm.Add(l, r)
	case libgvm.SignSUB:
		return libgvm.Sub(l, r)
	case libgvm.SignMUL:
		return libgvm.Mul(l, r)
	case libgvm.SignQUO:
		return libgvm.Quo(l, r)
	case libgvm.SignREM:
		return libgvm.Rem(l, r)
	default:
		return nil, fmt.Errorf("unknown sign_type: %v", b.Sign)
	}
}

func setStateTestCase() []abstraction.Instruction {
	return []abstraction.Instruction{
		gvm_instruction.SetState{
			Target:          "a",
			RightExpression: libgvm.Bool(true),
		},
		gvm_instruction.SetState{
			Target:          "b",
			RightExpression: libgvm.Bool(false),
		},
		gvm_instruction.SetState{
			Target: "c",
			RightExpression: BinaryExpression{
				Type:  libgvm.RefBool,
				Sign:  libgvm.SignLAnd,
				Left:  libgvm.Bool(false),
				Right: libgvm.Bool(true),
			},
		},
		gvm_instruction.SetState{
			Target: "d",
			RightExpression: BinaryExpression{
				Type:  libgvm.RefBool,
				Sign:  libgvm.SignLOr,
				Left:  libgvm.Bool(false),
				Right: libgvm.Bool(true),
			},
		},
	}
}

func funcSetA() []abstraction.Instruction {
	return []abstraction.Instruction{
		gvm_instruction.SetState{
			Target:          "a",
			RightExpression: libgvm.Bool(true),
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
			RightExpression: libgvm.Bool(true),
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
		gvm_instruction.SetLocal{Target: "q", RightExpression: libgvm.Int64(0)},
		// if n > 0 { q = fib(n - 1); }
		gvm_instruction.ConditionCallFunc{
			CallFunc: gvm_instruction.CallFunc{
				FN: "fib", Left: []string{"q"}, Right: []gvm.VTok{BinaryExpression{
					Type: libgvm.RefInt64, Sign: libgvm.SignSUB, Left: libgvm.FuncParam{T: libgvm.RefInt64, K: 0}, Right: libgvm.Int64(1),
				}},
			},
			Condition: BinaryExpression{
				Type: libgvm.RefBool, Sign: libgvm.SignGT, Left: libgvm.FuncParam{T: libgvm.RefInt64, K: 0}, Right: libgvm.Int64(0),
			},
		},
		// r = n + q; return r
		gvm_instruction.SetFuncReturn{Target: 0, RightExpression: BinaryExpression{
			Type: libgvm.RefInt64, Sign: libgvm.SignADD, Left: libgvm.LocalVariable{Name: "q"}, Right: libgvm.FuncParam{T: libgvm.RefInt64, K: 0},
		}},
	}
}

//noinspection SpellCheckingInspection
func TestFibonacci(t *testing.T) {
	runMemoryGVM(func(g *libgvm.GVMeX) {
		//fmt.Println(g.Machine.(*libgvm.Mem).Context)
	}, []abstraction.Instruction{
		gvm_instruction.CallFunc{FN: "fib", Left: []string{"res"}, Right: []abstraction.VTok{libgvm.Int64(3)}},
		doInst{g: func(g *abstraction.ExecCtx) error {
			fmt.Println("fib(3) =", g.This["res"])
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
				libgvm.Int64(3)}},
		})
	}
}
