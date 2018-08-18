package local

import (
	"fmt"
	"os"
)

// SymbolTable keeps track of the symbol table
type SymbolTable struct {
	Symbols map[string]interface{}
}

// Define adds a new symbol to the symbal table
func (table *SymbolTable) Define(name string, value interface{}) {
	if _, ok := table.Symbols[name]; ok {
		fmt.Printf("Already defined symbol: %s\n", name)
		os.Exit(1)
	}
}

// Resolve resolves the symbol from the table or fails
func (table *SymbolTable) Resolve(name string) interface{} {
	if val, ok := table.Symbols[name]; ok {
		return val
	}

	fmt.Printf("Unknown symbol: %s\n", name)
	os.Exit(1)

	return nil
}
