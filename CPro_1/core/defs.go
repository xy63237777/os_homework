package core

import (
	"io/ioutil"
	"os"
	"strings"
)

var ReserveWordMap = make(map[string]int)
var OperatorOrDelimiterMap = make(map[string]int)
const (
	CONFIG_PATH = "./config/"
	RESERVE_WORD_PATH = "reserveWord.conf"
	OPERATOR_OR_DELIMITER_PATH = "operatorOrDelimiter.conf"
)

func initReserveWordMap(i *int)  {
	doInitMap(ReserveWordMap,i,CONFIG_PATH+RESERVE_WORD_PATH)
}

func initOperatorOrDelimiterMap(i *int)  {
	doInitMap(OperatorOrDelimiterMap, i, CONFIG_PATH + OPERATOR_OR_DELIMITER_PATH)
}

func doInitMap(m map[string]int, i *int, path string)  {
	file, err := os.Open(path)
	defer file.Close()
	CheckErrorForExitOfMsg(err, " open file error for config ", path)
	bytes, err := ioutil.ReadAll(file)
	CheckErrorForExitOfMsg(err, "read file ", path, " error")
	strs := strings.Split(string(bytes), "$$")
	for _, str := range strs {
		m[str] = *i
		*i++
	}
}

func init()  {
	i := 1
	initReserveWordMap(&i)
	initOperatorOrDelimiterMap(&i)
}
