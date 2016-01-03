package main
import (
	"translation"
	"hvm"
)


func main() {
	translation.RunTranslator(new(hvm.VMCompiler))
}
