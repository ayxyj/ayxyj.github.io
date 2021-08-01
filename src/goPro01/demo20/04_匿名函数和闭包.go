package main

import(
	"fmt"
)

func main1(){

}

func test11(){
	a , str := 10 ,"go"
	//匿名函数，没有函数名字
	//函数定义，没有调用
	f1 := func(){
		fmt.Println("a=",a)
		fmt.Println("str=",str)
	}
	//调用
	f1()

	//给一个函数类型起别名
	type FuncType func()
	//声明变量
	var f2 FuncType
	f2 = f1 
	//调用
	f2()

	//定义一个匿名函数，同时进行调用
	func(){
		fmt.Printf("a=%d,str=%s\n" , a , str)
	}()

	//带参数的匿名函数并调用
	func(a , b int ){
		fmt.Printf("a=%d , b=%d\n" ,a ,b)
	}( 1 , 2)

	//有返回值的匿名函数
	max , min := func(a , b int )(max ,min int ){
		if a>b{
			max = a 
			min = b
			}else{
				max = b 
				min = a 
		}
		return
	}(10,20)
	fmt.Printf("max=%d ,min=%d\n" , max , min )
}