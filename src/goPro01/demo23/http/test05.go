package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	"sourcegraph.com/sourcegraph/go-selenium"
)

func main() {
	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "firefox"})
	if webDriver, err = selenium.NewRemote(caps, "http://localhost:8888"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}
	defer webDriver.Quit()

	err = webDriver.Get("http://www.yiyou.org")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	robotgo.MoveMouseSmooth(323, 186) //移动鼠标
	robotgo.MouseClick("left", true) //单击
	time.Sleep(time.Second * 5)
	robotgo.MoveMouseSmooth(422, 412)
	robotgo.MouseClick("left", true)
	time.Sleep(time.Second * 5)

	cookies, err := webDriver.GetCookies() //获取cookie
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 5)
	fmt.Print("Cookie:", cookies)
	fmt.Print(robotgo.GetMousePos())
}
