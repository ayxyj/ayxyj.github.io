package main

import(
	"fmt"
	"time"
)

func main()  {
	go func( name string ){
		fmt.Println("name :",name)
	}("zzu")
	time.Sleep(time.Second)
}
