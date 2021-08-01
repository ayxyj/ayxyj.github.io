package main 

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/" , MyHandler1)
	err := http.ListenAndServe(":8080",nil)
	if err == nil {
		fmt.Printf("listenAndServe faile!\n")
	}
}

func MyHandler1(resp http.ResponseWriter , req *http.Request)  {
	req.ParseForm()//先解析，不然表单无法处理
	fmt.Fprintf(resp , "hello test %s ",req.Form["test"])
}