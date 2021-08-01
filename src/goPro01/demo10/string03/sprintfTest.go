package main

import(
	"fmt"
)

func main()  {
	var progress = 2
	var target = 8

	// 两参数格式化
	title := fmt.Sprintf("已采集%d个药草, 还需要%d个完成任务", progress, target)

	fmt.Println(title)

	pi := 3.14159
	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)

	fmt.Println(variant)

	// 匿名结构体声明, 并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}

	fmt.Printf("使用'%%+v' %+v\n", profile)

	fmt.Printf("使用'%%#v' %#v\n", profile)

	fmt.Printf("使用'%%T' %T\n", profile)
}
/**
表：字符串格式化时常用动词及功能
动  词	功  能
%v	按值的本来值输出
%+v	在 %v 基础上，对结构体字段名和值进行展开
%#v	输出 Go 语言语法格式的值
%T	输出 Go 语言语法格式的类型和值
%%	输出 % 本体
%b	整型以二进制方式显示
%o	整型以八进制方式显示
%d	整型以十进制方式显示
%x	整型以十六进制方式显示
%X	整型以十六进制、字母大写方式显示
%U	Unicode 字符
%f	浮点数
%p	指针，十六进制方式显示
*/