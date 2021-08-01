package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

//常量
type Weekday int  
const(
	Sunday  Weekday = iota 
	Monday 
	Tuesday
	Wednesday
	Friday
	Saturday
)
type Flags int 
const(
	
)

func main()  {
	fmt.Printf("%x\n",15)
	fmt.Println("x=",strconv.Itoa(15))
	val , _ := strconv.Atoi("123")
	fmt.Println("x=",val)

	//一个 float32 类型的浮点数可以提供大约 6 个十进制数的精度，而 float64 则可以提供约 15 个十进制数的精度，通常应该优先使用 float64 类型
	//因为 float32 类型的累计计算误差很容易扩散，并且 float32 能精确表示的正整数并不是很大。
	var f float32 = 16777216 // 1 << 24     
	var f2 float32 = 1677721 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
	fmt.Println(f2 == f2+1)    // "false"!
	fmt.Println(f,f+1)
	fmt.Println(f2,f2+1)
	fmt.Printf("%b",f)

	//Go语言中不允许将整型强制转换为布尔型
	//var n bool 
	//fmt.Println(int(n) * 2)


	//由于 Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节，因此使用 len() 获得两个中文文字对应的 6 个字节。
	//如果希望按习惯上的字符个数来计算，就需要使用 Go 语言中 UTF-8 包提供的 RuneCountInString() 函数，统计 Unicode 字符数量。
	fmt.Println(len("忍者"))
	fmt.Println(utf8.RuneCountInString("忍者"))
	//也可以通过切片来实现
	str := []rune("忍者")
	fmt.Println(len(str))
	fmt.Println("----slice----rune----")
	//包含汉字的处理方法
	//ASCII 字符串遍历直接使用下标。
	//Unicode 字符串遍历用 for range。
	var str1  = "忍者"
	for  index , val := range str1{
		fmt.Printf("%d,%c\n",index , val)
	}
	str2 := []rune(str1) //切片
	for  index , val := range str2{
		fmt.Printf("%d,%c\n",index , val)
	}
	fmt.Println("----slice----rune----")
	fmt.Println(utf8.RuneCountInString("龙忍出鞘,fight!"))
	//Go 语言的字符串是不可变的。
	// 修改字符串时，可以将字符串转换为 []byte 进行修改。
	// []byte 和 string 可以通过强制类型转换互转。
	angel := "Heros never die"
	angleBytes := []byte(angel)
	for i := 5; i <= 10; i++ {
		angleBytes[i] = ' '
	}
	fmt.Println(string(angleBytes))
}