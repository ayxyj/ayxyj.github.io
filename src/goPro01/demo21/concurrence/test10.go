package main

import (
	"fmt"
	//"bufio"
	"sync"
	"time"
	"strconv"
)

/**
//添加等待数据,传递负数表示任务减1
func (wg *WaitGroup) Add(delta int)
//等待数减1
func (wg *WaitGroup) Done()
//使用goroutine等待于此
func (wg *WaitGroup) Wait()

注:
1、在goroutine调用waitGroup.Wait进行等待之前，需要保证waitGroup中等待数量大于1，即waitGroup.Add 在 waitGroup.Wait之前执行，否则就会被忽略
2、waitGroup.Done执行次数和waitGroup.Add添加的数量保持一致，过少会导致死锁，过多会导致程序的panic
*/

//sync.WaitGroup
//并发等待组适用于执行批量操作，等待所有的goroutine执行结束后统一返回结果的情况
func main()  {
	var waitGroup sync.WaitGroup
	//添加等待数量 goroutine为5
	waitGroup.Add(5)
	
	
	fmt.Println("start : ", time.Now().String())

	for i := 0; i < 5; i++ {
		go func(i int){
			fmt.Println("work "+strconv.Itoa(i)+" is done at "+time.Now().String())
			time.Sleep(time.Second)
			//减少等待数
			waitGroup.Done()
		}(i)
	}
	//使goroutine等待于此 ，等待所有的goroutine全部完成
	waitGroup.Wait()
	fmt.Println("end : ", time.Now().String())
}

/**
//主goroutine在执行waitGroup.Done之后，需要等待waitGroup中的等待数变为0之后才继续往后执行
PS D:\Go\GoProject\src\goPro01\demo21\concurrence> go run .\test10.go
start :  2021-03-14 16:39:30.0544475 +0800 CST m=+0.003988001
work 4 is done at 2021-03-14 16:39:30.0773862 +0800 CST m=+0.026926701
work 0 is done at 2021-03-14 16:39:30.0773862 +0800 CST m=+0.026926701
work 2 is done at 2021-03-14 16:39:30.0773862 +0800 CST m=+0.026926701
work 1 is done at 2021-03-14 16:39:30.0773862 +0800 CST m=+0.026926701
work 3 is done at 2021-03-14 16:39:30.0773862 +0800 CST m=+0.026926701
end :  2021-03-14 16:39:31.0787087 +0800 CST m=+1.028249201
*/