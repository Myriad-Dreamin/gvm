package gvm_instruction

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"testing"
)

func TestConditionGoto_Exec(t *testing.T) {
	type fields struct {
		Goto      Goto
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
		wantPC  uint64
	}{
		{
			name: "goto",
			fields: fields{
				Goto: Goto{
					Index: 10,
				},
				Condition: gvm_type.Bool(true),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine: sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX),
					PC:      1,
					This:    make(abstraction.Locals),
				},
			},
			wantPC: 10,
		},
		{
			name: "not-goto-and-pc++",
			fields: fields{
				Goto: Goto{
					Index: 10,
				},
				Condition: gvm_type.Bool(false),
			},
			args: args{
				&abstraction.ExecCtx{
					Machine: sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX),
					PC:      2,
					This:    make(abstraction.Locals),
				},
			},
			wantPC: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &ConditionGoto{
				Goto:      tt.fields.Goto,
				Condition: tt.fields.Condition,
			}
			if err := inst.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil && tt.args.g.PC != tt.wantPC {
				t.Errorf("Exec() got = %v, want %v", tt.args.g.PC, tt.wantPC)
			}
		})
	}
}

func TestGoto_Exec(t *testing.T) {
	type fields struct {
		Index uint64
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
	}{
		{
			name: "goto",
			fields: fields{
				Index: 10,
			},
			args: args{
				&abstraction.ExecCtx{
					Machine: sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX),
					PC:      1,
					This:    make(abstraction.Locals),
				},
			},
			wantPC: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &Goto{
				Index: tt.fields.Index,
			}
			if err := inst.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil && tt.args.g.PC != tt.wantPC {
				t.Errorf("Exec() got = %v, want %v", tt.args.g.PC, tt.wantPC)
			}
		})
	}
}
