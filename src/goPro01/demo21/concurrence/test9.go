package main 

import (
	"fmt"
	"sync"
	"time"
	//"strconv"   strconv.Itoa(int)
)

/**
//写加锁
func (rw *RWMutex) Lock()
//写解锁
func (rw *RWMutex) Unlock()
//读加锁
func (rw *RWMutex) RLock()
//读解锁
func (rw *RWMutex) RUnlock()

//读加解锁
func (rw *RWMutex) RLocker() Locker
RLocker returns a Locker interface that implements the Lock and Unlock methods by calling rw.RLock and rw.RUnlock. 
    83	func (rw *RWMutex) RLocker() Locker {
    84		return (*rlocker)(rw)
    85	}
    86	
    87	type rlocker RWMutex
    88	
    89	func (r *rlocker) Lock()   { (*RWMutex)(r).RLock() }
    90	func (r *rlocker) Unlock() { (*RWMutex)(r).RUnlock() }
*/
//读写锁 RWMutex
var rwLock sync.RWMutex
func main()  {
	var ch = make(chan int)
	fmt.Println("start write int to channel :",time.Now().String())
	for i := 0; i < 10; i++ {
		go func(x int){
			rwLock.Lock()
			defer rwLock.Unlock()
			ch <- x
			}(i)
			fmt.Println(i,"  ",time.Now().String())
		}
	time.Sleep(time.Second * 2)
	fmt.Println("data write to channel end")
	fmt.Println("end write int to channel :",time.Now().String())


	fmt.Println("start Read int to channel :",time.Now().String())
	for i := 0; i < 10; i++ {
		go func(){
			rwLock.RLocker()
			defer rwLock.RLocker()
			fmt.Println("get channel data :",<-ch,"  ",time.Now().String())
			
			}()
		}
	time.Sleep(time.Second * 2)
	fmt.Println("end Read int to channel :",time.Now().String())

	time.Sleep(time.Second * 10)
}



/**
申请写锁必须等到没有其他任何的读锁和其他的写锁才能申请成功
申请读锁必须等到没有其他写锁的时候，可以同时申请读锁
PS D:\Go\GoProject\src\goPro01\demo21\concurrence> go run .\test9.go
start write int to channel : 2021-03-14 15:54:41.471561 +0800 CST m=+0.004017401
0    2021-03-14 15:54:41.4944993 +0800 CST m=+0.026955701
1    2021-03-14 15:54:41.4954689 +0800 CST m=+0.027925301
2    2021-03-14 15:54:41.4954689 +0800 CST m=+0.027925301
3    2021-03-14 15:54:41.4954689 +0800 CST m=+0.027925301
4    2021-03-14 15:54:41.5034464 +0800 CST m=+0.035902801
5    2021-03-14 15:54:41.504445 +0800 CST m=+0.036901401
6    2021-03-14 15:54:41.504445 +0800 CST m=+0.036901401
7    2021-03-14 15:54:41.5054415 +0800 CST m=+0.037897901
8    2021-03-14 15:54:41.5054415 +0800 CST m=+0.037897901
9    2021-03-14 15:54:41.5054415 +0800 CST m=+0.037897901
data write to channel end
end write int to channel : 2021-03-14 15:54:43.5062792 +0800 CST m=+2.038735601
start Read int to channel : 2021-03-14 15:54:43.5070902 +0800 CST m=+2.039546601
get channel data : 0    2021-03-14 15:54:43.5070902 +0800 CST m=+2.039546601
get channel data : 9    2021-03-14 15:54:43.5070902 +0800 CST m=+2.039546601
get channel data : 1    2021-03-14 15:54:43.5070902 +0800 CST m=+2.039546601
get channel data : 2    2021-03-14 15:54:43.5070902 +0800 CST m=+2.039546601
get channel data : 3    2021-03-14 15:54:43.5070902 +0800 CST m=+2.039546601
get channel data : 4    2021-03-14 15:54:43.5070902 +0800 CST m=+2.039546601
get channel data : 5    2021-03-14 15:54:43.5070902 +0800 CST m=+2.039546601
get channel data : 6    2021-03-14 15:54:43.5070902 +0800 CST m=+2.039546601
get channel data : 7    2021-03-14 15:54:43.5070902 +0800 CST m=+2.039546601
get channel data : 8    2021-03-14 15:54:43.5070902 +0800 CST m=+2.039546601
end Read int to channel : 2021-03-14 15:54:45.5077443 +0800 CST m=+4.040200701
*/