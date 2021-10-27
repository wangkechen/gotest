package main

//闭包是指内层函数引用了外层函数中的变量或称为引用了自由变量的函数，其返回值也是一个函数

/*func closure(x int) (func(),func(int))  {
	fmt.Printf("初始值x为: %d,内存地址: %p\n",x,&x)
	f1 := func() {
		x = x +5
		fmt.Printf("x+5: %d, 内存地址: %p\n",x,&x)
	}
	f2 := func(y int) {
		x = x +y
		fmt.Printf("x+%d: %d, 内存地址: %p\n",y,x,&x)
	}
	return f1,f2
}
//闭包来访问函数中的局部变量（被访问操作的变量为指针指向关系，操作的是同一个局部变量）
//初始值x为: 10,内存地址: 0xc00018a1b0
//x+5: 15, 内存地址: 0xc00018a1b0
//x+10: 25, 内存地址: 0xc00018a1b0
//x+20: 45, 内存地址: 0xc00018a1b0
//初始值x为: 20,内存地址: 0xc00008e060
//x+5: 25, 内存地址: 0xc00008e060
//x+10: 35, 内存地址: 0xc00008e060
//x+20: 55, 内存地址: 0xc00008e060
func main() {
	func1,func2 := closure(10)
	func1()
	func2(10)
	func2(20)
	func3,func4 := closure(20)
	func3()
	func4(10)
	func4(20)
}*/

/*func outer(x int) func(int) int  {
	return func(y int) int {
		return x + y
	}
}
//输出结果：110
func main() {
	f := outer(10)
	fmt.Println(f(100))
}*/