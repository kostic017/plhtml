package semantic

type symbol struct {
	name string
}

type SymbolTable struct {
	level   int
	symbols map[string]symbol
}

func NewSymbolTable(level int) *SymbolTable {
	st := new(SymbolTable)
	st.level = level
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
