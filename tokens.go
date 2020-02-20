package main

import "fmt"

type token int

const (
	_              token = -iota
	tokDoctype     token = -iota
	tokLang        token = -iota
	tokHTML        token = -iota
	tokHead        token = -iota
	tokTitle       token = -iota
	tokBody        token = -iota
	tokMain        token = -iota
	tokVar         token = -iota
	tokClass       token = -iota
	tokOutput      token = -iota
	tokInput       token = -iota
	tokName        token = -iota
	tokData        token = -iota
	tokValue       token = -iota
	tokIntType     token = -iota
	tokRealType    token = -iota
	tokBoolType    token = -iota
	tokStringType  token = -iota
	tokIntConst    token = -iota
	tokRealConst   token = -iota
	tokBoolConst   token = -iota
	tokStringConst token = -iota
	tokAddOp       token = -iota
	tokSubOp       token = -iota
	tokMulOp       token = -iota
	tokDivOp       token = -iota
)

func (tok token) String() string {
	switch tok {
	case tokDoctype:
		return "tokDoctype"
	case tokLang:
		return "tokLang"
	case tokHTML:
		return "tokHTML"
	case tokHead:
		return "tokHead"
	case tokTitle:
		return "tokTitle"
	case tokBody:
		return "tokBody"
	case tokMain:
		return "tokMain"
	case tokVar:
		return "tokVar"
	case tokClass:
		return "tokClass"
	case tokOutput:
		return "tokOutput"
	case tokInput:
		return "tokInput"
	case tokName:
		return "tokName"
	case tokData:
		return "tokData"
	case tokValue:
		return "tokValue"
	case tokIntType:
		return "tokIntType"
	case tokRealType:
		return "tokRealType"
	case tokBoolType:
		return "tokBoolType"
	case tokStringType:
		return "tokStringType"
	case tokIntConst:
		return "tokIntConst"
	case tokRealConst:
		return "tokRealConst"
	case tokBoolConst:
		return "tokBoolConst"
	case tokStringConst:
		return "tokStringConst"
	case tokAddOp:
		return "tokAddOp"
	case tokSubOp:
		return "tokSubOp"
	case tokMulOp:
		return "tokMulOp"
	case tokDivOp:
		return "tokDivOp"
	default:
		return fmt.Sprintf("%d", int(tok))
	}
}
