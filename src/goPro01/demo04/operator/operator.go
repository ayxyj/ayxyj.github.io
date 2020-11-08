package main

import(
	"fmt"
)

func main()  {

	test4()
}


//操作符
func test4()  {
	var a bool 
	fmt.Printf("%+v , %#v , %T \n",a , a , a)

	//移位运算和位运算  补码进行移位操作和位操作，然后在转回成原码（负数），正数三码一致
	var a1 int = 1 >> 2
	var b1 int =-1 >> 2
	fmt.Printf("%d\n" , a1)
	fmt.Printf("%d\n" , b1)
}
func test3()  {
	//关系运算符 > <  =   >=  <= != ==
	//逻辑运算符 && || ！
	//位运算 &  ^ | >> << 
	//赋值运算符  =  += -= *= /= %=  >>=   <<=   &=  |=  ^=
	//小练习：
	var i , j  int
	fmt.Scanln(&i , &j)
	if i != j {
		if i > j {
			i &= j
			fmt.Println("i &= j ：" , i )
		}else{
			i ^= j
			fmt.Println("i &= j ：" , i )
		}
	}
	//不借助第三个变量 实现交换ab
	var a , b  int 
	fmt.Scanln(&a , &b)
	fmt.Printf("交换前：a = %d , b = %d\n", a , b)
	a = a ^ b
	b = a ^ b
	a = b ^ a
	fmt.Printf("交换后：a = %d , b = %d\n", a , b)

	var t1 , t2  uint
	fmt.Scanln(&t1 , &t2)
	//.\operator.go:47:65: invalid operation: t1 >> t2 (shift count type int, must be unsigned integer)
	//.\operator.go:47:76: invalid operation: t1 << t2 (shift count type int, must be unsigned integer)
	fmt.Println("位运算：" , t1 & t2 , t1 ^ t2  , t1 | t2 , t1 >> t2 , t1 << t2)
}
	
func test2()  {
	//小练习1)：假期还有68天，还能玩几个星期零几天
	fmt.Printf("还能玩%d个星期 %d天\n" , 68/7 , 68%7)
	//小练习2): 定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为:5/9*(华氏温度-32),请求出华氏温度对应的摄氏温度。
	var htem float64 //华氏温度
	fmt.Scanln(&htem)
	var stem = 5.0 / 9 * (htem - 32)
	fmt.Printf("输入的华氏温度%5.2f,对应的摄氏温度%5.2f\n",htem,stem)
}

func test1()  {
		/*
	>案例演示算术运算符的使用。
	+,-,*,/,%,++,--，重点讲解/、%
	自增:++
	自减:--
	*/
	fmt.Println("10%3="   , 10  %  3)
	fmt.Println("-10%3="  ,-10  %  3)
	fmt.Println("10%-3="  , 10  % -3)
	fmt.Println("-10%-3=" ,-10  % -3)

	var count int = 5;
	for i := 0; i < count; i++ {
		fmt.Println("hello Go！")
	}
}