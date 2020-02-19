#ifndef _H_ast_stmt
#define _H_ast_stmt

#include <vector>

class Stmt
{
};

class PrintStmt : public Stmt
{
protected:
    std::vector<Expr*> *args;
public:
    PrintStmt(std::vector<Expr*> *arguments) : args(arguments) {}
};

#endif
