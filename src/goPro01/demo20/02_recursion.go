package main

import(
	"fmt"
)

func main(){

}

func test02( a int ){
	if a==1  {
		fmt.Println("a=" , a)
		return 
	}
	test02(a-1)
	fmt.Println("a=" , a)
}

func test01(){
	//递归实现累加
	sum := test(1)
	fmt.Println("sum = " , sum)
	//递归打印
	test02(3)
}

func test( i int ) ( tmp int ){
	if i == 100 {
		return 100
	}
	return i + test(i+1)
}