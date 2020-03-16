package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"reflect"
	"testing"
)

func TestUnaryExpression_Eval(t *testing.T) {
	type fields struct {
		Type abstraction.RefType
		Sign SignType
		Left abstraction.VTok
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
		{name: "got", fields: fields{Sign: SignLNot, Type: RefBool, Left: Bool(true)},
			args: args{&abstraction.ExecCtx{}}, want: Bool(false)},
		{name: "got", fields: fields{Sign: SignLNot, Type: RefBool, Left: Bool(false)},
			args: args{&abstraction.ExecCtx{}}, want: Bool(true)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UnaryExpression{
				Type: tt.fields.Type,
				Sign: tt.fields.Sign,
				Left: tt.fields.Left,
			}
			got, err := u.Eval(tt.args.g)
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
