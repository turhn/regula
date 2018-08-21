package local

import (
	"fmt"
)

// SymbolTable keeps track of the symbol table
type SymbolTable struct {
	Symbols map[string]interface{}
}

// Define adds a new symbol to the symbal table
func (table *SymbolTable) Define(name string, value interface{}) {
	if table.Symbols == nil {
		table.Symbols = make(map[string]interface{}, 0)
	}
	table.Symbols[name] = value
}

// Resolve resolves the symbol from the table or fails
func (table *SymbolTable) Resolve(name string) interface{} {
	if val, ok := table.Symbols[name]; ok {
		return val
	}

	panic(fmt.Sprintf("Unresolved symbol %s", name))
}
