package main

import (
	"fmt"
	"os"
)

func main() {
	strs := []string {
		"auto", "break", "case", "char", "const", "continue",
		"default", "do", "double", "else", "enum", "extern",
		"float", "for", "goto", "if", "int", "long",
		"register", "return", "short", "signed", "sizeof", "static",
		"struct", "switch", "typedef", "union", "unsigned", "void",
		"volatile", "while",
	}
	file, err := os.Create("./config/reserveWord.conf")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	for i, str := range strs  {
		fmt.Println(str)
		n, err := file.WriteString(str)
		if i != len(strs)-1 {
			file.WriteString("$$")
		}
		fmt.Println(n, "  " ,err)
	}
}
