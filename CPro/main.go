package main

import "os"

var ReserveWordMap = make(map[string]int)
var OperatorOrDelimiterMap = make(map[string]int)
func init() {
	strs := []string{
		"auto", "break", "case", "char", "const", "continue",
		"default", "do", "double", "else", "enum", "extern",
		"float", "for", "goto", "if", "int", "long",
		"register", "return", "short", "signed", "sizeof", "static",
		"struct", "switch", "typedef", "union", "unsigned", "void",
		"volatile", "while",
	}
	for i, str := range strs {
		ReserveWordMap[str] = i + 1
	}
	strs = []string {
		"+", "-", "*", "/", "<", "<=", ">", ">=", "=", "==",
		"!=", ";", "(", ")", "^", ",", "\"", `\`, "#", "&",
		"&&", "|", "||", "%", "~", "<<", ">>", "[", "]", "{",
		"}", "\\", ".", `\?`, ":", "!",
	}
	for i, str := range strs {
		OperatorOrDelimiterMap[str] = i + +33
	}
}

func main() {
	file, e := os.Open()
}
