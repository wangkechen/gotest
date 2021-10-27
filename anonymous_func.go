package main

//匿名函数的两种使用方式
//一、在定义匿名函数的时候就可以直接使用（这种方式只使用一次）
/*func main(){
	res1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10, 30)  //括号里的10,30 就相当于参数列表，分别对应n1和n2

	fmt.Println("res1=",res1)
}*/

//匿名函数的两种使用方式
//二、将匿名函数赋给一个变量（函数变量），再通过该变量来调用匿名函数
/*func main(){
	//将匿名函数fun 赋给变量test_fun
	//则test_fun的数据类型是函数类型，可以通过test_fun完成调用
	test_fun := func (n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := test_fun(10, 30)
	res3 := test_fun(50, 30)
	fmt.Println("res2=", res2)
	fmt.Println("res3=", res3)
	fmt.Printf("%T", test_fun)
}*/