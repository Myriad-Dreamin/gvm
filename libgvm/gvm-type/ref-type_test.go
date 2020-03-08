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
		{"unknown", Undefined},
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
		{name: "uint8", args: args{t: RefUint8, v: uint8(1)}, want: Uint8(1)},
		{name: "uint16", args: args{t: RefUint16, v: uint16(1)}, want: Uint16(1)},
		{name: "uint32", args: args{t: RefUint32, v: uint32(1)}, want: Uint32(1)},
		{name: "uint64", args: args{t: RefUint64, v: uint64(1)}, want: Uint64(1)},
		{name: "uint128", args: args{t: RefUint128, v: big.NewInt(1)}, want: (*Uint128)(big.NewInt(1))},
		{name: "uint256", args: args{t: RefUint256, v: big.NewInt(1)}, want: (*Uint256)(big.NewInt(1))},
		{name: "int8", args: args{t: RefInt8, v: int8(1)}, want: Int8(1)},
		{name: "int16", args: args{t: RefInt16, v: int16(1)}, want: Int16(1)},
		{name: "int32", args: args{t: RefInt32, v: int32(1)}, want: Int32(1)},
		{name: "int64", args: args{t: RefInt64, v: int64(1)}, want: Int64(1)},
		{name: "int128", args: args{t: RefInt128, v: big.NewInt(1)}, want: (*Int128)(big.NewInt(1))},
		{name: "int256", args: args{t: RefInt256, v: big.NewInt(1)}, want: (*Int256)(big.NewInt(1))},
		{name: "bool", args: args{t: RefBool, v: true}, want: Bool(true)},
		{name: "bytes", args: args{t: RefBytes, v: []byte{1}}, want: Bytes([]byte{1})},
		{name: "string", args: args{t: RefString, v: ""}, want: String("")},
		{name: "bytes-nil", args: args{t: RefBytes, v: nil}, want: Bytes(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateRef(tt.args.t, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateRefWillPanic(t *testing.T) {
	type args struct {
		t abstraction.RefType
		v interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantPanic error
	}{
		{name: "uint128-nil", args: args{t: RefUint128, v: nil}, wantPanic: runtimeConvertNil(RefUint128)},
		{name: "uint256-nil", args: args{t: RefUint256, v: nil}, wantPanic: runtimeConvertNil(RefUint256)},
		{name: "int128-nil", args: args{t: RefInt128, v: nil}, wantPanic: runtimeConvertNil(RefInt128)},
		{name: "int256-nil", args: args{t: RefInt256, v: nil}, wantPanic: runtimeConvertNil(RefInt256)},
		{name: "unknown", args: args{t: RefUnknown, v: nil}, wantPanic: creatingUnknownReferenceType(RefUnknown)},
		{name: "some-type", args: args{t: 23333333, v: nil}, wantPanic: creatingUnknownReferenceType(23333333)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover().(error)
				if !assert.EqualValues(t, tt.wantPanic, err) {
					t.Errorf("CreateRef.panic() = %v, want.panic %v", err, tt.wantPanic)
				}
			}()

			_ = CreateRef(tt.args.t, tt.args.v)
		})
	}
}

func TestDecodeRef(t *testing.T) {
	type args struct {
		t abstraction.RefType
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    abstraction.Ref
		wantErr bool
	}{
		{name: "uint8", args: args{t: RefUint8, v: uint8(1)}, want: Uint8(1)},
		{name: "uint16", args: args{t: RefUint16, v: uint16(1)}, want: Uint16(1)},
		{name: "uint32", args: args{t: RefUint32, v: uint32(1)}, want: Uint32(1)},
		{name: "uint64", args: args{t: RefUint64, v: uint64(1)}, want: Uint64(1)},
		{name: "uint128", args: args{t: RefUint128, v: big.NewInt(1)}, want: (*Uint128)(big.NewInt(1))},
		{name: "uint256", args: args{t: RefUint256, v: big.NewInt(1)}, want: (*Uint256)(big.NewInt(1))},
		{name: "int8", args: args{t: RefInt8, v: int8(1)}, want: Int8(1)},
		{name: "int16", args: args{t: RefInt16, v: int16(1)}, want: Int16(1)},
		{name: "int32", args: args{t: RefInt32, v: int32(1)}, want: Int32(1)},
		{name: "int64", args: args{t: RefInt64, v: int64(1)}, want: Int64(1)},
		{name: "int128", args: args{t: RefInt128, v: big.NewInt(1)}, want: (*Int128)(big.NewInt(1))},
		{name: "int256", args: args{t: RefInt256, v: big.NewInt(1)}, want: (*Int256)(big.NewInt(1))},
		{name: "bool", args: args{t: RefBool, v: true}, want: Bool(true)},
		{name: "bytes", args: args{t: RefBytes, v: []byte{1}}, want: Bytes([]byte{1})},
		{name: "string", args: args{t: RefString, v: ""}, want: String("")},
		{name: "bytes-nil", args: args{t: RefBytes, v: nil}, want: Bytes(nil)},

		// must panic
		//{name: "uint128-nil", args: args{t: RefUint128, v: nil}, want: (*Uint128)(big.NewInt(0).SetBytes(nil))},
		//{name: "uint256-nil", args: args{t: RefUint256, v: nil}, want: (*Uint256)(big.NewInt(0).SetBytes(nil))},
		//{name: "int128-nil", args: args{t: RefInt128, v: nil}, want: (*Int128)(big.NewInt(0).SetBytes(nil))},
		//{name: "int256-nil", args: args{t: RefInt256, v: nil}, want: (*Int256)(big.NewInt(0).SetBytes(nil))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CreateRef(tt.args.t, tt.args.v)
			got, err := DecodeRef(tt.args.t, sugar.HandlerError(r.Encode()).([]byte))
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(assert.EqualValues(t, tt.want, got) && assert.EqualValues(t, r, got)) {
				t.Errorf("DecodeRef() got = %v, want %v", got, tt.want)
			}
		})
	}
}
