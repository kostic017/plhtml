package scanner

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	EOF TokenType = ""

	Doctype TokenType = "Doctype"
	Lang    TokenType = "Lang"
	HTML    TokenType = "HTML"
	Head    TokenType = "Head"
	Title   TokenType = "Title"
	Body    TokenType = "Body"
	Main    TokenType = "Main"
	Var     TokenType = "Var"
	Class   TokenType = "Class"
	Output  TokenType = "Output"
	Input   TokenType = "Input"
	Name    TokenType = "Name"
	Data    TokenType = "Data"
	Value   TokenType = "Value"
	Div     TokenType = "Div"
	If      TokenType = "If"
	While   TokenType = "While"

	AddOp TokenType = "AddOp"
	SubOp TokenType = "SubOp"
	MulOp TokenType = "MulOp"
	DivOp TokenType = "DivOp"
	LtOp  TokenType = "LtOp"
	GtOp  TokenType = "GtOp"
	LeqOp TokenType = "LeqOp"
	GeqOp TokenType = "GeqOp"
	EqOp  TokenType = "EqOp"
	NeqOp TokenType = "NeqOp"
	NotOp TokenType = "NotOp"

	IntType    TokenType = "IntType"
	RealType   TokenType = "RealType"
	BoolType   TokenType = "BoolType"
	StringType TokenType = "StringType"

	IntConst    TokenType = "IntConst"
	RealConst   TokenType = "RealConst"
	BoolConst   TokenType = "BoolConst"
	StringConst TokenType = "StringConst"

	Identifier TokenType = "Identifier"
)
