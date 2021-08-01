package main

import (
    "encoding/base64"
	"fmt"
	"goPro01/demo10/string04/utils"
)
func main()  {
	fmt.Println(utils.EncodeBase64("zxc"))
	fmt.Println(utils.EncodeBase64("严"))
	fmt.Printf("%x\n", "严")
	fmt.Printf("%b\n", 0xe4b8a5)

	fmt.Println(BToD2("00100011"))

}
func test() {
	// 需要处理的字符串
	/**
	message := "zxc" // 00011110 00100111 00100001 00100011  -> // enhj
	fmt.Printf("%b\n" , 'z') //01111010
	fmt.Printf("%b\n" , 'x') //01111000
	fmt.Printf("%b\n" , 'c') //01100011
	fmt.Println(BToD("00011110")) // 30 e
	fmt.Println(BToD("00100111")) // 39 n
	fmt.Println(BToD("00100001")) //33 h
	fmt.Println(BToD("00100011")) //35 j
	*/
	message := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
    // 编码消息
    encodedMessage := base64.StdEncoding.EncodeToString([]byte (message))
    // 输出编码完成的消息
    fmt.Println(encodedMessage)
    // 解码消息
    data, err := base64.StdEncoding.DecodeString(encodedMessage)
    // 出错处理
    if err != nil {
        fmt.Println(err)
    } else {
        // 打印解码完成的数据
        fmt.Println(string(data))
    }
}

//进制转换 2 -> 10
func BToD( s string ) ( num int ) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) - 48) << uint8(i)  
	}
	return num
}
func BToD2(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0x1) << uint8(i)
	}
	return
}
/**
Base64编码：
所谓Base64，就是说选出64个字符----小写字母a-z、大写字母A-Z、数字0-9、符号"+"、"/"（再加上作为垫字的"="，实际上是65个字符）
作为一个基本字符集。然后，其他所有符号都转换成这个字符集中的字符。

第一步，将每三个字节作为一组，一共是24个二进制位。
第二步，将这24个二进制位分为四组，每个组有6个二进制位。
第三步，在每组前面加两个00，扩展成32个二进制位，即四个字节。
第四步，根据表，得到扩展后的每个字节的对应符号，这就是Base64的编码值。

因为，Base64将三个字节转化成四个字节，因此Base64编码后的文本，会比原文本大出三分之一左右。
*/