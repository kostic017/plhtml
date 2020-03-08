package parser

import "../scanner"

func (parser *Parser) parseExpr() ExpressionNode {
	lhs := parser.parsePrimaryExpr()
	return parser.parseBinOpRhs(lhs, 0)
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

	panic("Invalid primary expression.")
}

func (parser *Parser) parseIdentifier() IdentifierNode {
	parser.expect(scanner.TokIdentifier)
	return IdentifierNode{Name: parser.current().StrVal}
}

func (parser *Parser) parseIntConst() IntConstNode {
	parser.expect(scanner.TokIntConst)
	return IntConstNode{Value: parser.current().IntVal}
}

func (parser *Parser) parseRealConst() RealConstNode {
	parser.expect(scanner.TokRealConst)
	return RealConstNode{Value: parser.current().RealVal}
}

func (parser *Parser) parseBoolConst() BoolConstNode {
	parser.expect(scanner.TokBoolConst)
	return BoolConstNode{Value: parser.current().BoolVal}
}

func (parser *Parser) parseStringConst() StringConstNode {
	parser.expect(scanner.TokStringConst)
	return StringConstNode{Value: parser.current().StrVal}
}

func (parser *Parser) parseParenExpr() ExpressionNode {
	parser.expect(TokenType('('))
	expr := parser.parseExpr()
	parser.expect(TokenType(')'))
	return expr
}

func (parser *Parser) parseUnaryExpr() UnaryExprNode {
	op := parser.expect(TokenType('+'), TokenType('-'), TokenType('!'))
	expr := parser.parsePrimaryExpr()
	return UnaryExprNode{Operator: op, Value: expr}
}
