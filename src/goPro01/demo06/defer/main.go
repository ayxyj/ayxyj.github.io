package main

import(
	"fmt"
)

func main()  {
	test3()
}
//===================================
type Users struct {
    name string
}

func (t *Users) GetName() { // 注意这里是 * 传地址 引用Users
    fmt.Println(t.name)
}
func GetName(t Users) { // 定义一个函数，名称自定义
    t.GetName() // 调用结构体USers的方法GetName
}
func test3() {
    list := []Users{{"乔峰"}, {"慕容复"}, {"清风扬"}}
    for _, t := range list {
		fmt.Printf("%T\n",t)
        defer GetName(t)
    }
}
//===================================
func test2()  {
    var users [5]struct{}
	for i := range users {
			defer Print(i)
	}
}
func Print(i int) {
    fmt.Println(i)
}

//===================================
func test1() {
	/**
	输出：4 4 4 4 4，很多人也包括我。
	预期的结果不是 4 3 2 1 0 吗？
	官网对defer闭包的使用大致是这个意思：
		函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,
		所以输出全都是4。那么如何正常输出预期的 4 3 2 1 0 呢？
	*/
    var users [5]struct{}
    for i := range users {
        defer func() { fmt.Println(i) }()
    }
}
//===================================
func sum( n1 , n2 int ) int  {
	defer fmt.Println("OK1 --- n1的值：",n1)
	defer fmt.Println("OK2 --- n2的值：",n2)
	//defer压入栈
	/**
	defer特性：
		1. 关键字 defer 用于注册延迟调用。
		2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
		3. 多个defer语句，按先进后出的方式执行。
		4. defer语句中的变量，在defer声明时就决定了。
	defer用途：
		1. 关闭文件句柄
		2. 锁资源释放
		3. 数据库连接释放
	*/
	res := n1 + n2 
	fmt.Println("OK3 --- res的值：" , res )
	return res 
}