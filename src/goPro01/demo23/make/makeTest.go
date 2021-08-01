package main 

import(
	"fmt"
	"time"
)

func main()  {
	test111()

}

func test112(maps map[int]string)  {
	// 在函数间传递映射并不会制造出该映射的一个副本，不是值传递，而是引用传递：

	for key , val := range maps {
		fmt.Println(key , ":" , val)
	}
	delete(maps, 2) 
}

func test111()  {
	m1 := map[int]string{1: "mike", 2: "yoyo", 3: "lily"}
    //迭代遍历1，第一个返回值是key，第二个返回值是value
	test112(m1)

    for k, v := range m1 {
        fmt.Printf("%d ----> %s\n", k, v)
	}
}

func test110()  {
	//注意：make只能创建slice、map和channel，并且返回一个有初始值(非零)。
	maps := make(map[string]string)
	sli := make([]string , 10 , 10 )
	chan1 := make(chan string , 10)
	maps["name"]="zzu"
	maps["sex"]="m"
	sli[0] = "test01"
	sli[1] = "test02"
	chan1 <- "test03"
	chan1 <- "test04"
	fmt.Println(maps)
	fmt.Println(sli ,  "   " , len(sli) ,  "   " , cap(sli))
	fmt.Println(<-chan1 ,  "   " , len(chan1), "   ", cap(chan1) )//test03            FIFO
	go PrintChan(chan1)//test04

	//append函数向 slice 尾部添加数据，返回新的 slice 对象
	s2 := make([]int, 5)
	fmt.Println(len(s2) , cap(s2))
    s2 = append(s2, 6)
	fmt.Println(len(s2) , cap(s2))
    fmt.Println(s2) //[0 0 0 0 0 6]

	time.Sleep(time.Second * 10)

}

func PrintChan(chan1 chan string)  {
	for v := range chan1 {
		fmt.Printf("%s",v)	
	}
}