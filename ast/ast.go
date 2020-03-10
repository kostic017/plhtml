package ast

import "../scanner"

type TokenType = scanner.TokenType

type AstNode interface {
	ToString() string
}

type StatementNode interface {
	AstNode
}

type ExpressionNode interface {
	AstNode
}
