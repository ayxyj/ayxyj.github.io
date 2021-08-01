package main

import (
	"fmt"
	"net/http"
	"time"
)


//管理服务端的行为

type MyHandlers struct {

}

func (mh *MyHandlers)ServeHTTP(resp http.ResponseWriter , req http.Request)  {
	fmt.Fprintf(resp , "hello success!")
}


func main()  {
	var handler MyHandlers
	var server = http.Server{
			Addr:":80",
			Handler:&handler
			ReadTimeOut: 2 *  time.Second
			MaxHeaderBytes: 1<<20
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("http server failed:%v\n" , err)
		return
	}
}

func MyHandler(resp http.ResponseWriter , req *http.Request) {
	defer req.Body.Close()
	fmt.Println("method:" , req.Method)
	fmt.Println("Url:" , req.URL)
	fmt.Println("header:" , req.Header)
	fmt.Println("body:" , req.Body)
	fmt.Println("remoteAddr"  , req.RemoteAddr)
	resp.Write([]byte("hello success!"))
}
