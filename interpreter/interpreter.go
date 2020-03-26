package interpreter

import (
	"bufio"
	"fmt"
	"go/constant"
	"os"
	"strconv"

	"plhtml/ast"
	"plhtml/logger"
	"plhtml/token"
	"plhtml/util"
)

type TokenType = token.Type

var myLogger = logger.New("INTERPRETER")

func SetLogLevel(level logger.LogLevel) {
	myLogger.SetLevel(level)
}

type Interpreter struct {
	in        *bufio.Scanner
	callStack *callStack
}

func New() *Interpreter {
	interp := new(Interpreter)
	interp.in = bufio.NewScanner(os.Stdin)
	interp.callStack = NewStack()
	return interp
}

/**********************************************
 * Visit methods should return constant.Value *
 **********************************************/

func (interp *Interpreter) VisitBinaryOpExpr(node ast.BinaryOpExprNode) interface{} {
	leftValue := node.LeftExpr.Accept(interp).(constant.Value)
	rightValue := node.RightExpr.Accept(interp).(constant.Value)
	myLogger.Debug("%s %s %s", leftValue.String(), node.Operator.String(), rightValue.String())

	if node.Operator == token.Plus && (isStr(leftValue) || isStr(rightValue)) {
		return strcat(leftValue, rightValue)
	}

	if isNum(leftValue) && isNum(rightValue) {
		return opsWithNums(leftValue, node.Operator, rightValue)
	}

	//if isBool(leftValue) && isBool(rightValue) {
	//    return opsWithBools(leftValue, node.Operator, rightValue)
	//}

	panic("Binary operator " + node.Operator.String() + " is not supported for operands of given types.")
}

func (interp *Interpreter) VisitBoolConst(node ast.BoolConstNode) interface{} {
	myLogger.Debug(fmt.Sprintf("Bool %t", node.Value))
	return constant.MakeBool(node.Value)
}

func (interp *Interpreter) VisitControlFlowStmt(node ast.ControlFlowStmtNode) {
	myLogger.Debug(node.Type.String())
	switch node.Type {
	case token.If:
		if constant.BoolVal(node.Condition.Accept(interp).(constant.Value)) {
			for _, stmt := range node.Statements {
				stmt.Accept(interp)
			}
		}
		break
	case token.While:
		for constant.BoolVal(node.Condition.Accept(interp).(constant.Value)) {
			for _, stmt := range node.Statements {
				stmt.Accept(interp)
			}
		}
	}
}

func (interp *Interpreter) VisitIdentifier(node ast.IdentifierNode) interface{} {
	arcRecord := interp.callStack.peek()
	value := arcRecord.variables[node.Name].(constant.Value)
	myLogger.Debug("%s: %s", node.Name, value.String())
	return value
}

func (interp *Interpreter) VisitIntConst(node ast.IntConstNode) interface{} {
	return constant.MakeInt64(int64(node.Value))
}

func (interp *Interpreter) VisitMainFunc(node ast.MainFuncNode) {
	interp.callStack.push(newActRecord()) // function local variables
	for _, stmt := range node.Statements {
		stmt.Accept(interp)
	}
	interp.callStack.pop()
}

func (interp *Interpreter) VisitProgram(node ast.ProgramNode) {
	fmt.Println(node.Title.Value)
	node.Body.Accept(interp)
}

func (interp *Interpreter) VisitProgramBody(node ast.ProgramBodyNode) {
	interp.callStack.push(newActRecord()) // for global variables
	node.MainFunc.Accept(interp)
	interp.callStack.pop()
}

func (interp *Interpreter) VisitReadStmt(node ast.ReadStmtNode) {
	// TODO correct type
	interp.in.Scan()

	val, err := util.StrToInt64(interp.in.Text())
	util.Check(err)

	actRecord := interp.callStack.peek()
	actRecord.variables[node.Identifier.Name] = constant.MakeInt64(val)
}

func (interp *Interpreter) VisitRealConst(node ast.RealConstNode) interface{} {
	return constant.MakeFloat64(node.Value)
}

func (interp *Interpreter) VisitStringConst(node ast.StringConstNode) interface{} {
	return constant.MakeString(node.Value)
}

func (interp *Interpreter) VisitUnaryExpr(node ast.UnaryExprNode) interface{} {
	exprValue := node.Expr.Accept(interp).(constant.Value)

	switch node.Operator {
	case token.Minus:
		if exprValue.Kind() == constant.Int {
			exprVal, _ := constant.Int64Val(exprValue)
			return constant.MakeInt64(-exprVal)
		} else if exprValue.Kind() == constant.Float {
			exprVal, _ := constant.Float64Val(exprValue)
			return constant.MakeFloat64(-exprVal)
		}
		break
	case token.Exclamation:
		if exprValue.Kind() == constant.Bool {
			return !constant.BoolVal(exprValue)
		}
	}

	panic("Unary operator " + node.Operator.String() + " is not supported for given types.")
}

func (interp *Interpreter) VisitVarAssign(node ast.VarAssignNode) {
	// TODO correct type
	value := node.Value.Accept(interp).(constant.Value)
	myLogger.Debug("%s = %s", node.Identifier.Name, value.String())
	actRecord := interp.callStack.peek()
	actRecord.variables[node.Identifier.Name] = value
}

func (interp *Interpreter) VisitVarDecl(node ast.VarDeclNode) {
	// TODO correct type and default value
	myLogger.Debug("%s %s", node.Type.Name, node.Identifier.Name)
	actRecord := interp.callStack.peek()
	actRecord.variables[node.Identifier.Name] = constant.MakeInt64(0)
}

func (interp *Interpreter) VisitWriteStmt(node ast.WriteStmtNode) {
	exprValue := node.Value.Accept(interp).(constant.Value)
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
			panic("Operator " + operator.String() + " cannot be applied to two integers.")
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
		panic("Operator " + operator.String() + " cannot be applied to two numbers.")
	}

}
