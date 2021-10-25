package inittest

import "fmt"

//引入包测试init 执行顺序
func init()  {
	//fmt.Println("Init Package")
}

func Wang()  {
	fmt.Println(">>>>")
}