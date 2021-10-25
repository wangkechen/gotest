package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func stringtoslicebyte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func slicebytetostring(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func main1()  {
	/*var r1 rune = '\u0041'
	var r2 rune = '\U00000041'
	fmt.Printf("%c %c \r\n",r1,r2)
	r3 := '中'
	r4 := '\u4E2D'
	r5 := []byte{'h','e','l','l','\xc3','\xb8'}
	fmt.Printf("%U %c %s \r\n",r3,r4,r5)
	str := []rune(string("养由基"))
	for _,r := range str {
		fmt.Printf("%U\r\n",r)
	}
	fmt.Println(string([]rune{0x517B,0x7531,0x57FA}))
	fmt.Println(string([]rune{'\u517B','\u7531','\u57FA'}))*/

	//a := []byte{1,2,3}
	//fmt.Println(a)  //[1 2 3]

	//a := []byte("1,2,3")
	//fmt.Println(a)  //[49 44 50 44 51]

	/*var s1,s2 string
	for i := 0; i < 5; i++ {
		//string函数的参数若是一个整型数字，它将该整型数字转换成ASCII码值等于该整形数字的字符。
		s1 += string(i)
		//将数字转换成对应的字符串类型的数字
		s2 += strconv.Itoa(i)
		//fmt.Println(s1,s2)
	}
	a1 := []byte(s1) //[0 1 2 3 4]
	a2 := []byte(s2) //01234  => [48 49 50 51 52]
	fmt.Println(a1,a2)*/

	/*s1 := "abc"
	b1 := []byte("def")
	copy(b1, s1)  //将s1复制到b1
	log.Println(s1, b1) //abc [97 98 99]

	s := "hello"
	b2 := stringtoslicebyte(s) //等同于 b2 := []byte(s)
	log.Println(b2) //[104 101 108 108 111]
	// b2[0] = byte(99) unexpected fault address

	b3 := []byte("test")
	s3 := slicebytetostring(b3)
	log.Println(s3) //test*/

	/*str2 := "hello"
	data2 := []byte(str2)
	fmt.Println(data2)
	str2 = string(data2)
	fmt.Println(str2)*/

	/*type Person struct {
		Name string
		Age int64
	}
	person := Person{"wang",25}
	value := reflect.ValueOf(person)
	fmt.Printf("%v\n", value.FieldByName("Name").Interface())*/

	/*str := "-100 123 200"
	//指定分隔符
	countSplit := strings.Split(str, " ")
	fmt.Println(countSplit, len(countSplit)) //[-100 123 200] 3
	//指定分割符号，指定分割次数
	countSplit = strings.SplitN(str, " ", 2)
	fmt.Println(countSplit, len(countSplit)) //[-100 123 200] 2*/

	dbLicensePlate := "闽BK0110-京QZ7566"
	input := "闽BK0110"
	//plateByte := strings.Contains(dbLicensePlate,input) //判断字符串是否包含
	plateByte := strings.Split(dbLicensePlate,"-")
	var flag bool
	fmt.Println(flag)
	if len(plateByte) > 0 {
		for _,v := range plateByte{
			if v == input {
				fmt.Printf("%T %v\n", v, v)
				flag = true
			}
		}
	}
	fmt.Println(flag)

	//fmt.Printf("%T\n", plateByte)


}