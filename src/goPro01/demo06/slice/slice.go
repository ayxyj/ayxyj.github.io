package main

import (
    "fmt"
)

func main() {
	test2()
}

/**
数组声明  语法:[n]T
创建分片  语法:[]T
*/
//=====================================
func test7()  {
	pls := [][]string {
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
		}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
//=====================================
/**
slice作为参数传递
你可以认为slice代表的是一个结构类：

type slice struct {
    Length        int
    Capacity      int
    ZerothElement *byte
}

slice含有三个部分： 长度、容量、指向零值元素的数组。
当把一个slice传递给函数之后，虽然它是值传递，但是指针变量将指向同一个底层数组。
因此当slice作为参数传递给函数之后，在函数内部对此slice做出的修改，在此函数的外部也是可见的。我们可以写个程序检查一下：
*/
func subtactOne(numbers []int) {
    for i := range numbers {
        numbers[i] -= 2
    }
}
func subtactTwo( a [3]int ){
	for i,_ := range a{
		a[i] -= 2
	}
}
//当把一个slice传递给函数之后，虽然它是值传递，但是指针变量将指向同一个底层数组。
//因此当slice作为参数传递给函数之后，在函数内部对此slice做出的修改，在此函数的外部也是可见的。
//如果你把一个数组传递函数，然后在函数内部对此数组进行的修改，在该函数外部则不可见。
func test6() {
	nos := []int{8, 7, 6}
	a   := [3]int{3, 4, 5}
    fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
    fmt.Println("slice after function call", nos) //modifications are visible outside
	subtactTwo(a)
    fmt.Println("array after function call", nos) //modifications are invisible
}
//=====================================
func test5()  {
	veggies := []string{"potatoes","tomatoes","brinjal"}
    fruits := []string{"oranges","apples"}
    food := append(veggies, fruits...)//使用...运算符可以把一个slice追加到另一个slice中
    fmt.Println("food:",food)
}
//=====================================
func test4()  {
    var names []string //zero value of a slice is nil  // slice
    var ages [3]int  //array
    fmt.Println("names:",names,"ages:",ages)
    if names == nil {
        fmt.Println("slice is nil going to append")
        names = append(names, "John", "Sebastian", "Vinay")
        fmt.Println("names contents:",names)
    }
}
//=====================================
/**
向slice中添加内容
我们都知道数组的长度都是固定的，不能增加。而slice是动态的，新元素可以使用append函数被追加到slice中。
append的函数定义是：func append(s []T, x ...T) []T。

X...T表示的是： 该函数可以接收多个参数X。这种类型的函数被称为：可变函数。

你可以在心里会有这样一个疑问，既然slice的内部机制是由底层的数组所支持的，而数组的长度是固定的。
那么为何slice的长度是可变的呢？ 其实底层是这样的，当一个元素被追加到slice中时，是再次创建一个新
数组，原有数组里面的元素被会copy到新数组中，同时返回一个指向新数组的slice。此时，新slice的容量
为老slice容量的二倍，是不是很神奇？下面的程序将会帮你看清这一点：
*/
func test3()  {
	cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
    cars = append(cars, "Toyota")
    fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
}
//=====================================
//cap() 容量
//len() 长度
func test2()  {
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fmt.Printf("length of slice %d capacity %d\n", len(fruitarray), cap(fruitarray)) //length of is 7 and capacity is 7
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
    fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
}
//a[start:end]这个语法，会从数组a的角标start到end-1里创建一个slice。
//slice共享一个底层数组时，slice做出的改变将影响到底层的数组。
//a[:]由于没有指定起始脚标和终止脚标，它默认会认起始脚标为0，终止脚标为len(a)。

//=====================================
/**
使用make创建slice
func make([]T,len,cap) 可以用来创建slice，接收三个参数： type, length,capacity。
capacity参数是可选的，如果不传的话，默认值为数组长度。make函数会创建一个array，并返回一个指向它的slice。
*/
func test1() {
    i := make([]int, 5, 5)
    fmt.Println(i)
}
