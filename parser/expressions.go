package parser

import "../scanner"

func (parser *Parser) parseExpr() ExpressionNode {
	// expr1 := parser.parseExpr()

	// for _, ok := parser.binOps[parser.peek().Type]; ok; _, ok = parser.binOps[parser.peek().Type] {
	// 	op := parser.next()
	// 	expr2 := parser.parsePrimaryExpr()
	// }

	return nil
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

	panic("Invalid expression.")
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
	expr := parser.parsePrimaryExpr()
	parser.expect(TokenType(')'))
	return expr
}

func (parser *Parser) parseUnaryExpr() UnaryExprNode {
	op := parser.expect(TokenType('+'), TokenType('-'), TokenType('!'))
	expr := parser.parsePrimaryExpr()
	return UnaryExprNode{Operator: op, Value: expr}
}
