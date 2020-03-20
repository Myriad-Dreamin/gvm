package gvm

import (
	"github.com/Myriad-Dreamin/gvm/libgvm"
)

type MemMachine = libgvm.Mem
type MachineBase = libgvm.MachineBase


type GVM = libgvm.GVM
type GVMeX = libgvm.GVMeX

func NewGVM() (*GVMeX, error) {
	return libgvm.NewGVM()
}

func Wrap(g Machine) *GVM {

	return &GVM{Machine: g}
}
