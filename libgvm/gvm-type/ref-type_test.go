package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"math/big"
	"reflect"
	"testing"
)

func TestGoodEDecode(t *testing.T) {
	tests := []struct {
		name     string
		variable abstraction.Ref
	}{
		{"uint8", Uint8(1)},
		{"uint16", Uint16(1)},
		{"uint32", Uint32(1)},
		{"uint64", Uint64(1)},

		{"int8", Int8(1)},
		{"int16", Int16(1)},
		{"int32", Int32(1)},
		{"int64", Int64(1)},

		{"uint128", (*Uint128)(big.NewInt(128))},
		{"uint256", (*Uint256)(big.NewInt(256))},
		{"int128", (*Int128)(big.NewInt(128))},
		{"int256", (*Int256)(big.NewInt(256))},

		{"bool", Bool(true)},
		{"string", String("123")},
		{"bytes", Bytes("123")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := sugar.HandlerError(
				tt.variable.(interface {
					Decode([]byte) (abstraction.Ref, error)
				}).Decode(sugar.HandlerError(tt.variable.Encode()).([]byte)))
			if !assert.EqualValues(t, tt.variable, v) {
				t.Errorf("got = %v, want = %v", v, tt.variable)
			}
		})
	}
}

func TestCreateRef(t *testing.T) {
	type args struct {
		t abstraction.RefType
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want abstraction.Ref
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateRef(tt.args.t, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeRef(t *testing.T) {
	type args struct {
		t abstraction.RefType
		r []byte
	}
	tests := []struct {
		name    string
		args    args
		want    abstraction.Ref
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeRef(tt.args.t, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeRef() got = %v, want %v", got, tt.want)
			}
		})
	}
}
