package interpreter

import (
	"../ast"
	"../logger"
	"../token"
	"../util"
	"fmt"
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

/**********************************************
 * Visit methods should return constant.Value *
 **********************************************/

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

	panic("Binary operator " + token.TypeToStr[node.Operator] + " is not supported for operands of given types.")
}

func (interpreter *Interpreter) VisitBoolConst(node ast.BoolConstNode) interface{} {
	return constant.MakeBool(node.Value)
}

func (interpreter *Interpreter) VisitControlFlowStmt(node ast.ControlFlowStmtNode) {
	switch node.Type {
	case token.If:
		if constant.BoolVal(node.Condition.Accept(interpreter).(constant.Value)) {
			for _, stmt := range node.Statements {
				stmt.Accept(interpreter)
			}
		}
		break
	case token.While:
		for constant.BoolVal(node.Condition.Accept(interpreter).(constant.Value)) {
			for _, stmt := range node.Statements {
				stmt.Accept(interpreter)
			}
		}
	}
}

func (interpreter *Interpreter) VisitIdentifier(node ast.IdentifierNode) interface{} {
	return nil // TODO
}

func (interpreter *Interpreter) VisitIntConst(node ast.IntConstNode) interface{} {
	return constant.MakeInt64(int64(node.Value))
}

func (interpreter *Interpreter) VisitMainFunc(node ast.MainFuncNode) {
	for _, stmt := range node.Statements {
		stmt.Accept(interpreter)
	}
}

func (interpreter *Interpreter) VisitProgram(node ast.ProgramNode) {
	node.Body.Accept(interpreter)
}

func (interpreter *Interpreter) VisitProgramBody(node ast.ProgramBodyNode) {
	node.MainFunc.Accept(interpreter)
}

func (interpreter *Interpreter) VisitReadStmt(node ast.ReadStmtNode) {
	// TODO
}

func (interpreter *Interpreter) VisitRealConst(node ast.RealConstNode) interface{} {
	return constant.MakeFloat64(node.Value)
}

func (interpreter *Interpreter) VisitStringConst(node ast.StringConstNode) interface{} {
	return constant.MakeString(node.Value)
}

func (interpreter *Interpreter) VisitUnaryExpr(node ast.UnaryExprNode) interface{} {
	exprValue := node.Expr.Accept(interpreter).(constant.Value)

	switch node.Operator {
	case token.Minus:
		if exprValue.Kind() == constant.Int {
			exprVal, _ := constant.Int64Val(exprValue)
			return -exprVal
		} else if exprValue.Kind() == constant.Float {
			exprVal, _ := constant.Float64Val(exprValue)
			return -exprVal
		}
		break
	case token.Exclamation:
		if exprValue.Kind() == constant.Bool {
			return !constant.BoolVal(exprValue)
		}
	}

	panic("Unary operator " + token.TypeToStr[node.Operator] + " is not supported for given types.")
}

func (interpreter *Interpreter) VisitVarAssign(node ast.VarAssignNode) {
	// TODO
}

func (interpreter *Interpreter) VisitVarDecl(node ast.VarDeclNode) {
	// TODO
}

func (interpreter *Interpreter) VisitWriteStmt(node ast.WriteStmtNode) {
	exprValue := node.Value.Accept(interpreter).(constant.Value)
	if exprValue.Kind() == constant.String {
		fmt.Print(constant.StringVal(exprValue))
	} else {
		panic("You can print strings only.")
	}
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
		return strcatImplicitConversion(leftValue, rightValue, true)
	}
	if rightValue.Kind() == constant.String {
		return strcatImplicitConversion(rightValue, leftValue, false)
	}
	panic("Could not perform string concatenation.")
}

func strcatImplicitConversion(stringValue constant.Value, otherValue constant.Value, ordered bool) constant.Value {

	stringVal := constant.StringVal(stringValue)

	switch otherValue.Kind() {
	case constant.Int:
		intVal, _ := constant.Int64Val(otherValue)
		intValStr := strconv.FormatInt(intVal, 10)
		if ordered {
			return constant.MakeString(stringVal + intValStr)
		} else {
			return constant.MakeString(intValStr + stringVal)
		}
	case constant.Float:
		floatVal, _ := constant.Float64Val(otherValue)
		floatValStr := util.FloatToString(floatVal)
		if ordered {
			return constant.MakeString(stringVal + floatValStr)
		} else {
			return constant.MakeString(floatValStr + stringVal)
		}
	case constant.Bool:
		boolVal := constant.BoolVal(otherValue)
		boolValStr := strconv.FormatBool(boolVal)
		if ordered {
			return constant.MakeString(stringVal + boolValStr)
		} else {
			return constant.MakeString(boolValStr + stringVal)
		}
	}

	panic("You can concatenate string with other strings, integers, floats and booleans.")

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
