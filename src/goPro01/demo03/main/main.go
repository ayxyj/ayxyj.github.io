package main

import(
	"fmt"
	"goPro01/demo03/model"
	//环境变量中配置了 GOPATH/src 因此引用包只需要写到src下一级即可
	//D:\Go\GoProject\src\goPro01\demo03\model (from $GOPATH)go
)
func main()  {
	/*
		如果变量名、函数名、常量名首字母大写，则可以被其他的包访问;
		如果首字母小写，则只能在本包中使用（注:可以简单的理解成，首字母大写是公开的，首字母小写是私有的)
		在golang没有public , private等关键字。
	*/
	fmt.Println("英雄名字：", model.HeroName)
	//fmt.Println("英雄性别：", model.sex)  //.\main.go:16:33: cannot refer to unexported name model.sex 错误 类似于java中private
}