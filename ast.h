#ifndef _H_ast
#define _H_ast

class Node
{
};

class Identifier : public Node
{
protected:
    const char* name;
public:
    Identifier(const char* n) : name(n) {}
};

#endif
