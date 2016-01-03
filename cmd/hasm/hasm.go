package main
import (
	"translation"
	"hasm"
)

func main() {
	translation.RunTranslator(new(hasm.Assembler))
}