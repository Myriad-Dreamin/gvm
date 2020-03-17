package gvm_instruction

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConditionSetState_Exec(t *testing.T) {
	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	sugar.HandlerError0(g.Save("test", gvm_type.Bool(true)))

	type fields struct {
		SetState  SetState
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
		wantPC uint64
		wantRef abstraction.Ref
	}{
		{
			name: "eval-and-pc++",
			fields: fields{
				SetState: SetState{
					Target:          "a",
					RightExpression: gvm_type.Bool(true),
				},
				Condition: gvm_type.Bool(true),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine:  sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX),
					PC:       1,
				},
			},
			wantPC: 2,
			wantRef: gvm_type.Bool(true),
		},
		{
			name: "not-eval-and-pc++",
			fields: fields{
				SetState: SetState{
					Target:          "a",
					RightExpression: gvm_type.Bool(true),
				},
				Condition: gvm_type.Bool(false),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine:  sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX),
					PC:       2,
				},
			},
			wantPC: 3,
			wantRef: gvm_type.Undefined,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &ConditionSetState{
				SetState:  tt.fields.SetState,
				Condition: tt.fields.Condition,
			}
			if err := inst.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			} else if err ==nil && !assert.EqualValues(t, tt.wantRef,
				sugar.HandlerError(tt.args.g.Machine.Load("a", gvm_type.RefBool))) {
				t.Errorf("Exec() got = %v, want %v",
					sugar.HandlerError(tt.args.g.Machine.Load("a", gvm_type.RefBool)), tt.wantRef)
			} else if err == nil && tt.args.g.PC != tt.wantPC {
				t.Errorf("Exec() got = %v, want %v", tt.args.g.PC, tt.wantPC)
			}
		})
	}
}

func TestSetState_Exec(t *testing.T) {
	type fields struct {
		Target          string
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
		wantPC uint64
		wantRef abstraction.Ref
	}{
		{
			name: "eval-and-pc++",
			fields: fields{
				Target:          "a",
				RightExpression: gvm_type.Bool(true),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine:  sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX),
					PC:       1,
				},
			},
			wantPC: 2,
			wantRef: gvm_type.Bool(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := SetState{
				Target:          tt.fields.Target,
				RightExpression: tt.fields.RightExpression,
			}
			if err := G.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			} else if err ==nil && !assert.EqualValues(t, tt.wantRef,
				sugar.HandlerError(tt.args.g.Machine.Load("a", gvm_type.RefBool))) {
				t.Errorf("Exec() got = %v, want %v",
					sugar.HandlerError(tt.args.g.Machine.Load("a", gvm_type.RefBool)), tt.wantRef)
			} else if err == nil && tt.args.g.PC != tt.wantPC {
				t.Errorf("Exec() got = %v, want %v", tt.args.g.PC, tt.wantPC)
			}
		})
	}
}
