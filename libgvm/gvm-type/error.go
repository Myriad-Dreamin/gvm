package gvm_type

import (
	"fmt"
	"github.com/Myriad-Dreamin/gvm/internal/abstraction"
)

func convertError(fr, to abstraction.Ref) (err error) {
	return fmt.Errorf("cant convert %s and %s to the same type", ExplainGVMType(fr.GetGVMType()), ExplainGVMType(to.GetGVMType()))
}

func convertUnsignedError(k abstraction.Ref) error {
	return fmt.Errorf("cant convert %s to the unsigned type", ExplainGVMType(k.GetGVMType()))
}

func invalidTypeError(k abstraction.Ref) error {
	return fmt.Errorf("invalid type: %v", ExplainGVMType(k.GetGVMType()))
}

func expressionTypeError(want abstraction.RefType, got abstraction.RefType) error {
	return fmt.Errorf("expression want type %s, got %s", ExplainGVMType(want), ExplainGVMType(got))
}
