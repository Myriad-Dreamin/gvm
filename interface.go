package gvm

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/gvm/libgvm"
)

type TokType = abstraction.TokType
type RefType = abstraction.RefType

type VTok = abstraction.VTok
type Ref = abstraction.Ref

type ExecCtx = abstraction.ExecCtx
type Instruction = abstraction.Instruction

type Machine = abstraction.Machine
type GVM = libgvm.GVM
type GVMeX = libgvm.GVMeX

func NewGVM() (*GVMeX, error) {
	return libgvm.NewGVM()
}

func Wrap(g Machine) *GVM {

	return &GVM{Machine: g}
}
