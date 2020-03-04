package parser

type AstNode interface {
}

type StatementNode interface {
}

type ExpressionNode interface {
}

type ProgramNode struct {
	Title StringConstNode
	Body  ProgramBodyNode
}

type ProgramBodyNode struct {
	MainFunc MainFuncNode
}

type MainFuncNode struct {
	Statements []StatementNode
}

type VarDeclNode struct {
	Identifier IdentifierNode
	Type       TokenType
}

type VarAssignNode struct {
	Identifier IdentifierNode
	Value      ExpressionNode
}

type IdentifierNode struct {
	Name string
}

type StringConstNode struct {
	Value string
}

type IntConstNode struct {
	Value int
}

type RealConstNode struct {
	Value float64
}

type BoolConstNode struct {
	Value bool
}

type WriteStmtNode struct {
	Value ExpressionNode
}

type ReadStmtNode struct {
	Identifier IdentifierNode
}

type WhileStmtNode struct {
	Condition  ExpressionNode
	Statements []StatementNode
}

type BinaryOpExprNode struct {
	Value1   ExpressionNode
	Value2   ExpressionNode
	Operator TokenType
}

type UnaryExprNode struct {
	Operator TokenType
	Value    ExpressionNode
}
