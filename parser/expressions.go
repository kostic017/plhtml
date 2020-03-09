package parser

import (
	"fmt"

	"../ast"
	"../scanner"
)

func (parser *Parser) parseExpr() ast.ExpressionNode {
	parser.logger.Debug("=BEG= Expression")
	lhs := parser.parsePrimaryExpr()
	expr := parser.parseBinOpRhs(lhs, 0)
	parser.logger.Debug("=END= Expression")
	return expr
}

func (parser *Parser) parseBinOpRhs(lhs ast.ExpressionNode, minPrec int) ast.ExpressionNode {

	for {
		prec := parser.peekBinOpPrec()

		if prec < minPrec {
			return lhs
		}

		binop := parser.next()
		rhs := parser.parsePrimaryExpr()

		// lhs binop rhs next_binop ...
		nextPrec := parser.peekBinOpPrec()

		if prec < nextPrec {
			// lhs binop (rhs next_binop ...)
			rhs = parser.parseBinOpRhs(rhs, prec+1)
		}

		// (lhs binop rhs) next_binop ...
		lhs = ast.BinaryOpExprNode{Expr1: lhs, Expr2: rhs, Operator: binop.Type}
	}

}

func (parser *Parser) peekBinOpPrec() int {
	prec, ok := parser.binOpsPrec[parser.peek().Type]
	if !ok {
		return -1 // if not binop
	}
	return prec
}

func (parser *Parser) parsePrimaryExpr() ast.ExpressionNode {

	switch parser.peek().Type {
	case scanner.TokIntConst:
		return parser.parseIntConst()
	case scanner.TokRealConst:
		return parser.parseRealConst()
	case scanner.TokBoolConst:
		return parser.parseBoolConst()
	case scanner.TokStringConst:
		return parser.parseStringConst()
	case scanner.TokIdentifier:
		return parser.parseIdentifier()
	case TokenType('('):
		return parser.parseParenExpr()
	case TokenType('+'), TokenType('-'), TokenType('!'):
		return parser.parseUnaryExpr()
	}

	panic(fmt.Sprintf("Invalid primary expression '%s'.", string(parser.peek().Type)))
}

func (parser *Parser) parseIdentifier() ast.IdentifierNode {
	parser.logger.Debug("=BEG= Identifier")
	parser.expect(scanner.TokIdentifier)
	parser.logger.Debug("=END= Identifier")
	return ast.IdentifierNode{Name: parser.current().StrVal}
}

func (parser *Parser) parseIntConst() ast.IntConstNode {
	parser.logger.Debug("=BEG= Integer Constant")
	parser.expect(scanner.TokIntConst)
	parser.logger.Debug("=END= Integer Constant")
	return ast.IntConstNode{Value: parser.current().IntVal}
}

func (parser *Parser) parseRealConst() ast.RealConstNode {
	parser.logger.Debug("=BEG= Real Constant")
	parser.expect(scanner.TokRealConst)
	parser.logger.Debug("=END= Real Constant")
	return ast.RealConstNode{Value: parser.current().RealVal}
}

func (parser *Parser) parseBoolConst() ast.BoolConstNode {
	parser.logger.Debug("=BEG= Boolean Constant")
	parser.expect(scanner.TokBoolConst)
	parser.logger.Debug("=END= Boolean Constant")
	return ast.BoolConstNode{Value: parser.current().BoolVal}
}

func (parser *Parser) parseStringConst() ast.StringConstNode {
	parser.logger.Debug("=BEG= String Constant")
	parser.expect(scanner.TokStringConst)
	parser.logger.Debug("=END= String Constant")
	return ast.StringConstNode{Value: parser.current().StrVal}
}

func (parser *Parser) parseParenExpr() ast.ExpressionNode {
	parser.logger.Debug("=BEG= ()")
	parser.expect(TokenType('('))
	expr := parser.parseExpr()
	parser.expect(TokenType(')'))
	parser.logger.Debug("=END= ()")
	return expr
}

func (parser *Parser) parseUnaryExpr() ast.UnaryExprNode {
	parser.logger.Debug("=BEG= Unary")
	op := parser.expect(TokenType('+'), TokenType('-'), TokenType('!'))
	expr := parser.parsePrimaryExpr()
	parser.logger.Debug("=END= Unary")
	return ast.UnaryExprNode{Operator: op, Expr: expr}
}
