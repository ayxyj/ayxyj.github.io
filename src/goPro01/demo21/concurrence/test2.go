package main

import(
	"fmt"
	"time"
)

func main() {
	v := new(int)
	go setVTo1(v)
	go setVTo2(v)

	time.Sleep(time.Second)
	fmt.Println(*v)
}

func setVTo1(v *int)  {
	*v = 1
	//time.Sleep(time.Second)
}

func setVTo2(v *int){
	*v = 2
	//time.Sleep(time.Second)
}