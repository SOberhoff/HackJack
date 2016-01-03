package hvm
import (
	"strings"
	"translation"
)

type VMCompiler struct{}

func (vmcompiler *VMCompiler) InputExtension() string {
	return "vm"
}

func (vmcompiler *VMCompiler) OutputExtension() string {
	return "asm"
}

func (vmcompiler *VMCompiler) String() string {
	return "VMCompiler"
}

func (vmcompiler *VMCompiler) Translate(input map[string]string) (output map[string]string) {
	baseName := determineBaseName(input)
	labelIndexer = translation.NewLabelIndexer()
	outputString := bootstrap()
	for baseName, inputString := range input {
		currentFileName = baseName
		outputString += translateSingleFile(inputString)
	}
	return map[string]string{baseName : outputString}
}

func determineBaseName(input map[string]string) string {
	for baseName, inputString := range input {
		if strings.Contains(inputString, "function Sys.init") {
			return baseName
		}
	}
	panic("no file with Sys.init function found")
}

var labelIndexer translation.LabelIndexer

var currentFileName string

func translateSingleFile(input string) (output string) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := translation.TrimComment(line)
		if trimmedLine == "" {
			continue
		} else {
			tokens := tokenize(trimmedLine)
			switch tokens[0] {
			case "push":
				output += "//push\n"
				output += push(tokens[1], tokens[2])
			case "pop":
				output += "//pop\n"
				output += pop(tokens[1], tokens[2])
			case "neg":
				output += "//neg\n"
				output += unaryComputation("-")
			case "not":
				output += "//not\n"
				output += unaryComputation("!")
			case "add":
				output += "//add\n"
				output += generateBinaryComputation("+")
			case "sub":
				output += "//sub\n"
				output += generateBinaryComputation("-")
			case "and":
				output += "//and\n"
				output += generateBinaryComputation("&")
			case "or":
				output += "//or\n"
				output += generateBinaryComputation("|")
			case "eq":
				output += "//eq\n"
				output += comparison("JEQ")
			case "lt":
				output += "//lt\n"
				output += comparison("JLT")
			case "gt":
				output += "//gt\n"
				output += comparison("JGT")
			case "label":
				output += label(tokens[1])
			case "goto":
				output += "//goto\n"
				output += unconditionalGoto(tokens[1])
			case "if-goto":
				output += "//if-goto\n"
				output += ifGoto(tokens[1])
			case "call":
				output += "//call\n"
				output += call(tokens[1], tokens[2])
			case "return":
				output += "//return\n"
				output += returnStatment()
			case "function":
				output += "//function\n"
				output += function(tokens[1], tokens[2])
			default:
				panic("unknown operation: " + tokens[0])
			}
		}
	}
	output += "(END)\n"
	output += "@END\n"
	output += "D;JMP\n"
	return
}

func bootstrap() (output string) {
	//initialize stack pointer
	output += "@256\n"
	output += "D=A\n"
	output += "@SP\n"
	output += "M=D\n"

	//call Sys.init
	output += call("Sys.init", "0")
	return
}

func tokenize(input string) []string {
	tokens := make([]string, 0)
	for _, token := range strings.Split(input, " ") {
		if token != "" {
			tokens = append(tokens, token)
		}
	}
	if 3 < len(tokens) {panic("Too many tokens in: " + input)}
	return tokens
}