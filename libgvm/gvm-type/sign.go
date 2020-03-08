package gvm_type

type SignType = uint16

//noinspection GoUnusedConst
const (
	SignUnknown SignType = iota

	SignEQ   // ==
	SignNEQ  // !=
	SignLE   // <=
	SignLT   // >
	SignGE   // >=
	SignGT   // <
	SignLAnd // &&
	SignLOr  // ||
	SignLNot // !

	SignAdd // +
	SignSub // -
	SignMul // *
	SignQuo // /
	SignRem // %

	SignAnd    // &
	SignOr     // |
	SignXor    // ^
	SignNot    // ~
	SignSHL    // <<
	SignSHR    // >>
	SignAndNot // &^

	//Sign//ADD_ASSIGN // +=
	//Sign//SUB_ASSIGN // -=
	//Sign//MUL_ASSIGN // *=
	//Sign//QUO_ASSIGN // /=
	//Sign//REM_ASSIGN // %=
	//
	//Sign//AND_ASSIGN     // &=
	//Sign//OR_ASSIGN      // |=
	//Sign//XOR_ASSIGN     // ^=
	//Sign//SHL_ASSIGN     // <<=
	//Sign//SHR_ASSIGN     // >>=
	//Sign//AND_NOT_ASSIGN // &^=

	SignLength

	SignLogicL = SignEQ
	SignLogicR = SignLNot + 1

	SignArithmeticL = SignAdd
	SignArithmeticR = SignRem + 1

	SignBitwiseOpL = SignAnd
	SignBitwiseOpR = SignAndNot + 1
)

func IsLogic(s SignType) bool {
	return SignLogicL <= s && s < SignLogicR
}

func IsArithmetic(s SignType) bool {
	return SignArithmeticL <= s && s < SignArithmeticR
}

func IsBitwiseOp(s SignType) bool {
	return SignBitwiseOpL <= s && s < SignBitwiseOpR
}

func IsStandardSignType(t SignType) bool {
	return t > SignUnknown && t < SignLength
}
