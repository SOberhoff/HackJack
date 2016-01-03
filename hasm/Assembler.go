package hasm
import (
	"fmt"
	"strconv"
)

type Assembler struct{}

func (assembler *Assembler) InputExtension() string {
	return "asm"
}

func (assembler *Assembler) OutputExtension() string {
	return "hack"
}

func (assembler *Assembler) String() string {
	return "Assembler"
}

func (assembler *Assembler) Translate(input map[string]string) (output map[string]string) {
	output = make(map[string]string, 0)
	for baseName, inputString := range input {
		output[baseName] = translateSingleFile(baseName, inputString)
	}
	return
}

func translateSingleFile(filename string, input string) (output string) {
	result := ""
	for _, loc := range ConvertToLoCs(input) {
		if loc.translation != "" {
			result += loc.translation + "\n"
		}
	}
	return result
}

func ConvertToLoCs(input string) []LoC {
	locs := splitIntoLoCs(input)
	symboltable := buildSymbolTable(locs)
	addTranslations(locs, symboltable)
	return locs
}

func buildSymbolTable(locs []LoC) *Symboltable {
	symboltable := NewSymboltable()
	for _, loc := range locs {
		if loc.isLabel {
			symboltable.DefineLabel(loc.getLabel(), loc.findLabelTarget(locs))
		}
	}
	for _, loc := range locs {
		if loc.isAInstruction {
			immediate, isVariable := loc.getImmediate()
			if isVariable {
				symboltable.DefineVariable(immediate)
			}
		}
	}
	return symboltable
}

func addTranslations(locs []LoC, symboltable *Symboltable) {
	for i, _ := range locs {
		if locs[i].isAInstruction {
			translateAInstruction(&locs[i], symboltable)
		} else if locs[i].isCInstruction {
			translateCInstruction(&locs[i])
		}
	}
}

func translateAInstruction(loc *LoC, symboltable *Symboltable) {
	immediate, isVariable := loc.getImmediate()
	if isVariable {
		loc.translation = "0" + symboltable.Lookup(immediate)
	} else {
		immediateInt, _ := strconv.ParseInt(immediate, 10, 64)
		if immediateInt < 0 {
			immediateInt = 32768 + immediateInt
		}
		loc.translation = "0" + fmt.Sprintf("%015b", immediateInt)
	}
}

func translateCInstruction(loc *LoC) {
	dest, comp, jmp := loc.splitCInstruction()
	loc.translation = "111" + translateComp(comp) + translateDest(dest) + translateJmp(jmp)
}

func translateComp(comp string) string {
	switch comp {
	case "0":
		return "0101010"
	case "1":
		return "0111111"
	case "-1":
		return "0111010"
	case "D":
		return "0001100"
	case "A":
		return "0110000"
	case "!D":
		return "0001101"
	case "!A":
		return "0110001"
	case "-D":
		return "0001111"
	case "-A":
		return "0110011"
	case "D+1", "1+D":
		return "0011111"
	case "A+1", "1+A":
		return "0110111"
	case "D-1", "-1+D":
		return "0001110"
	case "A-1", "-1+A":
		return "0110010"
	case "D+A", "A+D":
		return "0000010"
	case "D-A", "-A+D":
		return "0010011"
	case "A-D", "-D+A":
		return "0000111"
	case "D&A", "A&D":
		return "0000000"
	case "D|A", "A|D":
		return "0010101"
	case "M":
		return "1110000"
	case "!M":
		return "1110001"
	case "-M":
		return "1110011"
	case "M+1", "1+M":
		return "1110111"
	case "M-1", "-1+M":
		return "1110010"
	case "D+M", "M+D":
		return "1000010"
	case "D-M", "-M+D":
		return "1010011"
	case "M-D", "-D+M":
		return "1000111"
	case "D&M", "M&D":
		return "1000000"
	case "D|M", "M|D":
		return "1010101"
	default:
		panic("Unknown computation: " + comp)
	}
}

func translateDest(dest string) string {
	switch dest {
	case "":
		return "000"
	case "M":
		return "001"
	case "D":
		return "010"
	case "MD", "DM":
		return "011"
	case "A":
		return "100"
	case "AM", "MA":
		return "101"
	case "AD", "DA":
		return "110"
	case "AMD", "ADM", "MAD", "DAM", "MDA", "DMA":
		return "111"
	default:
		panic("unknown destination: " + dest)
	}
}

func translateJmp(jmp string) string {
	switch jmp {
	case "":
		return "000"
	case "JGT":
		return "001"
	case "JEQ":
		return "010"
	case "JGE":
		return "011"
	case "JLT":
		return "100"
	case "JNE":
		return "101"
	case "JLE":
		return "110"
	case "JMP":
		return "111"
	default:
		panic("unknown jump instruction: " + jmp)
	}
}