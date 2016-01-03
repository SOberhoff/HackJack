package hvm


//Generates code that will decrement the stack pointer once and leave the operand
//in D.
func popOneOperand() (output string) {
	//decrement stack pointer
	output += "@SP\n"
	output += "D=M-1\n"
	output += "@SP\n"
	output += "M=D\n"
	//load operand
	output += "A=D\n"
	output += "D=M\n"
	return
}

//Generates code that will decrement the stack pointer twice and leave the first
//(lower address) operand in D and the second operand in A.
func popTwoOperands() (output string) {
	//decrement stack pointer twice
	output += "@SP\n"
	output += "A=M-1\n"
	output += "D=A-1\n"
	output += "@SP\n"
	output += "M=D\n"
	//load operands
	output += "A=M\n"
	output += "D=M\n"
	output += "A=A+1\n"
	output += "A=M\n"
	return
}

//Generates code that will store the current value of D onto the stack and increment the stack pointer by one.
func pushD() (output string) {
	//store D onto the stack
	output += "@SP\n"
	output += "A=M\n"
	output += "M=D\n"
	//increment stack pointer
	output += "D=A+1\n"
	output += "@SP\n"
	output += "M=D\n"
	return
}

//Generates code that will store the current value of A onto the stack and increment the stack pointer by one.
func pushA() (output string) {
	output += "D=A\n"
	output += pushD()
	return
}