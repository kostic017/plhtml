package scope

import (
    "go/constant"
    "plhtml/logger"
)

var myLogger = logger.New("SCOPE")

func SetLogLevel(level logger.LogLevel) {
    myLogger.SetLevel(level)
}

type Scope struct {
    Id      int
    Parent  *Scope
    symbols map[string]*Symbol
}

type Symbol struct {
    Name  string
    Type  string
    Value constant.Value
}

const (
    TypeInteger = "integer"
    TypeReal    = "real"
    TypeBoolean = "boolean"
    TypeString  = "string"
)

func New(id int, parent *Scope) *Scope {
    scope := &Scope{
        Id:     id,
        Parent: parent,
    }
    if parent == nil {
        myLogger.Debug("Creating root scope %d.", id)
        scope.symbols = map[string]*Symbol{
            TypeInteger: {Name: TypeInteger},
            TypeReal:    {Name: TypeReal},
            TypeBoolean: {Name: TypeBoolean},
            TypeString:  {Name: TypeString},
        }
    } else {
        myLogger.Debug("Creating new scope %d as child of %d.", id, parent.Id)
        scope.symbols = make(map[string]*Symbol)
    }
    return scope
}

func (scope *Scope) Insert(sym *Symbol) {
    myLogger.Debug("Inserting %s into scope %d.", sym.Name, scope.Id)
    scope.symbols[sym.Name] = sym
}

func (scope *Scope) DeclaredLocally(name string) bool {
    myLogger.Debug("Checking if %s is declared in scope %d.", name, scope.Id)
    _, ok := scope.symbols[name]
    return ok
}

func (scope *Scope) Lookup(name string) (*Symbol, bool) {
    for currentScope := scope; currentScope != nil; currentScope = currentScope.Parent {
        myLogger.Debug("Looking for %s in scope %d.", name, currentScope.Id)
        if sym, ok := currentScope.symbols[name]; ok {
            myLogger.Debug("Found %s in scope %d.", name, currentScope.Id)
            return sym, true
        }
    }
    myLogger.Debug("Symbol %s not declared.", name)
    return nil, false
}

func (scope *Scope) GetValue(name string) constant.Value {
    sym, _ := scope.Lookup(name)
    return sym.Value
}

func (scope *Scope) SetValue(name string, value constant.Value) {
    sym, _ := scope.Lookup(name)
    sym.Value = value
}
