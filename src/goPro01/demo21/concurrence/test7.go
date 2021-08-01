package main

import (
	"fmt"
	"time"
)

func main()  {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Send(ch1 ,0)
	go Send(ch2 , 10)

	//主goroutine休眠1s保证调度成功
	time.Sleep(time.Second)

	for{
		//select 中的case 后语句必须是I/O操作 在没有响应且没有提供default语句，goroutine 将会被阻塞
		select{
		case val := <- ch1 : 
			fmt.Printf("get channel ch1 value : %d\n",val)
		case val := <- ch2 :
			fmt.Printf("get channel ch2 value : %d\n",val)
		case <-time.After(2 * time.Second) : //超时设置
			fmt.Println("time out !")
			return
		}
		//select 多路复用从ch1 ch2 中读取数据，如果多个case的ch同时到达，那么select会运行伪随机算法，随机选择一个case
	}
}

func Send( ch chan int , x int )  {
	for i := x ; i < x + 10; i++ {
		ch <- i 
	}
}