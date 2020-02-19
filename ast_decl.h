#ifndef _H_ast_decl
#define _H_ast_decl

#include "ast.h"

class Decl : public Node 
{
protected:
    Identifier *ident;
public:
    Decl(Identifier *id) : ident(id) {}
};

class VarDecl : public Decl
{
protected:
    Type *type;
public:
    VarDecl(Identifier *id, Type *t) : Decl(id), type(t) {}
};

#endif
