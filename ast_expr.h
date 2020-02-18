#ifndef _H_ast_expr
#define _H_ast_expr

#include "ast.h"

class Expr : public Node
{
};

class IntConst : public Expr
{
    int value;
public:
    IntConst(int val) : value(val) {}
};

class RealConst : public Expr
{
    float value;
public:
    RealConst(float val) : value(val) {}
};

class BoolConst : public Expr
{
    bool value;
public:
    BoolConst(bool val) : value(val) {}
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
    AssignExpr(Expr *lhs, Operator op, Expr *rhs) : CompoundExpr(lhs, op, rhs) {}
};

#endif
