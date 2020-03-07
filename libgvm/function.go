package libgvm

import "github.com/Myriad-Dreamin/gvm/internal/abstraction"

type FunctionImpl []abstraction.Instruction

func (f FunctionImpl) Fetch(pc uint64) (abstraction.Instruction, error) {
	return f[pc], nil
}

func (f FunctionImpl) Len() int {
	return len(f)
}
