package local

import (
	"fmt"
	"os"
)

type Scope interface {
	Define()
	Resolve()
}

type LocalTable struct {
	Scope
	Symbols map[string]*interface{}
}

func (localTable *LocalTable) Define(name string, value interface{}) {
	if _, ok := localTable.Symbols[name]; ok {
		fmt.Printf("Already defined symbol: %s", name)
		os.Exit(1)
	}
}

func (localTable *LocalTable) Resolve(name string) interface{} {
	if val, ok := localTable.Symbols[name]; ok {
		return val
	}

	fmt.Printf("Unknown symbol: %s", name)
	os.Exit(1)

	return nil
}
