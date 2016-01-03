package hvm
import "strconv"

func label(labelname string) (output string) {
	output += "(" + labelname + ")\n"
	return
}

func unconditionalGoto(labelname string) (output string) {
	output += "@" + labelname + "\n"
	output += "D;JMP\n"
	return
}

func ifGoto(labelname string) (output string) {
	output += popOneOperand()
	output += "@" + labelname + "\n"
	output += "D;JNE\n"
	return
}

func call(functionName string, operandCount string) (output string) {
	returnLabel := labelIndexer.AddIndex("return")
	//save state of current function
	output += "@" + returnLabel + "\n"
	output += pushA()
	output += "@LCL\n"
	output += "D=M\n"
	output += pushD()
	output += "@ARG\n"
	output += "D=M\n"
	output += pushD()
	output += "@THIS\n"
	output += "D=M\n"
	output += pushD()
	output += "@THAT\n"
	output += "D=M\n"
	output += pushD()

	//set up ARG
	output += "@SP\n"
	output += "D=M\n"
	output += "@5\n"
	output += "D=D-A\n"
	output += "@" + operandCount + "\n"
	output += "D=D-A\n"
	output += "@ARG\n"
	output += "M=D\n"

	//set up LCL
	output += "@SP\n"
	output += "D=M\n"
	output += "@LCL\n"
	output += "M=D\n"

	//jump to function
	output += unconditionalGoto(functionName)

	//return here
	output += "(" + returnLabel + ")\n"
	return
}

func returnStatment() (output string) {
	//store LCL in R[13]
	output += "@LCL\n"
	output += "D=M\n"
	output += "@13\n"
	output += "M=D\n"

	//store return address in R[14]
	output += "@5\n"
	output += "A=D-A\n"
	output += "A=M\n"
	output += "D=A\n"
	output += "@14\n"
	output += "M=D\n"

	//position the return value
	output += popOneOperand()
	output += "@ARG\n"
	output += "A=M\n"
	output += "M=D\n"

	//set stack pointer
	output += "D=A+1\n"
	output += "@SP\n"
	output += "M=D\n"

	//restore THAT
	output += "@13\n"
	output += "A=M-1\n"
	output += "D=M\n"
	output += "@THAT\n"
	output += "M=D\n"

	//restore THIS
	output += "@13\n"
	output += "D=M\n"
	output += "@2\n"
	output += "A=D-A\n"
	output += "D=M\n"
	output += "@THIS\n"
	output += "M=D\n"

	//restore ARG
	output += "@13\n"
	output += "D=M\n"
	output += "@3\n"
	output += "A=D-A\n"
	output += "D=M\n"
	output += "@ARG\n"
	output += "M=D\n"

	//restore LCL
	output += "@13\n"
	output += "D=M\n"
	output += "@4\n"
	output += "A=D-A\n"
	output += "D=M\n"
	output += "@LCL\n"
	output += "M=D\n"

	//return control to calling function
	output += "@14\n"
	output += "A=M\n"
	output += "D;JMP\n"
	return
}

func function(functionName string, operandCount string) (output string) {
	intOperandCount, _ := strconv.ParseInt(operandCount, 10, 64)
	output += "(" + functionName + ")\n"
	for i := int64(0); i < intOperandCount; i++ {
		output += "D=0\n"
		output += pushD()
	}
	return
}