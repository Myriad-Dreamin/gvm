package libgvm

type SignType = uint16

const (
	SignUnknown SignType = iota

	SignEQ
	SignLE
	SignLT
	SignGE
	SignGT
	SignLNot
	Sign// Logic And
	SignLAnd
	SignLOr

	SignADD // +
	SignSUB // -
	SignMUL // *
	SignQUO // /
	SignREM // %

	SignAND     // &
	SignOR      // |
	SignXOR     // ^
	SignSHL     // <<
	SignSHR     // >>
	SignAND_NOT // &^

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

	SignLogicL = SignEQ
	SignLogicR = SignLOr + 1
)

func IsLogic(s SignType) bool {
	return SignLogicL <= s && s < SignLogicR
}
