package semantic

type symbol struct {
    name string
}

type Scope struct {
    id      int
    parent  *Scope
    symbols map[string]symbol
}

func NewScope(id int, parent *Scope) *Scope {
    if parent == nil {
        myLogger.Debug("Creating global scope %d.", id)
    } else {
        myLogger.Debug("Creating new scope %d as child of %d.", id, parent.id)
    }
    scope := new(Scope)
    scope.id = id
    scope.parent = parent
    scope.initializeSymbolTable()
    return scope
}

func (scope *Scope) initializeSymbolTable() {
    scope.symbols = make(map[string]symbol)
    scope.symbols["integer"] = symbol{"integer"}
    scope.symbols["real"] = symbol{"real"}
    scope.symbols["boolean"] = symbol{"boolean"}
    scope.symbols["string"] = symbol{"string"}
}

func (scope *Scope) insert(sym symbol) {
    myLogger.Debug("Inserting %s into scope %d.", sym.name, scope.id)
    scope.symbols[sym.name] = sym
}

func (scope *Scope) lookup(name string) (symbol, bool) {
    for currentScope := scope; currentScope != nil; currentScope = currentScope.parent {
        myLogger.Debug("Looking for %s in scope %d.", name, currentScope.id)
        if sym, ok := currentScope.symbols[name]; ok {
            myLogger.Debug("Found %s in scope %d.", name, currentScope.id)
            return sym, true
        }
    }
    myLogger.Debug("Symbol %s not declared.", name)
    return symbol{}, false
}

func (scope *Scope) expect(name string) {
    if _, ok := scope.lookup(name); !ok {
        panic("Identifier " + name + " undefined.")
    }
}
