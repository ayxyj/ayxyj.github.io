package main 

import(
	"fmt"
	"bufio"
	"os"
)

func main()  {
	//创建一个缓冲通道
	ch := make(chan string)
	//读取
	go printInput(ch)
	
	//从命令行读取输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		val := scanner.Text()
		ch <- val
		if val=="EOF"{
			fmt.Println("End the game!")
			break
		}
	}
	defer close(ch)
}
//读通道数据
func printInput(ch chan string)  {
	for val := range ch {
		if val=="EOF"{
			break
		}
		fmt.Printf("Input is %s\n",val)
	}
}