package ast

import (
	"strings"

	"../token"
)

type TokenType = token.Type

type Node interface {
}

type StatementNode interface {
	Node
	ToString(lvl int) string
	Accept(v Visitor)
}

type ExpressionNode interface {
	Node
	ToString() string
	Accept(v Visitor) interface{}
}

type Visitor interface {
	VisitBinaryOpExpr(node BinaryOpExprNode) interface{}
	VisitBoolConst(node BoolConstNode) interface{}
	VisitControlFlowStmt(node ControlFlowStmtNode)
	VisitIdentifier(node IdentifierNode) interface{}
	VisitIntConst(node IntConstNode) interface{}
	VisitMainFunc(node MainFuncNode)
	VisitProgram(node ProgramNode)
	VisitProgramBody(node ProgramBodyNode)
	VisitReadStmt(node ReadStmtNode)
	VisitRealConst(node RealConstNode) interface{}
	VisitStringConst(node StringConstNode) interface{}
	VisitUnaryExpr(node UnaryExprNode) interface{}
	VisitVarAssign(node VarAssignNode)
	VisitVarDecl(node VarDeclNode)
	VisitWriteStmt(node WriteStmtNode)
}

func ident(lvl int, str string) string {
	return strings.Repeat(" ", lvl*4) + str
}
