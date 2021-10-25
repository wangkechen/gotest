package main

/*func main() {
	var i1, i2, i3 int
	c1 := make(chan int,1)
	c1 <- 1
	c2 := make(chan int,1)
	c2 <- 2
	c3 := make(chan int,1)
	c3 <- 3
	select {  //随机选一个执行，除了default
	case i1 = <- c1:
		fmt.Println("received ", i1, " from c1")
	case i2 = <- c2:
		fmt.Println("sent ", i2, " to c2")
	case i3 = <- c3:
		fmt.Println("sent ", i3, " to c3")
	default:
		fmt.Printf("no communication")
	}

	fmt.Println("end")
}*/

/*func Chann(ch chan int, stopCh chan bool) {
	var i int
	i = 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}

//for select 会循环检测条件，如果有满足则执行并退出，否则一直循环检测。
func main() {

	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("Receive", c)
			fmt.Println("cccc")
		case s := <-ch:
			fmt.Println("Receive", s)
			fmt.Println("ssss")
		case _ = <-stopCh:
			fmt.Println("stop")
			goto end
		}
	}
	end:
	println("end")
}*/

