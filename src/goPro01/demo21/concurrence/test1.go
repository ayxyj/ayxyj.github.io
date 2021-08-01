package main 


import (
	"fmt"
	"time"
)

func main(){
	go test()
	//主goroutine 休眠1s
	//下面语句如果没有，大多数情况是没有看到test中的打印结果
	//因为在 主goroutine(即main函数运行的goroutine)运行结束而test的goroutine还没执行
	//由于在GO主goutine执行结束，代表程序运行结束
	time.Sleep(time.Second)
}

func test()  {
	fmt.Println("work in signal goroutine")
}