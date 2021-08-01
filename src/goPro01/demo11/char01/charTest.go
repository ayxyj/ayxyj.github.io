package main

import(
	"fmt"
	"unicode"
)

func main()  {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U\n", ch, ch2, ch3)   // UTF-8 code point

	fmt.Println("\n----function----")
	//unicode function judge char type 
	fmt.Println(unicode.IsLetter('a'))
	fmt.Println(unicode.IsDigit('1'))
	fmt.Println(unicode.IsSpace(' '))

	//字符是按照unicode码点存储并在内存中使用 int 来表示，字符串是按照utf8进行编码
	fmt.Printf("%x\n",'严') //4e25  (unicode)
	fmt.Printf("%x\n","严") //e4b8a5(utf8)


}


/**
简单来说:
	Unicode是「字符集」	
	UTF-8是「编码规则」
其中:
	字符集:为每一个「字符」分配一个唯一的ID(学名为码位/码点/Code Point)
	编码规则:将「码位」转换为字节序列的规则（(编码/解码可以理解为加密/解密的过程)
*/

/**
(十六进制)        |              （二进制）
----------------------+---------------------------------------------
0000 0000-0000 007F | 0xxxxxxx
0000 0080-0000 07FF | 110xxxxx 10xxxxxx
0000 0800-0000 FFFF | 1110xxxx 10xxxxxx 10xxxxxx
0001 0000-0010 FFFF | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx

如果一个字节的第一位是0，则这个字节单独就是一个字符；如果第一位是1，则连续有多少个1，就表示当前字符占用多少个字节。
下面，还是以汉字严为例，演示如何实现 UTF-8 编码。
严的 Unicode 是4E25（100111000100101），根据上表，可以发现4E25处在第三行的范围内（0000 0800 - 0000 FFFF），
因此严的 UTF-8 编码需要三个字节，即格式是1110xxxx 10xxxxxx 10xxxxxx。然后，从严的最后一个二进制位开始，
依次从后向前填入格式中的x，多出的位补0。这样就得到了，严的 UTF-8 编码是11100100 10111000 10100101，转换成十六进制就是E4B8A5。
*/

/**
Unicode 包中内置了一些用于测试字符的函数，这些函数的返回值都是一个布尔值，如下所示（其中 ch 代表字符）：
判断是否为字母：unicode.IsLetter(ch)
判断是否为数字：unicode.IsDigit(ch)
判断是否为空白符号：unicode.IsSpace(ch)
*/

/**
Go语言的字符有以下两种：
一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。

byte 类型是 uint8 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，完全没有问题，例如 var ch byte = 'A'，字符使用单引号括起来。
在 ASCII 码表中，A 的值是 65，使用 16 进制表示则为 41，所以下面的写法是等效的：
var ch byte = 65 或 var ch byte = '\x41'      //（\x 总是紧跟着长度为 2 的 16 进制数）
另外一种可能的写法是\后面紧跟着长度为 3 的八进制数，例如 \377。

Go语言同样支持 Unicode（UTF-8），因此字符同样称为 Unicode 代码点或者 runes，并在内存中使用 int 来表示。
在文档中，一般使用格式 U+hhhh 来表示，其中 h 表示一个 16 进制数。
在书写 Unicode 字符时，需要在 16 进制数之前加上前缀\u或者\U。
因为 Unicode 至少占用 2 个字节，所以我们使用 int16 或者 int 类型来表示。
如果需要使用到 4 字节，则使用\u前缀，如果需要使用到 8 个字节，则使用\U前缀。
*/