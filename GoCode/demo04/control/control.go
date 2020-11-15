package main

import(
	"fmt"
	"math/rand"
	"time"
	"math"
)
//程序控制语句
func main()  {


	
}
func test10()  {
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.NaN())
	var z float64
	fmt.Println(math.IsNaN(z/z))
	fmt.Println(math.IsNaN(1/z))
	fmt.Println(1/z)
	fmt.Println(-1/z)
	/*
	3.4028234663852886e+38
	1.7976931348623157e+308
	NaN
	true
	false
	+Inf
	-Inf
	*/
}


var pc [256]byte
func init()  {
	for i := range pc{
		pc[i] = pc[i/2] + byte(i&1)
	}
	fmt.Println(pc)
}

func PopCount(x uint64)  int {
	return int(
		pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))] ) 
}



func test9()  {

	//交换1
	i , j := 10, 20
	fmt.Printf("前：i = %d , j = %d\n",i , j)
	i , j = j , i
	fmt.Printf("后：i = %d , j = %d\n",i , j)
	//交换2
	i = i ^ j
	j = i ^ j
	i = j ^ i 
	fmt.Printf("后：i = %d , j = %d\n",i , j)

}
func test8()  {
	//countine 结束本次循环  也可以使用标签  结束标签位置的循环
	//break    结束本层循环  
	//continue实现打印1——100之内的奇数[要求使用for循环+continue]

}
func test7()  {
	//随机数
	rand.Seed(time.Now().Unix())
	n := rand.Intn(20)+1
	fmt.Printf("n=%d\n",n)
	var count int = 1 
	label1:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if n>10 {
				break label1 //跳出指定标签位置的循环  正常break是就近原则
			}else{
				fmt.Printf("n=%d\n",n)
			}
			count++
		}
	}
	fmt.Printf("count=%d\n",count)
}

func test6()  {
	//for嵌套  9*9
	for i :=1 ;  i <= 9; i++ {
		for j := 1; j <= i ; j++{
				 fmt.Printf("%d*%d=%d	",i,j,i*j)
		}
		fmt.Printf("\n")
	}
}

func test5()  {
	//for
	var str = "hello_world!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n",str[i])
	}
	//for range
	for index , val := range str{
		fmt.Printf("%d,%c\n",index , val)
	}

	//包含汉字的处理方法
	var str1  = "hello_北京！"
	str2 := []rune(str1) //切片
	for  index , val := range str2{
		fmt.Printf("%d,%c\n",index , val)
	}

	//while 和 do...while 没有  ，需要通过for来实现
	var a = 9
	for{
		if a < 10 {
			fmt.Printf("先判断，后干事！\n")
		}else{
			fmt.Printf("while!\n")
			break
		}
		a++
	}

	//for ; ;
	//while 和 do...while 没有  ，需要通过for来实现
	var b = 9
	for{
		b++
		fmt.Printf("先干事！在判断！\n")
		if b < 10 {
				fmt.Printf("do ... while1!\n")
			}else{
				fmt.Printf("do ... while2!\n")
				break
		}
	}
	
	// for ; ** ; 
	var i = 1
	for i < 2 {
		fmt.Printf("hello_world!\n")
		i++
	}
}

func test4()  {
	
	//switch
	var date int 
	fmt.Println("请输入date：")
	fmt.Scanln(&date)
	switch date {
	case 1:
		fmt.Printf("周一！\n")
		//break  默认有
		//fallthrough  //穿透 默认下一层
	case 2:
		fmt.Printf("周二！\n")
		//break
	default :
		fmt.Printf("周34567\n")
	}
}

func test3()  {
	//判断分数段
	var grade int 
	fmt.Println("请输入Grade：")
	fmt.Scanln(&grade)
	if grade >= 90 {
		fmt.Println("优秀！")
		}else if grade >=80 {
			
			fmt.Println("中等！")
			}else {		
				fmt.Println("良好或差！")
	}
}
func test2()  {
	//闰年
	var date int 
	fmt.Println("请输入年份：")
	fmt.Scanln(&date)
	if (date%4==0 && date%100 != 0) || date%400 ==0 {
		fmt.Println("闰年！")
		}else{
			fmt.Println("平年！")
	}
}
func test1()  {
	//输入两个数之和能否整除3和5、
	var a ,b int 
	fmt.Println("请输入两个数：")
	fmt.Scanln(&a,&b)
	if (a+b)%3==0  && (a+b)%5==0 {
		fmt.Printf("yes\n")
	}else{
		fmt.Printf("no\n")
	}
}