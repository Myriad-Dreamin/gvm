package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"reflect"
	"testing"
)

func TestLocalVariable_Eval(t *testing.T) {
	type fields struct {
		Name string
		Type abstraction.RefType
	}
	type args struct {
		g *abstraction.ExecCtx
	}

	ctx := &abstraction.ExecCtx{
		This: make(abstraction.Locals),
	}

	ctx.This["b"] = Bool(true)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    abstraction.Ref
		wantErr bool
		err     error
	}{
		{name: "getCtx-param-a", fields: fields{Name: "a", Type: RefBool}, args: args{ctx}, want: Bool(true)},
		{name: "getCtx-param-a-type-error", fields: fields{Name: "a", Type: RefUint64},
			args: args{ctx}, wantErr: true, err: expressionTypeError(RefUint64, RefBool)},
		{name: "getCtx-param-b", fields: fields{Name: "b", Type: RefUnknown}, args: args{ctx}, want: Undefined},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LocalVariable{
				Name: tt.fields.Name,
				Type: tt.fields.Type,
			}
			got, err := l.Eval(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				if tt.err.Error() != err.Error() {
					t.Errorf("Eval.err = %v, want = %v", err.Error(), tt.err.Error())
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eval() got = %v, want %v", got, tt.want)
			}
		})
	}
}
