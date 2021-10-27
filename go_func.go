package main

//go func(){}已并发的方式调用匿名函数func

//https://studygolang.com/articles/21588?fr=sidebar
//https://blog.csdn.net/LMFranK/article/details/107525259

/*func main()  {
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ",i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ",i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}*/

//多次运行结果都是上面一样。先输出B: 9，然后10个A: 10，然后B输出0到8
//runtime. GOMAXPROCS(1) 强行指定了只创建一个 “P” 来处理并发，这使得例子中的 20 个 goroutine 会是串行的
//编译器会把 go 后面跟着的函数与参数都打包成g对象，等待系统调度。
/*func main()  {
	//设置协程调度只有一个P
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ",i)
			wg.Done()
		}()
	}
	for i:=0;i<10;i++ {
		go func(i int) {
			fmt.Println("B: ",i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}*/

//输出结果：0 1 2 3 ... 8 9 注意输出的不是10个10
/*func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go print(i)
	}
	time.Sleep(time.Second)
}*/

//输出结果：9 0 1 2 3 4 5 6 7 8
/*func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go println(i)
	}
	//runtime.Gosched()
	time.Sleep(time.Second)
}*/