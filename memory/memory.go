package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
)

type freeArea struct {
	id, size, address int
	flag bool
}

type node struct {
	area freeArea
	prev *node
	next *node
}

func NewNode(freeArea freeArea, prev, next *node) *node {
	return &node{
		area:freeArea,
		prev:prev,
		next:next,
	}
}

var areaHead, areaTail *node
var numberMap map[int]bool
var length *int = flag.Int("m", 1280, "Use -m <input filePath>")

func init() {
	flag.Parse()
	numberMap = make(map[int]bool)
	numberMap[0] = true
	temp := freeArea{size: 0,flag:true,}
	tempLast := freeArea{address: 0, size: *length, id: 0, flag: false}
	areaHead,areaTail = NewNode(temp,nil,nil),NewNode(tempLast,nil,nil)
	areaHead.next = areaTail
	areaTail.prev = areaHead
}

func CheckErrorForExit(err error)  {
	if err != nil {
		log.Fatalln(err)
	}
}

func CheckErrorForExitOfMsg(err error, msg ...interface{})  {
	if err != nil {
		log.Fatalln("发生了错误: ",msg,"\n", err)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func CheckErrorOfMsg(err error, msg ...interface{}) {
	if err != nil {
		log.Println("发生了错误: ",msg,"\n", err)
	}
}




func  (this *freeArea) checkId() string {
	if this.id == 0 {
		return "FREE"
	}
	return strconv.Itoa(this.id)
}

func (this *freeArea) checkStatus() string {
	if this.flag {
		return "Allocated"
	}
	return "UnAllocated"
}

func show()  {
	fmt.Println("____________________________________")
	for temp := areaHead.next; temp != nil; temp = temp.next {
		fmt.Println("memory number :  ",temp.area.checkId())
		fmt.Println("partition start address :  ",temp.area.address)
		fmt.Println("partition size :  ", temp.area.size)
		fmt.Println("partition status :  ", temp.area.checkStatus())
		fmt.Println("____________________________________")
	}
}


func doAllocation(id, size int) error {
	if size <= 0 {
		return errors.New("error for alloc error for size must be > 0")
	}
	for temp := areaHead.next; temp != nil; temp = temp.next {
		if !temp.area.flag && temp.area.size >= size {
			if temp.area.size == size {
				temp.area.id = id
				temp.area.flag = true
				return nil
			} else if temp.area.size > size {
				node := NewNode(freeArea{id: id, size: size, flag: true, address: temp.area.address}, temp.prev, temp)
				temp.area.address += size
				temp.area.size -= size
				temp.prev.next = node
				temp.prev = node
				return nil
			}
		}
	}
	return errors.New("error No eligible memory partition ")
}

func bestFit() {
	sort()
	allocation()
}

func sort()  {
	for doTemp := areaHead.next.next; doTemp != nil; doTemp = doTemp.next {
		for temp := doTemp; temp != nil; temp = temp.next {
			if temp.area.size < temp.prev.area.size {
				t_temp := temp.prev.prev
				temp.prev.next = temp.next
				temp.prev.prev.next = temp
				temp.next.prev = temp.prev
				temp.next = temp.prev
				temp.prev = t_temp
			}
		}
	}
}

func doRecycle(id int) error {
	if id <= 0 {
		return errors.New("error for recycle id must  > 0")
	}

	for temp := areaHead; temp != nil; temp = temp.next {
		if id == temp.area.id {
			temp.area.id = 0
			temp.area.flag = false
			if !temp.next.area.flag && temp.next != areaTail {
				temp.area.size += temp.next.area.size
				temp.next.next.prev = temp
				temp.next = temp.next.next
			}
			if temp.next == areaTail {
				areaTail.area.size += temp.area.size
				areaTail.area.address = temp.area.address
				areaTail.prev = temp.prev
				temp.prev.next = areaTail
				temp = areaTail
			}
			if !temp.prev.area.flag {
				if temp == areaTail {
					areaTail.area.address = temp.prev.area.address
					areaTail.area.size += temp.prev.area.size
					areaTail.prev.prev.next = areaTail
					areaTail.prev = areaTail
					return nil
				}
				temp.prev.area.size += temp.area.size
				temp.prev.next = temp.next
				temp.next.prev = temp.prev
			}

			return nil
		}
	}
	return errors.New("waring: not found memory id")
}

func recycle()  {
	var id int
	fmt.Print("Please enter the memory ID to release : ")
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("error for input : ",err)
		return
	}
	if _, ok := numberMap[id]; !ok {
		fmt.Println("memory ID not found Please re-input")
		return
	}
	err = doRecycle(id)
	CheckErrorOfMsg(err)
}

func firstFit() {
	allocation()
}

func worstFit()  {
	fmt.Println("暂未实现...")
}

func menu()  {
	fmt.Print("Please enter your action : ")
	fmt.Println( "1.申请内存    2.释放内存   3.退出 ")
}

func allocation()  {
	var id, size int
	fmt.Print("Please enter memory ID : ")
	_, err := fmt.Scan(&id)
	if err != nil {
		log.Println("error for input : ", err)
		return
	}
	if _, ok := numberMap[id]; ok {
		fmt.Println("memory ID Already exist Please re-input")
		return
	}
	fmt.Print("Please enter memory size : ")
	_, err = fmt.Scan(&size)
	if err != nil {
		log.Println("error for input : ", err)
		return
	}
	err = doAllocation(id, size)
	if err != nil {
		log.Println(err)
		return
	}
	numberMap[id] = true
}

func selectAlgorithm() {
	fmt.Print("Please enter your action : ")
	fmt.Println( "1.首次适应算法    2.最佳适应算法   3.最坏适应算法(未实现) ")
	var key int
	_, err := fmt.Scan(&key)
	CheckErrorForExitOfMsg(err,"error input error : ")
	switch key {
	case 1:firstFit()
	case 2:bestFit()
	case 3:worstFit()
	default:CheckError(errors.New("Please enter key = 1-3"))

	}

}

func main() {

	for ; ;  {
		menu()
		var key int
		_, err := fmt.Scan(&key)
		CheckErrorForExitOfMsg(err,"error input error : ")
		switch key {
		case 1:selectAlgorithm()
		show()
		case 2:recycle()
		sort()
		show()
		case 3:
			return
		default:
			fmt.Println("input error Re input")
		}
	}
}