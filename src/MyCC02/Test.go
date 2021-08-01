package main
import (
	"github.com/labstack/echo"
	"fmt"
	"net/http"
	"github.com/jinzhu/configor"
)
func main() {
	fmt.Println("test",configor.Config{})
}
func test()  {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World1!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}