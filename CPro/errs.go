package main

import "log"

func CheckErrorForExit(err error)  {
	if err != nil {
		log.Fatalln(err)
	}
}

func CheckErrorForExitOfMsg(err error, msg ...string)  {
	if err != nil {
		log.Fatalln("发生了错误: ",msg,"\n", err)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func CheckErrorOfMsg(err error, msg ...string) {
	if err != nil {
		log.Println("发生了错误: ",msg,"\n", err)
	}
}
