package gvm_instruction

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConditionSetLocal_Exec(t *testing.T) {
	type fields struct {
		SetLocal  SetLocal
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
				SetLocal: SetLocal{
					Target:          "a",
					RightExpression: gvm_type.Bool(true),
				},
				Condition: gvm_type.Bool(true),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine:  sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX),
					PC:       1,
					This: make(abstraction.Locals),
				},
			},
			wantPC: 2,
			wantRef: gvm_type.Bool(true),
		},
		{
			name: "not-eval-and-pc++",
			fields: fields{
				SetLocal: SetLocal{
					Target:          "a",
					RightExpression: gvm_type.Bool(true),
				},
				Condition: gvm_type.Bool(false),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine:  sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX),
					PC:       2,
					This: make(abstraction.Locals),
				},
			},
			wantPC: 3,
			wantRef: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &ConditionSetLocal{
				SetLocal:  tt.fields.SetLocal,
				Condition: tt.fields.Condition,
			}
			if err := inst.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			} else if err ==nil && !assert.EqualValues(t, tt.wantRef, tt.args.g.This["a"]) {
				t.Errorf("Exec() got = %v, want %v", tt.args.g.This["a"], tt.wantRef)
			} else if err == nil && tt.args.g.PC != tt.wantPC {
				t.Errorf("Exec() got = %v, want %v", tt.args.g.PC, tt.wantPC)
			}
		})
	}
}

func TestSetLocal_Exec(t *testing.T) {
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
					This: make(abstraction.Locals),
				},
			},
			wantPC: 2,
			wantRef: gvm_type.Bool(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := SetLocal{
				Target:          tt.fields.Target,
				RightExpression: tt.fields.RightExpression,
			}
			if err := G.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			} else if err ==nil && !assert.EqualValues(t, tt.wantRef, tt.args.g.This["a"]) {
				t.Errorf("Exec() got = %v, want %v", tt.args.g.This["a"], tt.wantRef)
			} else if err == nil && tt.args.g.PC != tt.wantPC {
				t.Errorf("Exec() got = %v, want %v", tt.args.g.PC, tt.wantPC)
			}
		})
	}
}
