package hasm
import (
	"fmt"
)

type Symboltable struct {
	variableCount int
	symbols       map[string]int
}

func NewSymboltable() *Symboltable {
	symboltable := &Symboltable{0, make(map[string]int)}
	symboltable.symbols["SP"] = 0
	symboltable.symbols["LCL"] = 1
	symboltable.symbols["ARG"] = 2
	symboltable.symbols["THIS"] = 3
	symboltable.symbols["THAT"] = 4
	symboltable.symbols["R0"] = 0
	symboltable.symbols["R1"] = 1
	symboltable.symbols["R2"] = 2
	symboltable.symbols["R3"] = 3
	symboltable.symbols["R4"] = 4
	symboltable.symbols["R5"] = 5
	symboltable.symbols["R6"] = 6
	symboltable.symbols["R7"] = 7
	symboltable.symbols["R8"] = 8
	symboltable.symbols["R9"] = 9
	symboltable.symbols["R10"] = 10
	symboltable.symbols["R11"] = 11
	symboltable.symbols["R12"] = 12
	symboltable.symbols["R13"] = 13
	symboltable.symbols["R14"] = 14
	symboltable.symbols["R15"] = 15
	symboltable.symbols["SCREEN"] = 0x4000
	symboltable.symbols["KBD"] = 0x6000
	return symboltable
}

func (symboltable *Symboltable) DefineVariable(symbol string) {
	_, alreadyDefined := symboltable.symbols[symbol]
	if !alreadyDefined {
		symboltable.symbols[symbol] = 16 + symboltable.variableCount
		symboltable.variableCount++
	}
}

func (symboltable *Symboltable) DefineLabel(symbol string, location int) {
	_, alreadyDefined := symboltable.symbols[symbol]
	if alreadyDefined {panic("Label " + symbol + " defined multiple times")}
	symboltable.symbols[symbol] = location
}

func (symboltable *Symboltable) Lookup(symbol string) string {
	location, ok := symboltable.symbols[symbol]
	if !ok {panic("symbol " + symbol + " not found")}
	return fmt.Sprintf("%015b", location)
}