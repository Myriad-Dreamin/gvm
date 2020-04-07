package libgvm

import (
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
)

func ProcessRunException(g *abstraction.ExecCtx, err error) error {

	if err == OutOfRange {
		err = popFrame(g)
	} else if trap, ok := err.(abstraction.Trap); ok {
		err = trap.DoTrap(g)
	}
	return err
}

func SaveFrame(g *abstraction.ExecCtx) error {
	return saveFrame(g)
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

func _Continue(g *abstraction.ExecCtx, err error, reportSaveError func(err error)) error {
	for err == nil {
		for ; err == nil; err = Iter(g) {
		}
		err = ProcessRunException(g, err)
	}
	if err != StopUnderFlow {
		reportSaveError(saveFrame(g))
	}

	return err
}

func Step(g abstraction.Machine) error {

	c := abstraction.ExecCtx{Machine: g}
	if err := loadFrameFromDisk(&c); err != nil {
		return err
	}

	return Iter(&c)
}

func Continue(g abstraction.Machine, reportSaveError func(err error)) error {
	c := abstraction.ExecCtx{Machine: g}
	if err := loadFrameFromDisk(&c); err != nil {
		return err
	}

	return _Continue(&c, nil, reportSaveError)
}

//trapCallFunc
func Run(g abstraction.Machine, fn string, reportSaveError func(err error)) (err error) {

	var c = abstraction.ExecCtx{Machine: g, Depth: 0, This: make(abstraction.Locals)}
	err = PushFrame(&c, fn)
	return _Continue(&c, err, reportSaveError)
}

func validate(g *abstraction.ExecCtx) error {
	if g.PC >= uint64(g.Len()) {
		return OutOfRange
	}
	return nil
}
