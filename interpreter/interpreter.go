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

// Visit methods should return constant.Value!

func (interpreter *Interpreter) VisitBinaryOpExpr(node ast.BinaryOpExprNode) interface{} {
	leftValue := node.LeftExpr.Accept(interpreter).(constant.Value)
	rightValue := node.RightExpr.Accept(interpreter).(constant.Value)

	if node.Operator == token.Plus && (isStr(leftValue) || isStr(rightValue)) {
		return strcat(leftValue, rightValue)
	}

	if isNum(leftValue) && isNum(rightValue) {
		return opsWithNums(leftValue, node.Operator, rightValue)
	}

	//if isBool(leftValue) && isBool(rightValue) {
	//    return opsWithBools(leftValue, node.Operator, rightValue)
	//}

	panic("Operator " + token.TypeToStr[node.Operator] + " is not supported for operands of given types.")
}

func (interpreter *Interpreter) VisitBoolConst(node ast.BoolConstNode) interface{} {
	return constant.MakeBool(node.Value)
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

func isNum(val constant.Value) bool {
	return val.Kind() == constant.Int || val.Kind() == constant.Float
}

func isStr(val constant.Value) bool {
	return val.Kind() == constant.String
}

func isBool(val constant.Value) bool {
	return val.Kind() == constant.Bool
}

func strcat(leftValue constant.Value, rightValue constant.Value) constant.Value {

	if leftValue.Kind() == constant.String && rightValue.Kind() == constant.String {
		return constant.MakeString(constant.StringVal(leftValue) + constant.StringVal(rightValue))
	}

	if leftValue.Kind() == constant.String {
		leftString := constant.StringVal(leftValue)
		if rightValue.Kind() == constant.Int {
			val, _ := constant.Int64Val(rightValue)
			return constant.MakeString(leftString + strconv.FormatInt(val, 10))
		}
		if rightValue.Kind() == constant.Float {
			val, _ := constant.Float64Val(rightValue)
			return constant.MakeString(leftString + util.FloatToString(val))
		}
		if rightValue.Kind() == constant.Bool {
			val := constant.BoolVal(rightValue)
			return constant.MakeString(leftString + strconv.FormatBool(val))
		}
	}

	if rightValue.Kind() == constant.String {
		rightString := constant.StringVal(rightValue)
		if leftValue.Kind() == constant.Int {
			val, _ := constant.Int64Val(leftValue)
			return constant.MakeString(strconv.FormatInt(val, 10) + rightString)
		}
		if leftValue.Kind() == constant.Float {
			val, _ := constant.Float64Val(leftValue)
			return constant.MakeString(util.FloatToString(val) + rightString)
		}
		if leftValue.Kind() == constant.Bool {
			val := constant.BoolVal(leftValue)
			return constant.MakeString(strconv.FormatBool(val) + rightString)
		}
	}

	panic("You can concatenate string with other strings, integers, floats or booleans.")

}

func opsWithNums(left constant.Value, operator TokenType, right constant.Value) constant.Value {

	if left.Kind() == constant.Int || right.Kind() == constant.Int {

		leftVal, _ := constant.Int64Val(left)
		rightVal, _ := constant.Int64Val(right)

		switch operator {
		case token.Plus:
			return constant.MakeInt64(leftVal + rightVal)
		case token.Minus:
			return constant.MakeInt64(leftVal - rightVal)
		case token.Asterisk:
			return constant.MakeInt64(leftVal * rightVal)
		case token.Slash:
			return constant.MakeInt64(leftVal / rightVal)
		case token.LtOp:
			return constant.MakeBool(leftVal < rightVal)
		case token.GtOp:
			return constant.MakeBool(leftVal > rightVal)
		case token.LeqOp:
			return constant.MakeBool(leftVal <= rightVal)
		case token.GeqOp:
			return constant.MakeBool(leftVal >= rightVal)
		case token.EqOp:
			return constant.MakeBool(leftVal == rightVal)
		case token.NeqOp:
			return constant.MakeBool(leftVal != rightVal)
		default:
			panic("Operator " + token.TypeToStr[operator] + " cannot be applied to two integers.")
		}

	}

	leftVal, _ := constant.Float64Val(left)
	rightVal, _ := constant.Float64Val(right)

	switch operator {
	case token.Plus:
		return constant.MakeFloat64(leftVal + rightVal)
	case token.Minus:
		return constant.MakeFloat64(leftVal - rightVal)
	case token.Asterisk:
		return constant.MakeFloat64(leftVal * rightVal)
	case token.Slash:
		return constant.MakeFloat64(leftVal / rightVal)
	case token.LtOp:
		return constant.MakeBool(leftVal < rightVal)
	case token.GtOp:
		return constant.MakeBool(leftVal > rightVal)
	case token.LeqOp:
		return constant.MakeBool(leftVal <= rightVal)
	case token.GeqOp:
		return constant.MakeBool(leftVal >= rightVal)
	case token.EqOp:
		return constant.MakeBool(leftVal == rightVal)
	case token.NeqOp:
		return constant.MakeBool(leftVal != rightVal)
	default:
		panic("Operator " + token.TypeToStr[operator] + " cannot be applied to two numbers.")
	}

}
