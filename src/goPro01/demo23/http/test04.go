package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

func main() {
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/block/sqlite/content", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Body)


	//// encode
	//hello := resp.Body
	//debyte := base64Encode([]byte(hello))
	//fmt.Printf("encode body :",debyte)
	//
	//// decode
	//enbyte, err := base64Decode(debyte)
	//fmt.Printf("decode body:",enbyte)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//if hello != string(enbyte) {
	//	fmt.Println("hello is not equal to enbyte")
	//}
	//
	//fmt.Println(string(enbyte))



	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%s\n", string(body))
}