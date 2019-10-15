package files

import (
	"CPro/core"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

//过滤注释
func FilterResource(bys []byte) []byte {
	i := 0
	count := 0
	number := 0
	no := 0
	tempbyte := make([]byte, len(bys))
	for i := 0; bys[i] == ' ' || bys[i] == '\n'; i++ {
		if(bys[i] == '\n') {
			no++
		}
	}
	for ; i < len(bys); i++ {
		if bys[i] == '/'&& bys[i + 1] == '/' {//若为单行注释“//”,则去除注释后面的东西，直至遇到回车换行
			for bys[i] != '\n' {
				i++//向后扫描
				number++
			}
			no++
		}
		if bys[i] == '/' && bys [i + 1] == '*' {//若为多行注释“/* 。。。*/”则去除该内容
			i += 2
			number+=2
			for bys[i] != '*' || bys[i + 1] != '/' {
				if bys[i] == '\n' {
					no++
				}
				i++ //继续扫描
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
	return tempbyte

}

func ParseFile(bys []byte,writer io.Writer)  {
	count :=  0
	var sync int

	for i := 0; i < len(bys); i++ {
		if bys[i] == ' ' {
			str := string(bys[count : i])
			sync,str = core.GetSyncAndString(str)
			if sync != core.NO_SEARCH {
				writerFile(writer,GetString("<",str," , ",strconv.Itoa(sync),">","\n"))
			}
			//处理后续空格
			for bys[i] == ' ' {
				i++
			}
			i--
			count = i+1
		} else if bys[i] == '.' && (core.IsNum(string(bys[i-1])) || core.IsNum(string(bys[i+1]))) {
			i++
			continue
		} else if i <= len(bys) - 2 && core.SearchOperatorOrDelimiter(string(bys[i:i+2])) != core.NO_SEARCH {
			str := string(bys[count : i])
			fmt.Println("--- ", str)
			sync,str = core.GetSyncAndString(str)
			if sync != core.NO_SEARCH {
				writerFile(writer,GetString("<",str," , ",strconv.Itoa(sync),">","\n"))
			}
			str = string(bys[i:i+2])
			sync = core.SearchOperatorOrDelimiter(str)
			writerFile(writer,GetString("<",str," , ",strconv.Itoa(sync),">","\n"))
			count = i+1
			i++
		} else if core.SearchOperatorOrDelimiter(string(bys[i])) != core.NO_SEARCH {
			str := string(bys[count : i])
			sync,str = core.GetSyncAndString(str)
			if sync != core.NO_SEARCH {
				writerFile(writer,GetString("<",str," , ",strconv.Itoa(sync),">","\n"))
			}
			str = string(bys[i:i+1])
			sync = core.SearchOperatorOrDelimiter(str)
			writerFile(writer,GetString("<",str," , ",strconv.Itoa(sync),">","\n"))
			count = i+1

		}
	}
}

func doWriter(writer io.Writer, str string) {
	//sync := GetSync(str)
	//if sync != NO_SEARCH {
	//	writerFile(writer,GetString("<",str," , ",strconv.Itoa(sync),">","\n"))
	//}
}

func GetString(str ...string) string {
	if len(str) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _,s := range str {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func writerFile(writer io.Writer, str string)  {
	_, err := io.WriteString(writer, str)
	core.CheckErrorOfMsg(err, "Writer Error : ")
}