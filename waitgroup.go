package main

//https://blog.csdn.net/u013474436/article/details/88749749

//主线程为了等待goroutine都运行完毕，不得不在程序的末尾使用time.Sleep()
//来睡眠一段时间，等待其他线程充分运行。对于简单的代码，100个for循环可以在1秒之内运行完毕，
//time.Sleep() 也可以达到想要的效果。
/*func main()  {
	for i:=0;i<100;i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Second)
}*/

//使用管道是能达到我们的目的的，而且不但能达到目的，
//还能十分完美的达到目的。 但是管道在这里显得有些大材小用，
//因为它被设计出来不仅仅只是在这里用作简单的同步处理，
//在这里使用管道实际上是不合适的。而且假设我们有一万、十万甚至更多的for循环，
//也要申请同样数量大小的管道出来，对内存也是不小的开销。
/*func main()  {
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}
	for i:= 0;i<100;i++{
		<-c
	}
}*/

//go语言中有一个其他的工具sync.WaitGroup 能更加方便的帮助我们达到这个目的。
//WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait()
//用来控制计数器的数量。Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，
//wait() 会阻塞代码的运行，直到计数器地值减为0。
/*func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}*/