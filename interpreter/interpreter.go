package interpreter

import (
	"../ast"
	"../logger"
	"../token"
	"../util"
	"go/constant"
	"strconv"
)

type TokenType = token.Type

var myLogger = logger.New("INTERPRETER")

func SetLogLevel(level logger.LogLevel) {
	myLogger.SetLevel(level)
}

type Interpreter struct {
}

func New() *Interpreter {
	interpreter := new(Interpreter)
	return interpreter
}

func (interpreter *Interpreter) VisitBinaryOpExpr(node ast.BinaryOpExprNode) interface{} {
	left := node.LeftExpr.Accept(interpreter).(constant.Value)
	right := node.RightExpr.Accept(interpreter).(constant.Value)

	if node.Operator == token.Plus {
		if val, ok := calc(left, node.Operator, right); ok {
			return val
		}
		if val, ok := strcat(left, right); ok {
			return val
		}
	}

	panic("Operator + is not supported for operands of given types.")
}

func (interpreter *Interpreter) VisitBoolConst(node ast.BoolConstNode) interface{} {
	return nil
}

func (interpreter *Interpreter) VisitControlFlowStmt(node ast.ControlFlowStmtNode) interface{} {
	node.Condition.Accept(interpreter)
	for _, stmt := range node.Statements {
		stmt.Accept(interpreter)
	}
	return nil
}

func (interpreter *Interpreter) VisitIdentifier(node ast.IdentifierNode) interface{} {
	return nil
}

func (interpreter *Interpreter) VisitIntConst(node ast.IntConstNode) interface{} {
	return nil
}

func (interpreter *Interpreter) VisitMainFunc(node ast.MainFuncNode) interface{} {
	for _, stmt := range node.Statements {
		stmt.Accept(interpreter)
	}
	return nil
}

func (interpreter *Interpreter) VisitProgram(node ast.ProgramNode) interface{} {
	node.Body.Accept(interpreter)
	return nil
}

func (interpreter *Interpreter) VisitProgramBody(node ast.ProgramBodyNode) interface{} {
	node.MainFunc.Accept(interpreter)
	return nil
}

func (interpreter *Interpreter) VisitReadStmt(node ast.ReadStmtNode) interface{} {
	return nil
}

func (interpreter *Interpreter) VisitRealConst(node ast.RealConstNode) interface{} {
	return nil
}

func (interpreter *Interpreter) VisitStringConst(node ast.StringConstNode) interface{} {
	return nil
}

func (interpreter *Interpreter) VisitUnaryExpr(node ast.UnaryExprNode) interface{} {
	node.Expr.Accept(interpreter)
	return nil
}

func (interpreter *Interpreter) VisitVarAssign(node ast.VarAssignNode) interface{} {
	node.Value.Accept(interpreter)
	return nil
}

func (interpreter *Interpreter) VisitVarDecl(node ast.VarDeclNode) interface{} {
	return nil
}

func (interpreter *Interpreter) VisitWriteStmt(node ast.WriteStmtNode) interface{} {
	node.Value.Accept(interpreter)
	return nil
}

func calc(leftValue constant.Value, operator TokenType, rightValue constant.Value) (constant.Value, bool) {

	leftIsNum := leftValue.Kind() == constant.Int || leftValue.Kind() == constant.Float
	rightIsNum := rightValue.Kind() == constant.Int || rightValue.Kind() == constant.Float

	if !leftIsNum || !rightIsNum {
		return nil, false
	}

	if leftValue.Kind() == constant.Int || rightValue.Kind() == constant.Int {

		leftVal, _ := constant.Int64Val(leftValue)
		rightVal, _ := constant.Int64Val(rightValue)

		switch operator {
		case token.Plus:
			return constant.MakeInt64(leftVal + rightVal), true
		case token.Minus:
			return constant.MakeInt64(leftVal - rightVal), true
		case token.Asterisk:
			return constant.MakeInt64(leftVal * rightVal), true
		case token.Slash:
			return constant.MakeInt64(leftVal / rightVal), true
		default:
			return nil, false
		}

	}

	leftVal, _ := constant.Float64Val(leftValue)
	rightVal, _ := constant.Float64Val(rightValue)

	switch operator {
	case token.Plus:
		return constant.MakeFloat64(leftVal + rightVal), true
	case token.Minus:
		return constant.MakeFloat64(leftVal - rightVal), true
	case token.Asterisk:
		return constant.MakeFloat64(leftVal * rightVal), true
	case token.Slash:
		return constant.MakeFloat64(leftVal / rightVal), true
	default:
		return nil, false
	}

}

func strcat(leftValue constant.Value, rightValue constant.Value) (constant.Value, bool) {

	if leftValue.Kind() == constant.String && rightValue.Kind() == constant.String {
		return constant.MakeString(constant.StringVal(leftValue) + constant.StringVal(rightValue)), true
	}

	if leftValue.Kind() == constant.String {
		leftString := constant.StringVal(leftValue)
		if rightValue.Kind() == constant.Int {
			val, _ := constant.Int64Val(rightValue)
			return constant.MakeString(leftString + strconv.FormatInt(val, 10)), true
		}
		if rightValue.Kind() == constant.Float {
			val, _ := constant.Float64Val(rightValue)
			return constant.MakeString(leftString + util.FloatToString(val)), true
		}
		if rightValue.Kind() == constant.Bool {
			val := constant.BoolVal(rightValue)
			return constant.MakeString(leftString + strconv.FormatBool(val)), true
		}
	}

	if rightValue.Kind() == constant.String {
		rightString := constant.StringVal(rightValue)
		if leftValue.Kind() == constant.Int {
			val, _ := constant.Int64Val(leftValue)
			return constant.MakeString(strconv.FormatInt(val, 10) + rightString), true
		}
		if leftValue.Kind() == constant.Float {
			val, _ := constant.Float64Val(leftValue)
			return constant.MakeString(util.FloatToString(val) + rightString), true
		}
		if leftValue.Kind() == constant.Bool {
			val := constant.BoolVal(leftValue)
			return constant.MakeString(strconv.FormatBool(val) + rightString), true
		}
	}

	return nil, false

}
