package parser

import (
	"fmt"

	"../scanner"
)

func (parser *Parser) parseExpr() ExpressionNode {
	parser.logger.Debug("=BEG= Expression")
	lhs := parser.parsePrimaryExpr()
	expr := parser.parseBinOpRhs(lhs, 0)
	parser.logger.Debug("=END= Expression")
	return expr
}

func (parser *Parser) parseBinOpRhs(lhs ExpressionNode, minPrec int) ExpressionNode {

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
		lhs = BinaryOpExprNode{Value1: lhs, Value2: rhs, Operator: binop.Type}
	}

}

func (parser *Parser) peekBinOpPrec() int {
	prec, ok := parser.binOpsPrec[parser.peek().Type]
	if !ok {
		return -1 // if not binop
	}
	return prec
}

func (parser *Parser) parsePrimaryExpr() ExpressionNode {

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

func (parser *Parser) parseIdentifier() IdentifierNode {
	parser.logger.Debug("=BEG= Identifier")
	parser.expect(scanner.TokIdentifier)
	parser.logger.Debug("=END= Identifier")
	return IdentifierNode{Name: parser.current().StrVal}
}

func (parser *Parser) parseIntConst() IntConstNode {
	parser.logger.Debug("=BEG= Integer Constant")
	parser.expect(scanner.TokIntConst)
	parser.logger.Debug("=END= Integer Constant")
	return IntConstNode{Value: parser.current().IntVal}
}

func (parser *Parser) parseRealConst() RealConstNode {
	parser.logger.Debug("=BEG= Real Constant")
	parser.expect(scanner.TokRealConst)
	parser.logger.Debug("=END= Real Constant")
	return RealConstNode{Value: parser.current().RealVal}
}

func (parser *Parser) parseBoolConst() BoolConstNode {
	parser.logger.Debug("=BEG= Boolean Constant")
	parser.expect(scanner.TokBoolConst)
	parser.logger.Debug("=END= Boolean Constant")
	return BoolConstNode{Value: parser.current().BoolVal}
}

func (parser *Parser) parseStringConst() StringConstNode {
	parser.logger.Debug("=BEG= String Constant")
	parser.expect(scanner.TokStringConst)
	parser.logger.Debug("=END= String Constant")
	return StringConstNode{Value: parser.current().StrVal}
}

func (parser *Parser) parseParenExpr() ExpressionNode {
	parser.logger.Debug("=BEG= ()")
	parser.expect(TokenType('('))
	expr := parser.parseExpr()
	parser.expect(TokenType(')'))
	parser.logger.Debug("=END= ()")
	return expr
}

func (parser *Parser) parseUnaryExpr() UnaryExprNode {
	parser.logger.Debug("=BEG= Unary")
	op := parser.expect(TokenType('+'), TokenType('-'), TokenType('!'))
	expr := parser.parsePrimaryExpr()
	parser.logger.Debug("=END= Unary")
	return UnaryExprNode{Operator: op, Value: expr}
}
