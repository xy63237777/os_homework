package core

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

func IsEmpty(str string) bool {
	if len(str) == 0 || (len(str) == 1 && str[0] == ' ') {
		return true
	}
	if str[0] == ' ' && str[len(str) - 1] == ' ' && len(str) < 2 {
		return true
	}
	for i := 0; str[i] != ' '; i++  {
		return false
	}
	return true
}

func TrimSuffix(str string) string {
	index := 0
	for str[index] == ' ' {
		index++
	}
	return str[index:]
}

func GetSyncAndString(str string) (int, string) {
	sync := NO_SEARCH

	if SearchReserve(str) != NO_SEARCH {
		sync = SearchReserve(str)
	} else if IsNum(str) {
		return 99,DecToBinaryForInt(str)
	} else if  IsIdentifier(str){
		sync = 100
	} else if IsEmpty(str) {

	} else {
		//fmt.Println("error : ", str)
	}
	return sync,str
}