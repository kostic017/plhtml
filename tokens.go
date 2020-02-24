package main

type TokenType string

type Token struct {
	Type    TokenType
	IntVal  int
	BoolVal bool
	RealVal float64
	StrVal  string
}

const (
	TokEOF TokenType = ""

	TokDoctype TokenType = "TokDoctype"
	TokLang    TokenType = "TokLang"
	TokHTML    TokenType = "TokHTML"
	TokHead    TokenType = "TokHead"
	TokTitle   TokenType = "TokTitle"
	TokBody    TokenType = "TokBody"
	TokMain    TokenType = "TokMain"
	TokVar     TokenType = "TokVar"
	TokClass   TokenType = "TokClass"
	TokOutput  TokenType = "TokOutput"
	TokInput   TokenType = "TokInput"
	TokName    TokenType = "TokName"
	TokData    TokenType = "TokData"
	TokValue   TokenType = "TokValue"
	TokDiv     TokenType = "TokDiv"
	TokIf      TokenType = "TokIf"
	TokWhile   TokenType = "TokWhile"

	TokAddOp TokenType = "TokAddOp"
	TokSubOp TokenType = "TokSubOp"
	TokMulOp TokenType = "TokMulOp"
	TokDivOp TokenType = "TokDivOp"
	TokLtOp  TokenType = "TokLtOp"
	TokGtOp  TokenType = "TokGtOp"
	TokLeqOp TokenType = "TokLeqOp"
	TokGeqOp TokenType = "TokGeqOp"
	TokEqOp  TokenType = "TokEqOp"
	TokNeqOp TokenType = "TokNeqOp"
	TokNotOp TokenType = "TokNotOp"

	TokIntType    TokenType = "TokIntType"
	TokRealType   TokenType = "TokRealType"
	TokBoolType   TokenType = "TokBoolType"
	TokStringType TokenType = "TokStringType"

	TokIntConst    TokenType = "TokIntConst"
	TokRealConst   TokenType = "TokRealConst"
	TokBoolConst   TokenType = "TokBoolConst"
	TokStringConst TokenType = "TokStringConst"

	TokIdentifier TokenType = "TokIdentifier"
)
