package ast

import (
	"strings"

	"../token"
)

type TokenType = token.Type

type Node interface {
	Accept(v Visitor) interface{}
}

type StatementNode interface {
	Node
	ToString(lvl int) string
}

type ExpressionNode interface {
	Node
	ToString() string
}

type Visitor interface {
	VisitBinaryOpExpr(node BinaryOpExprNode) interface{}
	VisitBoolConst(node BoolConstNode) interface{}
	VisitControlFlowStmt(node ControlFlowStmtNode) interface{}
	VisitIdentifier(node IdentifierNode) interface{}
	VisitIntConst(node IntConstNode) interface{}
	VisitMainFunc(node MainFuncNode) interface{}
	VisitProgram(node ProgramNode) interface{}
	VisitProgramBody(node ProgramBodyNode) interface{}
	VisitReadStmt(node ReadStmtNode) interface{}
	VisitRealConst(node RealConstNode) interface{}
	VisitStringConst(node StringConstNode) interface{}
	VisitUnaryExpr(node UnaryExprNode) interface{}
	VisitVarAssign(node VarAssignNode) interface{}
	VisitVarDecl(node VarDeclNode) interface{}
	VisitWriteStmt(node WriteStmtNode) interface{}
}

func ident(lvl int, str string) string {
	return strings.Repeat(" ", lvl*4) + str
}
