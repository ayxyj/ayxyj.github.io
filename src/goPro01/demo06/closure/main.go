package main

import (
	"fmt"
	"strings"
)
//闭包:一个函数和与其相关的引用环境组合的一个整体（实体） closure
//累加器 
func AddUpper() func (int) int{
	var n int = 10
	return func ( x int ) int{
		n += x
		return n
	}
	//返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成了一个整体，构成闭包
	//可以理解闭包就是类，函数是操作，n是字段。函数和他使用到n构成一个闭包
}
func main()  {

	//example 1 
	f := AddUpper()
	fmt.Printf("%d, ",f(1))
	fmt.Printf("%d, ",f(2))
	fmt.Printf("%d, ",f(3))
	//example 2
	ff := makeSuffix(".jpg")//返回值是一个函数
	fmt.Println(ff("hello"))
	fmt.Println(ff("hello.jpg"))
	//我们体会一下闭包的好处，如果使用传统的方法，也可以轻松实现这个功能，但是传统方法需要每次都传入后缀名，比如jpg ,而闭包因为可以保留上次引用的某个值，所以我们传入一次就可以反复使用。大家可以仔细的体会吧!
}
//------------------------------------------------------------------------------------
//文件名看是否有后缀格式，没有则加上
func makeSuffix( suffix string ) func ( string ) string {
	return func ( name string ) string {
		if strings.HasSuffix( name , suffix ){
			return name 
		}else{
			return name + suffix
		}
	}
}