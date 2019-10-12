package main

import (
	"fmt"
	"os"
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

//过滤注释
func FilterResource(bys []byte)  {
	count := 0
	tempbyte := make([]byte, len(bys))
	for i := 0; i <= len(bys); i++ {
		if bys[i] == '/'&& bys[i + 1] == '/' {//若为单行注释“//”,则去除注释后面的东西，直至遇到回车换行
			for bys[i] != '\n' {
				i++//向后扫描
			}
		}
		if bys[i] == '/' && bys [i + 1] == '*' {//若为多行注释“/* 。。。*/”则去除该内容
			i += 2
			for bys[i] != '*' || bys[i + 1] != '/' {
				i++//继续扫描
				if i == len(bys) {
					fmt.Printf("注释出错，没有找到 */，程序结束！！！\n")
					os.Exit(1)
				}
			}
			i += 2//跨过“*/”
		}
		if bys[i] != '\n' && bys[i] != '\t'&& bys [i] != '\v'&& bys [i] != '\r' {//若出现无用字符，则过滤；否则加载
			tempbyte[count] = bys[i]
			count++
		}
	}
	bys = tempbyte
}

func IsNum(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func IsIdentifier(str string) bool {
	pattern := "[a-zA-Z][a-zA-Z\\D]*"
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
	if IsNum(str) {
		sync = 99
	} else if IsIdentifier(str) {
		sync = 100
	} else if SearchReserve(str) != NO_SEARCH {
		sync = SearchReserve(str)
	} else if SearchOperatorOrDelimiter(str) != NO_SEARCH {
		sync = SearchOperatorOrDelimiter(str)
	}
	return sync
}