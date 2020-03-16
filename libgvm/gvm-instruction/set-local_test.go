package gvm_instruction

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &ConditionSetLocal{
				SetLocal:  tt.fields.SetLocal,
				Condition: tt.fields.Condition,
			}
			if err := inst.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := SetLocal{
				Target:          tt.fields.Target,
				RightExpression: tt.fields.RightExpression,
			}
			if err := G.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
