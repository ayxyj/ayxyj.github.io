package main

import(
	"bytes"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
)

//请求结构体
type requestBody struct {
	Key string `json:"key"`
	Info string `json:"info"`
	UserId string `json:"userid"`
}

//结果结构体
type responseBody struct {
	Code int `json:"code"`
	Text string `json:"text"`
	List []string `json:"list"`
	Url  string `json:"url"`
}

//请求机器人
func process( inputChan <- chan string , userid string ){
	for{
		//从通道中接受输入
		input := <- inputChan
		if input == "EOF" {
			break
		}
		//构建请求体
		reqData := &requestBody {
			Key : "792bcfdwawfafwadw4a85d16w1a561d65w1a65d165w",
			Info : input,
			UserId : userid,
		}
		//转义Json
		byteData , _ := json.Marshal( &reqData )
		//请求聊天机器人接口
		req , err := http.NewRequest(
			"POST",
			"https://zzugo.ayxyj.cn/",
			bytes.NewReader( byteData ))
		req.Header.Set( "Content-type" , "application/json;charset=UTF-8" )
		client := http.Client{}
		resp , err := client.Do( req )
		if err != nil {
			fmt.Println( "NewWork Error!" )
		}else{
			//将结果从Json中解析并输出到命令行
			body , _ := ioutil.ReadAll( resp.Body )
			var respData responseBody
			json.Unmarshal( body , &respData )
			fmt.Println("AI:"+respData.Text)
		}
		resp.Body.Close()
	}
}
func main(){
		var input string 
		fmt.Println( "输入`EOF`结束程序！" )
		//创建通道
		channel := make( chan string )
		//main结束时关闭通道
		defer  close( channel )
		//启动goroutine运行机器人回答线程
		go process( channel , string(rand.Int63()) )
		for{
			//从命令行中读取输入
			fmt.Scanf( "%s" , &input )
			//将输入放入通道
			channel <- input 
			//结束程序
			if input == "EOF"{
				fmt.Println("Bye!")
				break
			}
		}
}