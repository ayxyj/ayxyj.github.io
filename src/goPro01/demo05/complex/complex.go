
package main

import( 
	"fmt"
)

func main()  {
	var x complex128 = complex(1,2)
	var y complex128 = complex(3,4)
	fmt.Println(x*y)
	fmt.Println(real(x*y))
	fmt.Println(imag(x*y))
	var str = "hello , 北京"
	str1 := []rune(str)
	for index , value := range str1{
		fmt.Printf("index = %d , value = %c\n",index,value)
	}

	fmt.Printf("prefix1=%t\n", hasPrefix("hello , world!" , "hello"))
	fmt.Printf("prefix2=%t\n", hasPrefix("hello , world!" , "hello1"))
	fmt.Printf("suffix1=%t\n", hasPrefix("hello , world!" , "world!"))
	fmt.Printf("suffix2=%t\n", hasPrefix("hello , world!" , "world1"))
}

func hasPrefix ( s string , prefix string )  bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix 
}
func hasSuffix ( s string , suffix string )  bool  {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix 
}