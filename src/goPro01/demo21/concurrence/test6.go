package main

import (
	"fmt"
	"bufio"
	"time"
	"os"
)

func main()  {
	//创建通道
	var ch = make(chan string , 10)
	//单输入通道
	//var chInput chan <- string = ch 
	//单输出通道
	//var chOutput string <- chan = ch 
	//输入数据
	go send(ch)
	go pull(ch)
	time.Sleep(time.Second * 100)
}
func send(ch chan string)  {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		if scanner.Text() == "EOF"{
			fmt.Println("Game!~")
			break
		}
		ch <- scanner.Text()
	}
}
func pull(ch chan string)  {
	time.Sleep(time.Second * 10)
	for val := range ch {
		if val == "EOF"{
			break
		}
		fmt.Printf("  %s  ",val)
	}
}