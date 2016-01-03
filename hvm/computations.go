package hvm


//Generates code that will perform a unary computation on the current operand on the stack.
//"-" -> neg
//"!" -> not
func unaryComputation(computation string) (output string) {
	output += popOneOperand()
	output += "D=" + computation + "D\n"
	output += pushD()
	return
}

//Generates code that will perform a binary computation on the current operands on the stack.
//"+" -> add
//"-" -> sub
//"&" -> and
//"|" -> or
func generateBinaryComputation(computation string) (output string) {
	output += popTwoOperands()
	output += "D=D" + computation + "A\n"
	output += pushD()
	return
}

//Generates code that will perform a comparison on the current operands on the stack.
//"JEQ" -> eq
//"JGT" -> gt
//"JLT" -> lt
func comparison(condition string) (output string) {
	trueLabel := labelIndexer.AddIndex("TRUE")
	endLabel := labelIndexer.AddIndex("END")
	output += popTwoOperands()
	output += "D=D-A\n"
	output += "@" + trueLabel + "\n"
	output += "D;" + condition + "\n"
	output += "D=0\n"
	output += pushD()
	output += "@" + endLabel + "\n"
	output += "D;JMP\n"
	output += "(" + trueLabel + ")\n"
	output += "D=-1\n"
	output += pushD()
	output += "(" + endLabel + ")\n"
	return
}

//Generates code that will compute (*segment + index) or (segment + index) as appropriate
// and store it in the target register.
func computeAddress(segment string, index string, targetRegister string) (output string) {
	switch segment {
	case "local":
		output += "@LCL\n"
		output += "D=M\n"
	case "argument":
		output += "@ARG\n"
		output += "D=M\n"
	case "this":
		output += "@THIS\n"
		output += "D=M\n"
	case "that":
		output += "@THAT\n"
		output += "D=M\n"
	case "pointer":
		output += "@THIS\n"
		output += "D=A\n"
	case "temp":
		output += "@5\n"
		output += "D=A\n"
	case "static":
		output += "@" + currentFileName + index + "\n"
		output += targetRegister + "=A\n"
		return
	default:
		panic("unknown segment: " + segment)
	}
	output += "@" + index + "\n"
	output += targetRegister + "=D+A\n"
	return
}
