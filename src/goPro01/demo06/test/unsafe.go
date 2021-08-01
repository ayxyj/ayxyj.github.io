package main

import(
	"fmt"
	"unsafe"
)

func main()  {
	var a int //随系统位数 64b == 8B
	var b int32
	fmt.Printf("%T,%d\n",a,unsafe.Sizeof(a))
	fmt.Printf("%T,%d\n",b,unsafe.Sizeof(b))

	fmt.Println("------------------------")
	var name = "Sam"
	const sex = "m"
	const trueConst = true
    fmt.Printf("type %T value %v\n", name, name)
	fmt.Printf("type %T value %v\n", sex, sex)
	fmt.Printf("type %T value %v\n", trueConst, trueConst)
}