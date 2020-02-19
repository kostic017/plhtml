#ifndef _H_ast_type
#define _H_ast_type

#include "ast.h"

class Type : public Node
{
protected:
    const char* typeName;
public:
    static Type *intType,
                *realType,
                *boolType,
                *stringType;
    Type(const char* name) : typeName(name) {}
};

#endif
