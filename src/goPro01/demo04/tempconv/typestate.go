package main

import (
	"fmt"
)
//类型声明
type Cels float64 //摄氏温度 Celsius
type Fahr float64 //华氏温度 Fahrenheit
const(
	Abso Cels = -273.15
	Free Cels = 0
	Boil Cels = 100
)

func CToF(c Cels)  Fahr{
	return Fahr( c * 9 / 5 + 31 )
}
func FToC(f Fahr)  Cels{
	return Cels( (f-32) * 5 / 9 )
}

func main()  {
	var t Fahr = 100
	fmt.Printf("%v\n" , FToC(t))
	fmt.Printf("%v\n" , Abso)
	fmt.Printf("%v\n" , Free)
	fmt.Printf("%v\n" , Boil)
}

