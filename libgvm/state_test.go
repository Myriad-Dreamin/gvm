package libgvm

import (
	"fmt"
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_pushPop(t *testing.T) {
	g := sugar.HandlerError(NewGVM()).(*GVMeX)
	g2 := sugar.HandlerError(NewGVM()).(*GVMeX)

	var newFn, newFn2 = "newFn", "newFn2"
	sugar.HandlerError0(g.AddFunction("main", []abstraction.Instruction{gotoImpl{}}))
	sugar.HandlerError0(g.AddFunction(newFn, nil))
	sugar.HandlerError0(g.AddFunction(newFn2, []abstraction.Instruction{gotoImpl{}, gotoImpl{}}))
	sugar.HandlerError0(g2.AddFunction("main", []abstraction.Instruction{gotoImpl{}}))
	sugar.HandlerError0(g2.AddFunction(newFn, nil))
	sugar.HandlerError0(g2.AddFunction(newFn2, []abstraction.Instruction{gotoImpl{}, gotoImpl{}}))
	gCtx := &abstraction.ExecCtx{
		Machine:  g,
		Function: sugar.HandlerError(g.GetFunction("main")).(abstraction.Function),
		PC:       1,
		Depth:    1,
		FN:       "main",
		Parent:   abstraction.Locals{},
		This:     abstraction.Locals{"a": gvm_type.Bool(true)},
	}
	g2Ctx := &abstraction.ExecCtx{
		Machine:  g2,
		Function: sugar.HandlerError(g.GetFunction("main")).(abstraction.Function),
		PC:       1,
		Depth:    1,
		FN:       "main",
		Parent:   abstraction.Locals{"b": gvm_type.Bool(true)},
		This:     abstraction.Locals{"a": gvm_type.Bool(true)}}
	sugar.HandlerError0(PushFrame(g2Ctx, newFn))
	sugar.HandlerError0(saveLocals(gCtx, 0, gCtx.Parent))
	sugar.HandlerError0(saveLocals(g2Ctx, 0, g2Ctx.Parent))
	sugar.HandlerError0(saveFrame(gCtx))
	sugar.HandlerError0(saveFrame(g2Ctx))
	type args struct {
		g  *abstraction.ExecCtx
		fn string
	}
	tests := []struct {
		name string
		args args
		f    abstraction.Function
	}{
		{name: "pushPop", args: args{g: gCtx, fn: newFn}, f: sugar.HandlerError(g.GetFunction(newFn)).(abstraction.Function)},
		{name: "pushPopWithDepth0", args: args{g: g2Ctx, fn: newFn2},
			f: sugar.HandlerError(g.GetFunction(newFn2)).(abstraction.Function)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, pc, fn, pt, ts := tt.args.g.Function, tt.args.g.PC, tt.args.g.FN, tt.args.g.Parent, tt.args.g.This
			d := tt.args.g.Depth
			if err := PushFrame(tt.args.g, tt.args.fn); err != nil {
				t.Fatal(err)
			}
			if !assert.EqualValues(t, d+1, tt.args.g.Depth) {
				t.Errorf("want = %v, got = %v", d+1, tt.args.g.Depth)
			}
			if !assert.EqualValues(t, tt.f, tt.args.g.Function) {
				t.Errorf("want = %v, got = %v", f, tt.args.g.Function)
			}
			if !assert.EqualValues(t, make(abstraction.Locals), tt.args.g.This) {
				t.Errorf("want = %v, got = %v", make(abstraction.Locals), tt.args.g.This)
			}
			if !assert.EqualValues(t, ts, tt.args.g.Parent) {
				t.Errorf("want = %v, got = %v", ts, tt.args.g.Parent)
			}
			if !assert.EqualValues(t, 0, tt.args.g.PC) {
				t.Errorf("want = %v, got = %v", 0, tt.args.g.PC)
			}
			if !assert.EqualValues(t, tt.args.fn, tt.args.g.FN) {
				t.Errorf("want = %v, got = %v", tt.args.fn, tt.args.g.FN)
			}
			if err := popFrame(tt.args.g); err != nil && err != StopUnderFlow {
				t.Fatal(err)
			}
			if !assert.EqualValues(t, d, tt.args.g.Depth) {
				t.Errorf("want = %v, got = %v", d, tt.args.g.Depth)
			}
			if !assert.EqualValues(t, f, tt.args.g.Function) {
				t.Errorf("want = %v, got = %v", f, tt.args.g.Function)
			}
			if !assert.EqualValues(t, ts, tt.args.g.This) {
				t.Errorf("want = %v, got = %v", ts, tt.args.g.This)
			}
			if !assert.EqualValues(t, pt, tt.args.g.Parent) {
				t.Errorf("want = %v, got = %v", pt, tt.args.g.Parent)
			}
			if !assert.EqualValues(t, pc, tt.args.g.PC) {
				t.Errorf("want = %v, got = %v", pc, tt.args.g.PC)
			}
			if !assert.EqualValues(t, fn, tt.args.g.FN) {
				t.Errorf("want = %v, got = %v", fn, tt.args.g.FN)
			}

			if _, ok :=
				tt.args.g.Machine.(*GVMeX).g.Context["_gvm_locals_"+strconv.FormatUint(tt.args.g.Depth+1, 16)]; ok {
				t.Errorf("locals[%v] was not delete", tt.args.g.Depth+1)
			}
		})
	}
}

//noinspection SpellCheckingInspection,SpellCheckingInspection,SpellCheckingInspection
func Benchmark_serializeMapSR(b *testing.B) {
	var mp = map[string]abstraction.Ref{
		"a":          gvm_type.Bool(true),
		"bbbb":       gvm_type.Bool(true),
		"ccccccc":    gvm_type.Bool(true),
		"dddddddddd": gvm_type.Bool(true),
	}
	g := sugar.HandlerError(NewGVM()).(*GVMeX)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sugar.HandlerError0(saveLocals(g, 0, mp))
	}
}
func Benchmark_deserializeMapSR(b *testing.B) {
	var mp = map[string]abstraction.Ref{
		"a": gvm_type.Bool(true),
	}
	g := sugar.HandlerError(NewGVM()).(*GVMeX)
	sugar.HandlerError0(saveLocals(g, 0, mp))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sugar.HandlerError(loadLocals(g, 0)).(map[string]abstraction.Ref)
	}
}

func Test_serializeMapSR(t *testing.T) {
	var mp = map[string]abstraction.Ref{
		"a": gvm_type.Bool(true),
		"b": gvm_type.String("123132124"),
	}

	g := sugar.HandlerError(NewGVM()).(*GVMeX)

	sugar.HandlerError0(saveLocals(g, 0, mp))
	fmt.Println(g.g.Context)
	fmt.Println(len(g.g.Context["_gvm_locals_0"].Unwrap().([]byte)))
	mp2 := sugar.HandlerError(loadLocals(g, 0)).(map[string]abstraction.Ref)

	assert.EqualValues(t, mp, mp2)
	fmt.Println(mp2)

	mp = map[string]abstraction.Ref{
		"a": gvm_type.Bool(true),
		"b": gvm_type.Bool(true),
	}

	sugar.HandlerError0(saveLocals(g, 0, mp))
	fmt.Println(g.g.Context)
	fmt.Println(len(g.g.Context["_gvm_locals_0"].Unwrap().([]byte)))
	mp2 = sugar.HandlerError(loadLocals(g, 0)).(map[string]abstraction.Ref)

	assert.EqualValues(t, mp, mp2)
	fmt.Println(mp2)
}
