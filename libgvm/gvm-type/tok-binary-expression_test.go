package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"reflect"
	"testing"
)

func TestBinaryExpression_Eval(t *testing.T) {
	type fields struct {
		Type  abstraction.RefType
		Sign  SignType
		Left  abstraction.VTok
		Right abstraction.VTok
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
		err     error
	}{
		{name: "add", fields: fields{Type: RefUint64, Sign: SignAdd, Left: Uint64(1), Right: Uint64(2)}, args: args{
			nil,
		}, want: Uint64(3)},
		{name: "sub", fields: fields{Type: RefUint64, Sign: SignSub, Left: Uint64(2), Right: Uint64(1)}, args: args{
			nil,
		}, want: Uint64(1)},
		{name: "type error", fields: fields{Type: RefInt64, Sign: SignSub, Left: Uint64(2), Right: Uint64(1)}, args: args{
			nil,
		}, want: Uint64(1), wantErr: true, err: expressionTypeError(RefInt64, RefUint64)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BinaryExpression{
				Type:  tt.fields.Type,
				Sign:  tt.fields.Sign,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			got, err := b.Eval(tt.args.g)
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
