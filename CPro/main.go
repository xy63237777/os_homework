package main

import (
	"io/ioutil"
	"os"
)

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
	file, err := os.Open("./123.txt")
	CheckErrorForExitOfMsg(err, "Error for open file ")
	bytes, err := ioutil.ReadAll(file)
	//fmt.Println(string(bytes))
	CheckErrorForExitOfMsg(err, "Error for ReadAll file", file.Name())
	bytes = FilterResource(bytes)
	ParseFile(bytes,os.Stdout)
}


