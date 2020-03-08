package gvm_type

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
)

const (
	RefUnknown abstraction.RefType = iota
	RefBytes
	RefString
	RefUint8
	RefUint16 // 4
	RefUint32
	RefUint64
	RefUint128
	RefUint256
	RefInt8 // 9
	RefInt16
	RefInt32
	RefInt64
	RefInt128
	RefInt256 // 14

	RefBool // 15

	// Slice
	Length
)

func IsStandardRefType(t abstraction.RefType) bool {
	return t > RefUnknown && t < Length
}
