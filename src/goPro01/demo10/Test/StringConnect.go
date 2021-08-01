package main

import(
	"fmt"
	"bytes"
)

func main()  {
	hammer := "吃我一锤"
	sickle := "死吧"
	// 声明字节缓冲
	var stringBuilder bytes.Buffer
	// 把字符串写入缓冲
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	// 将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())

	//bytes.Buffer 是可以缓冲并可以往里面写入各种字节数组的。字符串也是一种字节数组，使用 WriteString() 方法进行写入。
	// 将需要连接的字符串，通过调用 WriteString() 方法，写入 stringBuilder 中，然后再通过 stringBuilder.String() 方法将缓冲转换为字符串。	
}
