package libgvm

import "github.com/Myriad-Dreamin/gvm/internal/abstraction"

type MachineBase struct{}

func (g *MachineBase) CreateRef(t abstraction.RefType, v interface{}) abstraction.Ref {
	return CreateRef(t, v)
}

func (g *MachineBase) DecodeRef(t abstraction.RefType, r []byte) (abstraction.Ref, error) {
	return DecodeRef(t, r)
}