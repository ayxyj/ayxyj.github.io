package main 

import (
	"fmt"
)


func Add( a , b  int )( res int ){
	return a + b
}

func Sub(  a , b int )( res int ){
	return a - b
}

//函数类型
type TwoFunc func( int , int ) int 

func main1(){
	res := Calc( 1 , 1 , Add)
	fmt.Println( "res1 = " , res )
	res = Calc( 1 , 1 , Sub)
	fmt.Println( "res2 = " , res )
}

//回调函数，函数有一个参数是函数类型，这个函数就是回调函数
//计算器，可以进行四则运算
//多态，多种形态，调用同一个接口，不同的表现，可以实现不同表现，加减乘除
//现有想法，后面再实现功能

func Calc(a , b int , TwoParam TwoFunc) (res int){
	res = TwoParam( a , b )
	return 
}

func test123(){
	var TwoParam TwoFunc // 声明函数变量
	TwoParam = Add //是变量就可以复制 将add赋值 ， 类似c中的指针
	res := TwoParam( 1 , 1 )
	fmt.Println( "res1 = " , res )
	//赋值sub
	TwoParam = Sub
	res = TwoParam( 2 , 1 )
	fmt.Println( "res2 = " , res )
}