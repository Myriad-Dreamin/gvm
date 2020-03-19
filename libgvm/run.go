package libgvm

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
)

func Step(g abstraction.Machine) error {

	c := abstraction.ExecCtx{Machine: g}
	if err := loadFrameFromDisk(&c); err != nil {
		return err
	}

	return Iter(&c)
}

func Continue(g abstraction.Machine) error {
	c := abstraction.ExecCtx{Machine: g}
	if err := loadFrameFromDisk(&c); err != nil {
		return err
	}

	return _Continue(&c)
}

func _Continue(g *abstraction.ExecCtx) (err error) {
	for ; err == nil; err = Iter(g) {
	}

	return err
}

//trapCallFunc
func Run(g abstraction.Machine, fn string, reportSaveError func(err error)) (err error) {

	var c = &abstraction.ExecCtx{Machine: g, Depth: 0, This: make(abstraction.Locals)}
	err = PushFrame(c, fn)
	for err == nil {
		err = _Continue(c)

		if err == OutOfRange {
			err = popFrame(c)
		} else if trap, ok := err.(abstraction.Trap); ok {
			err = trap.DoTrap(c)
		}
	}
	if err != StopUnderFlow {
		reportSaveError(saveFrame(c))
	}
	return err
}

func Iter(g *abstraction.ExecCtx) (err error) {
	if err = validate(g); err != nil {
		return
	}
	inst, err := g.Fetch(g.PC)
	if err != nil {
		return err
	}
	return inst.Exec(g)
}

func validate(g *abstraction.ExecCtx) error {
	if g.PC >= uint64(g.Len()) {
		return OutOfRange
	}
	return nil
}
