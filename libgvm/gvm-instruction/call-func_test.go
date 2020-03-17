package gvm_instruction

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	gvm_trap "github.com/Myriad-Dreamin/gvm/libgvm/gvm-trap"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConditionCallFunc_Exec(t *testing.T) {
	type ComparingCtxField struct {
		abstraction.Function
		Depth        uint64
		PC           uint64
		FN           string
		Parent, This abstraction.Locals
	}

	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	g2 := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)

	sugar.HandlerError0(g.AddFunction("test", []abstraction.Instruction{
		SetState{},
		SetFuncReturn{},
	}))

	sugar.HandlerError0(g2.AddFunction("test", []abstraction.Instruction{
		SetFuncReturn{},
		SetState{},
	}))

	sugar.HandlerError0(g.AddFunction("main", []abstraction.Instruction{
		SetFuncReturn{},
	}))

	sugar.HandlerError0(g2.AddFunction("main", []abstraction.Instruction{
		SetState{},
	}))

	fn := sugar.HandlerError(g.GetFunction("test")).(abstraction.Function)

	//fn2 := sugar.HandlerError(g2.GetFunction("test")).(abstraction.Function)

	main := sugar.HandlerError(g.GetFunction("main")).(abstraction.Function)

	main2 := sugar.HandlerError(g2.GetFunction("main")).(abstraction.Function)

	want1 := abstraction.Locals{}
	gvm_type.SetReturnField(&abstraction.ExecCtx{This:want1}, 0, "a")
	gvm_type.SetReturnField(&abstraction.ExecCtx{This:want1}, 1, "b")
	gvm_trap.SetReturnParamCount(&abstraction.ExecCtx{This: want1}, 2)
	gvm_type.AddFuncParam(&abstraction.ExecCtx{This: want1}, 0, gvm_type.Bool(true))
	gvm_type.AddFuncParam(&abstraction.ExecCtx{This: want1}, 1, gvm_type.Uint64(1))
	gvm_type.AddFuncParam(&abstraction.ExecCtx{This: want1}, 2, gvm_type.Uint64(2))
	gvm_trap.SetFuncParamCount(&abstraction.ExecCtx{This: want1}, 3)


	type fields struct {
		CallFunc  CallFunc
		Condition abstraction.VTok
	}
	type args struct {
		g *abstraction.ExecCtx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantCtx  *ComparingCtxField
	}{
		{
			name: "call-and-pc++",
			fields: fields{
				CallFunc: CallFunc{
					FN:    "test",
					Left:  []string{"a", "b"},
					Right: []abstraction.VTok{gvm_type.Bool(true), gvm_type.Uint64(1), gvm_type.Uint64(2)},
				},
				Condition: gvm_type.Bool(true),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine: g,
					PC:      1,
					Function: main,
					FN:      "main",
					Depth:   2,
					This:    abstraction.Locals{"c": gvm_type.Bool(true)},
					Parent:  abstraction.Locals{"d": gvm_type.Bool(true)},
				},
			},
			wantCtx:  &ComparingCtxField{
				PC:      0,
				Function: fn,
				FN:      "test",
				Depth:   3,
				This:    want1,
				Parent:  abstraction.Locals{"c": gvm_type.Bool(true)},
			},
		},
		{
			name: "not-call-and-pc++",
			fields: fields{
				CallFunc: CallFunc{
					FN:    "test",
					Left:  []string{"a", "b"},
					Right: []abstraction.VTok{gvm_type.Bool(true), gvm_type.Uint64(1), gvm_type.Uint64(2)},
				},
				Condition: gvm_type.Bool(false),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine: g2,
					PC:      1,
					Function: main2,
					FN:      "main",
					Depth:   2,
					This:    abstraction.Locals{"c": gvm_type.Bool(true)},
					Parent:  abstraction.Locals{"d": gvm_type.Bool(true)},
				},
			},
			wantCtx:  &ComparingCtxField{
				PC:      2,
				Function: main2,
				FN:      "main",
				Depth:   2,
				This:    abstraction.Locals{"c": gvm_type.Bool(true)},
				Parent:  abstraction.Locals{"d": gvm_type.Bool(true)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := ConditionCallFunc{
				CallFunc:  tt.fields.CallFunc,
				Condition: tt.fields.Condition,
			}

			if err := inst.Exec(tt.args.g); err != nil {
				if trap, ok := err.(abstraction.Trap); ok {
					err = trap.DoTrap(tt.args.g)
					if (err != nil) != tt.wantErr {
						t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
					} else if err == nil && !assert.EqualValues(t, tt.wantCtx, &ComparingCtxField{
						Function: tt.args.g.Function,
						Depth:    tt.args.g.Depth,
						PC:       tt.args.g.PC,
						FN:       tt.args.g.FN,
						Parent:  tt.args.g.Parent,
						This:    tt.args.g.This,
					}) {
						t.Errorf("Exec() got = %v, want %v", &ComparingCtxField{
							Function: tt.args.g.Function,
							Depth:    tt.args.g.Depth,
							PC:       tt.args.g.PC,
							FN:       tt.args.g.FN,
							Parent:  tt.args.g.Parent,
							This:    tt.args.g.This,
						}, tt.wantCtx)
					}
				} else if true != tt.wantErr {
					t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
