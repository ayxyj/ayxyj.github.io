package main 

import (
	"fmt"
	"time"
	"sync"
)
//互斥锁
func main()  {
	//按照逻辑实现顺序执行
	var lock sync.Mutex //互斥锁

	go func(){
		lock.Lock()//加锁
		defer lock.Unlock()//解锁 
		fmt.Println("func1 time start :", time.Now().String())
		time.Sleep(time.Second)
		fmt.Println("func1 time end :", time.Now().String())
	}()
	time.Sleep(time.Second)
	go func(){
		lock.Lock()//加锁
		defer lock.Unlock()//解锁 
		fmt.Println("func2 time start :", time.Now().String())
		time.Sleep(time.Second)
		fmt.Println("func2 time end :", time.Now().String())
	}()
	time.Sleep(time.Second * 4)
}

/**
互斥锁保证每次只有一个goroutine访问同步代码块中的资源，即访问同一资源
PS D:\Go\GoProject\src\goPro01\demo21\concurrence> go run .\test8.go
func1 time start : 2021-03-14 16:08:23.8317599 +0800 CST m=+0.004986401
func1 time end : 2021-03-14 16:08:24.8540267 +0800 CST m=+1.027253201
func2 time start : 2021-03-14 16:08:24.8540267 +0800 CST m=+1.027253201
func2 time end : 2021-03-14 16:08:25.8545308 +0800 CST m=+2.027757301
*/
