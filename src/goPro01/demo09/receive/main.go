package main

import (
	"fmt"
)
var m  map[int]users
var count int 
//object
type users struct{
	username string
	password string 
}

//对象方法
func (u *users) insert( username string , password string ) {
	u.username = username
	u.password = password
	count ++
	m[count] = *u
}

func main()  {
	fmt.Println("----接收器----")	
	m = make(map[int]users)
	u := new(users)
	u.insert("xiayujie1" , "123456")
	u.insert("xiayujie2" , "123456")
	u.insert("xiayujie3" , "123456")
	for _,v := range m {
		fmt.Println(v)
	}
}
/**
接收器——方法作用的目标
接收器的格式如下：
func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    函数体
}
对各部分的说明：
接收器变量：接收器中的参数变量名在命名时，官方建议使用接收器类型名的第一个小写字母，而不是 self、this 之类的命名。例如，Socket 类型的接收器变量应该命名为 s，Connector 类型的接收器变量应该命名为 c 等。
接收器类型：接收器类型和参数类似，可以是指针类型和非指针类型。
方法名、参数列表、返回参数：格式与函数定义一致。
*/