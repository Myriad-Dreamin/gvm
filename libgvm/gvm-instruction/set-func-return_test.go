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

func TestConditionSetFuncReturn_Exec(t *testing.T) {

	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)

	sugar.HandlerError0(g.AddFunction("test", []abstraction.Instruction{
		SetState{},
		SetFuncReturn{},
	}))

	sugar.HandlerError0(g.AddFunction("main", []abstraction.Instruction{
		SetFuncReturn{},
	}))

	thisCtx := abstraction.Locals{}
	gvm_type.SetReturnField(&abstraction.ExecCtx{This: thisCtx}, 0, "a")
	gvm_type.SetReturnField(&abstraction.ExecCtx{This: thisCtx}, 1, "b")
	gvm_trap.SetReturnParamCount(&abstraction.ExecCtx{This: thisCtx}, 2)
	type fields struct {
		SetFuncReturn SetFuncReturn
		Condition     abstraction.VTok
	}
	type args struct {
		g *abstraction.ExecCtx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantPC  uint64
		wantRef abstraction.Ref
	}{
		{
			name: "eval-and-pc++",
			fields: fields{
				SetFuncReturn: SetFuncReturn{
					Target:          0,
					RightExpression: gvm_type.Bool(true),
				},
				Condition: gvm_type.Bool(true),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine: sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX),
					PC:      1,
					This:    thisCtx,
					Parent:  make(abstraction.Locals),
				},
			},
			wantPC:  2,
			wantRef: gvm_type.Bool(true),
		},
		{
			name: "not-eval-and-pc++",
			fields: fields{
				SetFuncReturn: SetFuncReturn{
					Target:          0,
					RightExpression: gvm_type.Bool(true),
				},
				Condition: gvm_type.Bool(false),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine: sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX),
					PC:      2,
					This:    thisCtx,
					Parent:  make(abstraction.Locals),
				},
			},
			wantPC:  3,
			wantRef: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &ConditionSetFuncReturn{
				SetFuncReturn: tt.fields.SetFuncReturn,
				Condition:     tt.fields.Condition,
			}
			if err := inst.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil && !assert.EqualValues(t, tt.wantRef, tt.args.g.Parent["a"]) {
				t.Errorf("Exec() got = %v, want = %v", tt.args.g.Parent["a"], tt.wantRef)
			}
		})
	}
}

func TestSetFuncReturn_Exec(t *testing.T) {

	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)

	sugar.HandlerError0(g.AddFunction("test", []abstraction.Instruction{
		SetState{},
		SetFuncReturn{},
	}))

	sugar.HandlerError0(g.AddFunction("main", []abstraction.Instruction{
		SetFuncReturn{},
	}))

	thisCtx := abstraction.Locals{}
	gvm_type.SetReturnField(&abstraction.ExecCtx{This: thisCtx}, 0, "a")
	gvm_type.SetReturnField(&abstraction.ExecCtx{This: thisCtx}, 1, "b")
	gvm_trap.SetReturnParamCount(&abstraction.ExecCtx{This: thisCtx}, 2)

	type fields struct {
		Target          int
		RightExpression abstraction.VTok
	}
	type args struct {
		g *abstraction.ExecCtx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantPC  uint64
		wantRef abstraction.Ref
	}{
		{
			name: "eval-and-pc++",
			fields: fields{
				Target:          0,
				RightExpression: gvm_type.Bool(true),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine: sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX),
					PC:      1,
					This:    thisCtx,
					Parent:  make(abstraction.Locals),
				},
			},
			wantPC:  2,
			wantRef: gvm_type.Bool(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := SetFuncReturn{
				Target:          tt.fields.Target,
				RightExpression: tt.fields.RightExpression,
			}
			if err := G.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil && !assert.EqualValues(t, tt.wantRef, tt.args.g.Parent["a"]) {
				t.Errorf("Exec() got = %v, want = %v", tt.args.g.Parent["a"], tt.wantRef)
			}
		})
	}
}
