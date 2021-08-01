package main

import (
	"fmt"
	"unsafe"
)
//执行的流程： 全局变量定义 ->init函数->main函数
//init函数最主要的作用是完成一些初始化的工作
//导入包的函数首先执行



//globe var 
var t = test
func test(){
	fmt.Printf("First : hello globe var ！\n")
}

//init函数
func init(){
	t()  
	var i = 100
	fmt.Printf("Second : init %T , %d \n" , i , unsafe.Sizeof(i))
}
func main(){
	fmt.Printf("Third : hello main!")
}