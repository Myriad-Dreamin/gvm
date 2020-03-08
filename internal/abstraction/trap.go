package abstraction

type Trap interface {
	DoTrap(g *ExecCtx) error
}
