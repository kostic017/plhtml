package parser

import (
    "fmt"

    "../ast"
    "../scanner"
)

func (parser *Parser) parseExpr() ast.ExpressionNode {
    return parser.parseExprL(1)
}

func (parser *Parser) parseExprL(l int) ast.ExpressionNode {
    var operators []TokenType

    switch l {
    case 1:
        operators = []TokenType{scanner.TokEqOp, scanner.TokNeqOp}
    case 2:
        operators = []TokenType{scanner.TokLtOp, scanner.TokGtOp, scanner.TokLeqOp, scanner.TokGeqOp}
    case 3:
        operators = []TokenType{TokenType('+'), TokenType('-')}
    case 4:
        operators = []TokenType{TokenType('*'), TokenType('/')}
    case 5:
        return parser.parseFactor()
    }

    expr := parser.parseExprL(l + 1)
    if parser.expectOpt(operators...) {
        return ast.BinaryOpExprNode{LeftExpr: expr, Operator: parser.next().Type, RightExpr: parser.parseExprL(l + 1)}
    }
    return expr
}

func (parser *Parser) parseFactor() ast.ExpressionNode {
    nextToken := parser.peek()
    switch nextToken.Type {
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
    default:
        panic(fmt.Sprintf("Invalid factor '%s' at %d:%d.", string(nextToken.Type), nextToken.Line, nextToken.Column))
    }
}

func (parser *Parser) parseIdentifier() ast.IdentifierNode {
    myLogger.Debug("=BEG= Identifier")
    parser.expect(scanner.TokIdentifier)
    myLogger.Debug("=END= Identifier")
    return ast.IdentifierNode{Name: parser.current().StrVal}
}

func (parser *Parser) parseIntConst() ast.IntConstNode {
    myLogger.Debug("=BEG= Integer Constant")
    parser.expect(scanner.TokIntConst)
    myLogger.Debug("=END= Integer Constant")
    return ast.IntConstNode{Value: parser.current().IntVal}
}

func (parser *Parser) parseRealConst() ast.RealConstNode {
    myLogger.Debug("=BEG= Real Constant")
    parser.expect(scanner.TokRealConst)
    myLogger.Debug("=END= Real Constant")
    return ast.RealConstNode{Value: parser.current().RealVal}
}

func (parser *Parser) parseBoolConst() ast.BoolConstNode {
    myLogger.Debug("=BEG= Boolean Constant")
    parser.expect(scanner.TokBoolConst)
    myLogger.Debug("=END= Boolean Constant")
    return ast.BoolConstNode{Value: parser.current().BoolVal}
}

func (parser *Parser) parseStringConst() ast.StringConstNode {
    myLogger.Debug("=BEG= String Constant")
    parser.expect(scanner.TokStringConst)
    myLogger.Debug("=END= String Constant")
    return ast.StringConstNode{Value: parser.current().StrVal}
}

func (parser *Parser) parseParenExpr() ast.ExpressionNode {
    myLogger.Debug("=BEG= ()")
    parser.expect(TokenType('('))
    expr := parser.parseExpr()
    parser.expect(TokenType(')'))
    myLogger.Debug("=END= ()")
    return expr
}

func (parser *Parser) parseUnaryExpr() ast.UnaryExprNode {
    myLogger.Debug("=BEG= Unary")
    op := parser.expect(TokenType('+'), TokenType('-'), TokenType('!'))
    expr := parser.parseFactor()
    myLogger.Debug("=END= Unary")
    return ast.UnaryExprNode{Operator: op, Expr: expr}
}
