package libgvm

import "errors"

var OutOfRange = errors.New("gvm stopped")

var StopUnderFlow = errors.New("depth underflow")
