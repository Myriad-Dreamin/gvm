package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"reflect"
	"testing"
)

func TestFuncReturnParam_Eval(t *testing.T) {
	type fields struct {
		T abstraction.RefType
		K int
	}
	type args struct {
		g *abstraction.ExecCtx
	}

	funcCtx := &abstraction.ExecCtx{
		Parent: make(abstraction.Locals),
		This:   make(abstraction.Locals),
	}

	SetReturnField(funcCtx, 0, "a")
	funcCtx.Parent["a"] = Uint64(1)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    abstraction.Ref
		wantErr bool
		err     error
	}{
		{name: "getFuncCtx-return0", fields: fields{T: RefUint64, K: 0}, args: args{funcCtx}, want: Uint64(1)},
		{name: "getFuncCtx-return0-type-error", fields: fields{T: RefInt64, K: 0},
			args: args{funcCtx}, wantErr: true, err: expressionTypeError(RefInt64, RefUint64)},
		{name: "getFuncCtx-return2", fields: fields{T: RefUnknown, K: 2}, args: args{funcCtx}, want: Undefined},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FuncReturnParam{
				T: tt.fields.T,
				K: tt.fields.K,
			}
			got, err := f.Eval(tt.args.g)
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
