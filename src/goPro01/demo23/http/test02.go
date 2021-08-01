package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", MyHandler2)
	err := http.ListenAndServe("127.0.0.1:80", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}

func MyHandler2(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	fmt.Println("method:", req.Method)
	fmt.Println("Url:", req.URL)
	fmt.Println("header:", req.Header)
	fmt.Println("body:", req.Body)
	fmt.Println("remoteAddr", req.RemoteAddr)
	resp.Write([]byte("hello success!"))
}

//method: GET

//Url: /

//header: map[
//Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9]
//Accept-Encoding:[gzip, deflate, br]
//Accept-Language:[en,zh-CN;q=0.9,zh;q=0.8]
//Connection:[keep-alive]
//Sec-Ch-Ua:["Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"]
//Sec-Ch-Ua-Mobile:[?0] Sec-Fetch-Dest:[document]
//Sec-Fetch-Mode:[navigate]
//Sec-Fetch-Site:[none]
//Sec-Fetch-User:[?1]
//Upgrade-Insecure-Requests:[1]
//User-Agent:[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36]]

//body: {}

//remoteAddr 127.0.0.1:50767
