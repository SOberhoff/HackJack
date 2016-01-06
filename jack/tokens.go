package jack

type keyword string

const (
	CLASS keyword = "class"
	CONSTRUCTOR keyword = "constructor"
	FUNCTION keyword = "function"
	METHOD keyword = "method"
	FIELD keyword = "field"
	STATIC keyword = "static"
	VAR keyword = "var"
	INT keyword = "int"
	CHAR keyword = "char"
	BOOLEAN keyword = "boolean"
	VOID keyword = "void"
	TRUE keyword = "true"
	FALSE keyword = "false"
	NULL keyword = "null"
	THIS keyword = "this"
	LET keyword = "let"
	DO keyword = "do"
	IF keyword = "if"
	ELSE keyword = "else"
	WHILE keyword = "while"
	RETURN keyword = "return"
)

type symbol string

const (
	OPEN_BRACE symbol = "{"
	CLOSE_BRACE symbol = "}"
	OPEN_PAREN symbol = "("
	CLOSE_PAREN symbol = ")"
	OPEN_BRACKET symbol = "["
	CLOSE_BRACKET symbol = "]"
	DOT symbol = "."
	COMMA symbol = ","
	SEMICOLON symbol = ";"
	PLUS symbol = "+"
	MINUS symbol = "-"
	STAR symbol = "*"
	SLASH symbol = "/"
	AND symbol = "&"
	PIPE symbol = "|"
	LESS_THAN symbol = "<"
	GREATER_THAN symbol = ">"
	EQUALS symbol = "="
	TILDE symbol = "~"
)

type intToken int

type stringToken string

type identifier string


