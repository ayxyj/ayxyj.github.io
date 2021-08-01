package main

import (
	"fmt"
)

func main()  {
	test110()
}
type students struct{
	username string 
	password string 
	age 	 int 
	sex      string
} 
func test110()  {
	
	s1 := new(students)
	s1.username = "zzu1"
	(*s1).password = "123456"
	s2 := students{"zzu","123456",18,"m"}
	fmt.Println(s1)
	fmt.Println(s2)
	test111(s2)
	fmt.Println("值传递：",s2)
	test112(&s2)
	fmt.Println("引用传递：",s2)
}


//引用传递
func test112(stu *students)  {
	stu.username = "test111"
	fmt.Println(stu)
}
//值传递
func test111( stu students )  {
	stu.username = "test111"
	fmt.Println(stu)
}