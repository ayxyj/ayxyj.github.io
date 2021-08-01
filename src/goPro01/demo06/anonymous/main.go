package main

import(
	"fmt"
)

//匿名函数：
/*
Go支持匿名函数，匿名函数就是没有名字的函数，如果我们某个函数只是希望使用一次，可以考
虑使用匿名函数，匿名函数也可以实现多次调用。

*/
//全局匿名函数
//匿名函数赋给一个全局变量，这个函数变量可以在全局使用类似函数别名
var (
	Fun1 = func (n1 , n2 int ) int {
		return n1 + n2 
	}
)
func main()  {
	//局部匿名函数
	//将匿名函数赋给一个变量（函数变量） ， 在通过该变量来调用匿名函数
	a := func( n1 , n2 int ) int {
		return n1 + n2 
	}
	//局部匿名直接使用
	res := func(n1 , n2 int ) int {
		return n1 + n2 
	}(1, 2)
	fmt.Printf(" %d , %d , %d" , a(1 , 2) , Fun1(1 , 2) , res)
}