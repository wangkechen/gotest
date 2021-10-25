package main

import (
	"fmt"
	"time"
)

func main2() {
	/*var b bytes.Buffer
	b.Write([]byte("你好Hello"))
	fmt.Println(b.String())*/


	/*b1 := new(bytes.Buffer)
	b1.Write([]byte("world"))
	fmt.Printf("%T\n",b1) //*bytes.Buffer
	fmt.Printf("%T\n",b1.String()) //string*/

	/*c := make(chan int,2)
	c <- 10
	c <- 20
	close(c)
	v,ok := <-c
	fmt.Println(v,ok)  // 10 true
	v,ok = <-c
	fmt.Println(v,ok) // 20 true
	v,ok = <-c
	fmt.Println(v,ok) // 0 false*/

	/*ch := make(chan int64,2)
	ch <- 100
	ch <- 200
	defer close(ch)
	val := <- ch
	fmt.Println(val)
	val2 := <- ch
	fmt.Println(val2)*/

	/*
	//defer先入后出，先输出1，再输出0
	i := 0
	defer fmt.Println(i) //输出0，因为i此时就是0
	i++
	defer fmt.Println(i) //输出1，因为i此时就是1*/

	/*ch := make(chan int)
	go func() {
		v := <-ch
		fmt.Println(v)
	}()
	//把ch<-1这一行代码放到子线程代码的后面,channel在主线程中被赋值后，主线程就会阻塞，直到channel的值在子线程中被取出
	ch <- 1
	fmt.Println("2")*/

	ch := make(chan int)
	go produce(ch)
	go consumer(ch)
	time.Sleep(1 * time.Second)
}

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
