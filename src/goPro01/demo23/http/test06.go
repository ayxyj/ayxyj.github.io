package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//客户端主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err) //log.Fatal()会产生panic
		return
	}

	defer conn.Close() //关闭
	//http://127.0.0.1:8080/block/sqlite
	requestHeader := "GET /block/sqlite  HTTP/1.1\r\n" +
		"Accept: image/gif, image/jpeg, image/pjpeg, application/x-ms-application, application/xaml+xml, application/x-ms-xbap, */*\r\n" +
		"Accept-Language: zh-Hans-CN,zh-Hans;q=0.8,en-US;q=0.5,en;q=0.3\r\n" +
		"User-Agent: Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)\r\n" +
		"Accept-Encoding: gzip, deflate" +
		"\r\nHost: 127.0.0.1:8080\r\n" +
		"Connection: Keep-Alive\r\n\r\n"

	//先发送请求包
	conn.Write([]byte(requestHeader))

	buf := make([]byte, 4096) //缓冲区

	//阻塞等待服务器回复的数据
	n, err := conn.Read(buf) //n代码接收数据的长度
	if err != nil {
		fmt.Println(err)
		return
	}
	//切片截取，只截取有效数据
	result := buf[:n]
	fmt.Printf("接收到数据[%d]:\n%s\n", n, string(result))
}

