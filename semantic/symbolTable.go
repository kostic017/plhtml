package semantic

type symbol struct {
	name string
}

type SymbolTable struct {
	symbols map[string]symbol
}

func NewSymbolTable() *SymbolTable {
	st := new(SymbolTable)
	st.symbols = make(map[string]symbol)
	return st
}

func (st *SymbolTable) insert(sym symbol) {
	st.symbols[sym.name] = sym
}

func (st *SymbolTable) lookup(name string) (symbol, bool) {
	sym, ok := st.symbols[name]
	return sym, ok
}

func (st *SymbolTable) expect(name string) {
	if _, ok := st.lookup(name); !ok {
		panic("Identifier " + name + " undefined!")
	}
}
