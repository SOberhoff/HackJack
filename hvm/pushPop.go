package hvm

func push(segment string, index string) (output string) {
	if segment == "constant" {
		output += "@" + index + "\n"
		output += "D=A\n"
		output += pushD()
	} else {
		output += computeAddress(segment, index, "A")
		output += "D=M\n";
		output += pushD()
	}
	return
}

func pop(segment string, index string) (output string) {
	output += computeAddress(segment, index, "D")
	output += "@13\n"
	output += "M=D\n"
	output += popOneOperand()
	output += "@13\n"
	output += "A=M\n"
	output += "M=D\n"
	return
}
