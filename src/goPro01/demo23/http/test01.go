package main1

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	//控制器 ， 访问路径
	//传入处理函数sayhello
	http.HandleFunc("/", sayHello)
	//设置监听端口号
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

//handler 处理请求函数    serve.HandlerMux
func sayHello(resp http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println(req.URL)
	fmt.Println(req.URL.Path)
	fmt.Println(req.RemoteAddr)
	fmt.Println(req.Host)
	/**
	map[username:[zzu]]
	/?username=zzu
	/
	127.0.0.1:49593
	127.0.0.1:8899
	*/
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, "#"))
		/**
		key: username
		val: zzu
		map[]
		/favicon.ico
		/favicon.ico
		127.0.0.1:49593
		127.0.0.1:8899
		*/
	}
	//resp写入返回客户端的内容
	fmt.Fprintf(resp, "hello web %s!\n", req.Form["username"])
}
