package libgvm

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
)

type MachineBase struct{}

func (g *MachineBase) CreateRef(t abstraction.RefType, v interface{}) abstraction.Ref {
	return gvm_type.CreateRef(t, v)
}

func (g *MachineBase) DecodeRef(t abstraction.RefType, r []byte) (abstraction.Ref, error) {
	return gvm_type.DecodeRef(t, r)
}
