#ifndef _H_ast_expr
#define _H_ast_expr

#include "ast.h"

enum Operator
{
    PLUS,
    MINUS,
    TIMES,
    DIVIDE,
    EQUAL
};

class Expr : public Node
{
};

class IntConst : public Expr
{
protected:
    int value;
public:
    IntConst(int val) : value(val) {}
};

class RealConst : public Expr
{
protected:
    float value;
public:
    RealConst(float val) : value(val) {}
};

class BoolConst : public Expr
{
protected:
    bool value;
public:
    BoolConst(bool val) : value(val) {}
};

class StringConst : public Expr
{
protected:
    const char* value;
public:
    StringConst(const char* val) : value(val) {}
};

class CompoundExpr : public Expr
{
protected:
    Expr *left;
    Expr *right;
    Operator oper;
public:
    CompoundExpr(Operator op, Expr *rhs) : oper(op), right(rhs)  {}
    CompoundExpr(Expr *lhs, Operator op, Expr *rhs) : left(lhs), oper(op), right(rhs) {}
};

class ArithmeticExpr : public CompoundExpr
{
public:
    ArithmeticExpr(Operator op, Expr *rhs) : CompoundExpr(op, rhs) {}
    ArithmeticExpr(Expr *lhs, Operator op, Expr *rhs) : CompoundExpr(lhs, op, rhs) {}
};

class RelationalExpr : public CompoundExpr
{
public:
    RelationalExpr(Expr *lhs, Operator op, Expr *rhs) : CompoundExpr(lhs, op, rhs) {}
};

class LogicalExpr : public CompoundExpr
{
public:
    LogicalExpr(Operator op, Expr *rhs) : CompoundExpr(op, rhs) {}
    LogicalExpr(Expr *lhs, Operator op, Expr *rhs) : CompoundExpr(lhs, op, rhs) {}
};

class AssignExpr : public CompoundExpr
{
public:
    AssignExpr(Expr *lhs, Expr *rhs) : CompoundExpr(lhs, EQUAL, rhs) {}
};

class LVal : public Expr
{
};

class ReadValueExpr : public Expr
{
};

#endif
