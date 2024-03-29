package main

import (
	"CPro/core"
	"CPro/files"
	"flag"
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

var input *string = flag.String("in", "./123.txt", "Use -in <input filePath>")
var output *string = flag.String("out", "", "Use -out <output filePath>")
func main() {
	flag.Parse()
	file, err := os.Open(*input)
	core.CheckErrorForExitOfMsg(err, "Error for open file ")
	bytes, err := ioutil.ReadAll(file)
	//fmt.Println(string(bytes))
	core.CheckErrorForExitOfMsg(err, "Error for ReadAll file", file.Name())
	bytes = files.FilterResource(bytes)
	if len(*output) == 0 {
		files.ParseFile(bytes,os.Stdout)
	} else {
		createFile, err := os.Create(*output)
		core.CheckErrorOfMsg(err, "Error for CreateFile :")
		files.ParseFile(bytes,createFile)
	}
}


