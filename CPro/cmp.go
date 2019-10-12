package main

import (
	"regexp"
	"strconv"
	"strings"
)

const  (
	NO_SEARCH = -1
)

//查找保留字
func SearchReserve(str string) int {
	i, ok := ReserveWordMap[str]
	if ok {
		return i
	}
	return NO_SEARCH
}

func SearchOperatorOrDelimiter(str string) int  {
	i, ok := OperatorOrDelimiterMap[str]
	if ok {
		return i
	}
	return NO_SEARCH
}



func IsNum(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func IsIdentifier(str string) bool {
	pattern := "[a-zA-Z][a-zA-Z\\D_]*"
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}

func TrimString(str string) string {
	if str[0] == ' ' && str[len(str) - 1] == ' ' {
		return str
	}
	return strings.Trim(str," ")
}

func TrimSuffix(str string) string {
	index := 0
	for str[index] == ' ' {
		index++
	}
	return str[index:]
}

func GetSync(str string) int {
	sync := NO_SEARCH

	if SearchReserve(str) != NO_SEARCH {
		sync = SearchReserve(str)
	} else if IsNum(str) {
		sync = 99
	} else if  IsIdentifier(str){
		sync = 100
	}
	return sync
}