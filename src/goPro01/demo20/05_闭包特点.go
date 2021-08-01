package main

import(
	"fmt"
	"math"
)

func main(){
	test02()
}

func test02(){
	var e uint16 = math.MaxUint8 + 1 
	fmt.Printf("e value %v\n",e)
	var f = uint8(e)
	fmt.Printf("d value %v",f)
}

func test01() func() (res int) {
	var x int 
	return func() int {
		x++
		return x*x
	}
}

func test(){
	f1 := test01()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())

}