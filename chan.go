package main

import (
	"time"
	"fmt"
)

var chs chan int = make(chan int, 1)

func write()  {
	time.Sleep(3*time.Second)
	chs <- 1
}

//golang的select与channel配合使用。它用于等待一个或者多个channel的输出。
//select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
//select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。
func read()  {
	select {
		case ch1 := <- chs:
			fmt.Println(ch1)
			return
		case <- time.After(time.Second): //time.After(xxx)返回的就是chan
			fmt.Println("read time out")
			return
	}
}

func mainaz()  {
	go write()
	read()
}
//输出结果 read time out
