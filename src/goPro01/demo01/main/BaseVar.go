package main

import "fmt"

//声明全局变量
var r1  = 250 
var sex = "man"
var r2  = 30.55
var(
	q1 = 100
	pass = "123456"
	q3 = 150
)

func main() {
	test13()
}


func test13()  {
	//最大公约数
	var x ,y int
	fmt.Scanln(&x , &y)
	for y != 0{
		x , y = y , x%y
	}
	fmt.Printf("最大公约数%d\n" , x)
}
func test12()  {
	//new
	ptr := new(int)
	fmt.Printf("%v\n", ptr)
	*ptr = 2
	fmt.Printf("%v\n", *ptr)
}

func test11()  {
	fmt.Println("hello go!\n")
	//单变量声明
	var i int = 10 
	var j = 100
	k := 1000
	fmt.Println("output single var:")
	fmt.Println("i=", i , "j=" , j , "k=" , k , "\n")
	//多变量声明
	var n1 , n2 , n3 int 
	var username string
	fmt.Println("output multi1 var:")
	fmt.Println("n1=", n1 ,"n2=" , n2 , "n3=" , n3 , "\n")
	fmt.Println("username=",username,"\n")

	var t1 , name , t2 = 10 , "tom" , 100
	fmt.Println("output multi2 var:")
	fmt.Println("t1=" , t1 , "name=" , name , "t2=" , t2 , "\n")

	s1 , name1 , s2 := 10 , "tom" , 100
	fmt.Println("output multi3 var:")
	fmt.Println("s1=" , s1 , "name1=" , name1 , "s2=" , s2 , "\n")
	//全局变量
	fmt.Println("output global var:")
	fmt.Println("r1=" , r1 , "sex=" , sex , "r2=" , r2 , "\n")
	fmt.Println("q1=" , q1 , "pass=" , pass , "q2=" , r2 , "\n")
}