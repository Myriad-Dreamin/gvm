package gvm_type_test

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"reflect"
	"testing"
)

func TestStateVariable_Eval(t *testing.T) {
	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	sugar.HandlerError0(g.Save("test", gvm_type.Bool(true)))
	type fields struct {
		Field string
		Type  abstraction.RefType
	}
	type args struct {
		g *abstraction.ExecCtx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    abstraction.Ref
		wantErr bool
	}{
		{name: "got", fields: fields{Field: "test", Type: gvm_type.RefBool}, args: args{&abstraction.ExecCtx{Machine: g}}, want: gvm_type.Bool(true)},
		{name: "undefined", fields: fields{Field: "test2", Type: gvm_type.RefUnknown}, args: args{&abstraction.ExecCtx{Machine: g}}, want: gvm_type.Undefined},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := gvm_type.StateVariable{
				Field: tt.fields.Field,
				Type:  tt.fields.Type,
			}
			got, err := s.Eval(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eval() got = %v, want %v", got, tt.want)
			}
		})
	}
}
