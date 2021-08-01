package main

import(
	"fmt"
	"time"
	"bufio"
	"os"
)

func main()  {
	//创建通道  通道缓存大小为10
	ch := make(chan string , 10)
	//读通道
	go printChannel(ch)
	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		val := scanner.Text()
		ch <- val
		if val=="EOF"{
			fmt.Println("game over!")
			break
		}
	}
	defer close(ch)

}

func printChannel(ch chan string)  {
	time.Sleep(time.Second * 10)
	for val := range ch {
		if val == "EOF"{
			break
		}
		fmt.Printf("input %s\n",val)
	}
}