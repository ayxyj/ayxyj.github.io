package main

import(
	"fmt"
)

func main(){
	const (
		//iota 只能给常量枚举
		a = iota
		b = iota
		c = iota
	)

	const (
		//遇到const 时 iota自动清零
		d = iota
		//同一行iota同一个值
		e,f,g = iota,iota,iota
		h = iota
	)

	fmt.Printf("枚举：\n  a=%d,b=%d,c=%d \n清零：\n  d=%d \n同值：\n  e=%d ,f=%d , g=%d \n增：\n  h=%d",a,b,c,d,e,f,g,h)
}