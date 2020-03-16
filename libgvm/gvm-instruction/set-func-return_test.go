package gvm_instruction

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"testing"
)

func TestConditionSetFuncReturn_Exec(t *testing.T) {
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &ConditionSetFuncReturn{
				SetFuncReturn: tt.fields.SetFuncReturn,
				Condition:     tt.fields.Condition,
			}
			if err := inst.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetFuncReturn_Exec(t *testing.T) {
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := SetFuncReturn{
				Target:          tt.fields.Target,
				RightExpression: tt.fields.RightExpression,
			}
			if err := G.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
