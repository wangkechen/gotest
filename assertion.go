package main

import "fmt"

//类型断言是将接口类型的值x，转换成类型T
//类型断言的必要条件是x是接口类型,非接口类型的x不能做类型断言
func f(i interface{})  {
	v,ok := i.(string)
	if ok {
		fmt.Println("ok, string is:", v)
	}
	switch i.(type) {
	case string:
		fmt.Println(i,"is string")
	case int:
		fmt.Println(i,"is int")
	}
}
//ok, string is: hefs
//hefs is string
//33 is int
//类型转化失败
//0
//cc的类型是：int
/*func main() {
	i := "hefs"
	f(i)
	a := 33
	f(a)

	var bb interface{} = "string"
	cc, ok := bb.(int)
	if !ok {
		fmt.Println("类型转化失败")
		fmt.Println(cc)
	}
	fmt.Printf("cc的类型是：%T",cc)
}*/
