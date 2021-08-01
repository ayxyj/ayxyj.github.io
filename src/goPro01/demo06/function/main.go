package main

import(
	"fmt"
	"unsafe"
)

//注意：包：
/*
1)给一个文件打包，该包对应的文件夹，比如function文件夹对应的包名就是function，文件的包名通常和文件所在的文件夹名一致，一般为小写字母
2)当一个文件要使用其他包函数和变量时候，需要先引入对应的包
3)为了让其他包能够访问本包函数，函数名首字母要大写（类似public）
4)访问其他包函数变量时候，语法是 包名.函数名
5)如果包名较长们可以给包去别名 alias "package" 注意：取别名后，原包名在该文件下就不能用了
6)在同一包下，不能有相同的函数名(也不能有相同的全局变量名)，否则报重复定义
7)如果你要编译成一个可执行程序文件，就需要将这个包声明为main, 即package main。这个就是一个语法规范，如果你是写一个库，包名可以自定义
*/

//注意：函数：
/*
1)函数参数列表 返回值可以是多个
2)形参和实参 ，值传递和地址传递，是否改变实参value
3)函数命名遵循标识符规范，首字母大写，可以被其他包访问
4)Golang不支持函数重载
5)Golang中函数也是一种数据类型
*/
func main()  {
	a := sum1
	fmt.Println(myFunc( a , 1 , 2 ))//ok
	// b := sum2 
	//fmt.Println(myFunc( b , 1 , 2 ))//error
}

//定义函数类型
type mySum func( int , int ) int 
func sum1( n1 , n2 int ) int {
	return n1 + n2
}
func sum2( n1 , n2 , n3 int ) int {
	return n1 + n2
}
//使用type自定义的数据类型来简化定义 , 函数类似的数据类型作为参数使用
func myFunc(funcvar mySum , n1 , n2 int ) int {
	return funcvar( n1 , n2 )
}

//-------------------------------------------------

//函数也是一种数据类型
func test2(){
	//函数也是一种数据类型
	//type可以自定义数据类型
	type i int
	type ss func(int , int) int
	t  := test1
	t(1 , 2  , 3 , 4)//调用test6函数
	var a i = 100
	var b int = 100
	fmt.Printf("i defined = %d , %T , %d\n" , a , a , unsafe.Sizeof(a))
	fmt.Printf("i defined = %d , %T , %d\n" , b , b , unsafe.Sizeof(b))
}
func test1(sum int , args... int) int {
	//可变参数使用
	for i , j := range args {
		sum += j
		fmt.Printf("%d , %d \n", i , j)
	}
	return sum
}