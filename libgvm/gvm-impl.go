package libgvm

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
)

type GVM struct {
	abstraction.Machine
}

func (i *GVM) Continue() error {
	return Continue(i.Machine, sugar.HandlerError0)
}

func (i *GVM) Step() error {
	return Step(i.Machine)
}

func (i *GVM) Run(fn string) error {
	return Run(i.Machine, fn, sugar.HandlerError0)
}

func (i *GVM) RunWithHandle(fn string, reportSaveError func(err error)) error {
	return Run(i.Machine, fn, reportSaveError)
}
