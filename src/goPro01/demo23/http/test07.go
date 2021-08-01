package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


/**
	// encode
	hello := "hello world"
	debyte := base64Encode([]byte(hello))

	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(enbyte))
}
*/


func main() {

	//get方式请求一个资源
	// resp, err := http.Get("http://127.0.0.1:8080/block/sqlite")
	//resp ,err :=http.Get("http://zzugo.ayxyj.cn/")
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close() //关闭

	fmt.Println("header = ", resp.Header)
	fmt.Printf("resp status %s\nstatusCode %d\n", resp.Status, resp.StatusCode)
	fmt.Printf("body type = %T\n", resp.Body)

	buf := make([]byte, 2048) //切片缓冲区
	var buf1 []byte //切片缓冲区
	var tmp string

	type Block struct{
		Id int 				`json:"id"`
		CreateTime int 		`json:"createTime"`
		UpdateTime int 		`json:"updateTime"`
		PublicKey string	`json:"publicKey"`
		Content string 		`json:"content"`
		Target string		`json:"target"`
		Origin string		`json:"origin"`
		MessageId string	`json:"messageId"`
	}

	type Body struct{
		Code int 			`json:"code"`
		Message string 		`json:"message"`
		Data []*Block		`json:"data"`
	}
	for {
		n, err := resp.Body.Read(buf) //读取body包内容
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}

		if n == 0 {
			//fmt.Println("读取内容结束")
			break
		}
		tmp += string(buf[:n]) //累加读取的内容
		buf1 = buf[:n]
	}

	fmt.Println(string(buf[:]))
	fmt.Println("=========================")
	fmt.Println(string(buf1[:]))
	//fmt.Println("buf = ", string(tmp))
	var body Body
	// 解析json数据到post中
	err = json.Unmarshal(buf1, &body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(body)
	block0 := body.Data[0]
	block1 := body.Data[1]

	fmt.Println()

	// 进行加密
	e := base64.StdEncoding.EncodeToString([]byte("180,100,175,50"))
	fmt.Printf("获取区块%d 的 180,100,175,50 加密后的结果： %s    \n", block0.Id , e)
	// 进行解密
	str, err := base64.StdEncoding.DecodeString(block0.Content)
	if err != nil{
		fmt.Println("Error: ", err)
	}
	fmt.Printf("%s       解密后的结果： %s   \n" , e , string(str))
	fmt.Println()
	fmt.Println()

	fmt.Printf("获取到区块 %d 的 178 的 sha256 加密信息： \n" , block1.Id)
	fmt.Println(block1.Content)
	sEnc := sha256.Sum256([]byte("178"))
	fmt.Println("178 进行 sha256 加密后结果与区块中的加密信息一致:")
	fmt.Printf("%x\n", sEnc)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
}



