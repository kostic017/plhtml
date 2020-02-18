#ifndef _H_ast
#define _H_ast

enum Operator
{
    PLUS,
    MINUS,
    TIMES,
    DIVIDE
};

class Node
{
};

class Identifier : public Node
{
    char* name;
public:
    Identifier(const char* n) : name(n) {}
};

#endif
