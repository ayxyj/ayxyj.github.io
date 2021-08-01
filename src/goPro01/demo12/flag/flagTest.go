package main
// 导入系统包
import (
    "flag"
	"fmt"
	"os"
	"math"
)
// 定义命令行参数
var mode = flag.String("mode", "", "process mode")
func main() {
	func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	}()
    // 解析命令行参数
    flag.Parse()
    // 输出命令行参数
	fmt.Println(*mode)
	//go run .\flagTest.go  --mode=test
	//test 
	test()
}

func test()  {
	//（square root）
	getSqrtVal := func(x float64) float64{
		return math.Sqrt(x)
	}(64)
	fmt.Println(getSqrtVal)
}