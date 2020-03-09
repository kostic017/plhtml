package ast

import "../scanner"

type TokenType = scanner.TokenType

type AstNode interface {
	Print()
}

type StatementNode interface {
	AstNode
}

type ExpressionNode interface {
	AstNode
}
