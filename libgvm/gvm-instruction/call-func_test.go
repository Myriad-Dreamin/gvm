package gvm_instruction

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"testing"
)

func TestConditionCallFunc_Exec(t *testing.T) {
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := ConditionCallFunc{
				CallFunc:  tt.fields.CallFunc,
				Condition: tt.fields.Condition,
			}
			if err := inst.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
