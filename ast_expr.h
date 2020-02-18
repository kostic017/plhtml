#ifndef _H_ast_expr
#define _H_ast_expr

#include "ast.h"

template<class T>
class Expr : public Node
{
protected:
    T value;
public:
    Expr(T val) : value(val) {}
};

class IntConst : public Expr<int>
{
};

class RealConst : public Expr<float>
{
};

class BoolConst : public Expr<bool>
{
};

class CompoundExpr : public Expr
{
protected:
    Expr *left;
    Expr *right;
    Operator op;
public:
    CompoundExpr(Operator op, Expr *rhs);
    CompoundExpr(Expr *lhs, Operator op, Expr *rhs);
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
