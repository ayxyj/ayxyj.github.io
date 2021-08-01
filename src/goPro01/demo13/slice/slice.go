package main

import(
	"fmt"
	"unicode/utf8"
)

func main()  {
	test02()
	test01()
}

func test02()  {
	//arr
	//define:
		//var name [size]T  
		//var name [...]T(.......)
		//var name = new([size]T)
	var name [5]string
	var name1 = [...]string {"zzu" , "ayit" , "sdust"}
	name2 := new([5]string)

	name[0] = "name:zzu"
	name2[0] = "name2:zzu"
	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(*name2)

	fmt.Println("\n---------------------------------------")
	//slice
	//define:
		//arrayname[start:end]
		//make([]T , size , cap)
		//var name = []T{}

		
	var name4  =[...]string{"hello","Java","GO","Python"}
	fmt.Println(len(name4))//4

	sli1 := name4[0:3]
	fmt.Println(sli1) //[hello Java GO]
	fmt.Println(len(sli1))//3
	fmt.Println(cap(sli1))//4 等于name4长度

	sli4 := name4[:]
	fmt.Println(sli4) //[hello Java GO Python]
	fmt.Println(len(sli4))//4 
	fmt.Println(cap(sli4))//4 等于name4长度

	var sli2 = make([]string , 5 , 15)
	sli2[0] = name4[0]
	fmt.Println(sli2)//[hello]
	fmt.Println(len(sli2))//5
	fmt.Println(cap(sli2))//15

	var sli3 = []string{name4[0]}
	fmt.Println(sli3)//[hello]
	fmt.Println(len(sli3))//1
	fmt.Println(cap(sli3))//1

	//function
	//append copy
	newSli3 := append(sli3 , name4[1] ,name4[2])
	fmt.Println(newSli3)//[hello Java]
	fmt.Println(len(newSli3))//2
	fmt.Println(cap(newSli3))//2

	//var new1sli3 []string  //[] 0 0 因为新创建得切片长度和容量都是0,所有当长度和容量不够时，复制得个数和新建切片相同
	var new1sli3 = make([]string , 5 , 10)
	copy(new1sli3, sli3)
	fmt.Println(new1sli3)//[hello Java]
	fmt.Println(len(new1sli3))//5
	fmt.Println(cap(new1sli3))//10
	fmt.Println("\n---------------------------------------")
}
func test01()  {
	var name string ="郑州大学zzu"
	fmt.Println(len(name))//3 * 4 + 3  
	fmt.Println(utf8.RuneCountInString(name)) // 7
	for _, v := range name {
		fmt.Printf("  %c",v)
	}
	fmt.Println()
	for _, v := range []byte(name) {
		fmt.Printf("  %c",v)
	}
	fmt.Println()
	var newName = []rune(name) //切片  []T(string)
	for _, v := range newName {
		fmt.Printf("  %c",v)
	}
	fmt.Println("\n---------------------------------------")
}