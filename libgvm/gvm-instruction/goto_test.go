package gvm_instruction

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &ConditionGoto{
				Goto:      tt.fields.Goto,
				Condition: tt.fields.Condition,
			}
			if err := inst.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &Goto{
				Index: tt.fields.Index,
			}
			if err := inst.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
