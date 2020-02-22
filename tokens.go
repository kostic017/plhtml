package main

type token string

const (
	tokEOF token = ""

	tokDoctype token = "tokDoctype"
	tokLang    token = "tokLang"
	tokHTML    token = "tokHTML"
	tokHead    token = "tokHead"
	tokTitle   token = "tokTitle"
	tokBody    token = "tokBody"
	tokMain    token = "tokMain"
	tokVar     token = "tokVar"
	tokClass   token = "tokClass"
	tokOutput  token = "tokOutput"
	tokInput   token = "tokInput"
	tokName    token = "tokName"
	tokData    token = "tokData"
	tokValue   token = "tokValue"
	tokIf      token = "tokIf"
	tokWhile   token = "tokWhile"
	tokFor     token = "tokFor"
	tokEqual   token = "tokEqual"

	tokIdentifier token = "tokIdentifier"

	tokIntType    token = "tokIntType"
	tokRealType   token = "tokRealType"
	tokBoolType   token = "tokBoolType"
	tokStringType token = "tokStringType"

	tokIntConst    token = "tokIntConst"
	tokRealConst   token = "tokRealConst"
	tokBoolConst   token = "tokBoolConst"
	tokStringConst token = "tokStringConst"
)
