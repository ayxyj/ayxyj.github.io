package main

import (
	"fmt"
	"unsafe"
	"strconv"
)
//基本数据类型总结
//整型
//int  int8 int16 int32 int 64
//uint uint8 uint16 uint32 uint 64

//浮点型
//float32 float64

//逻辑
//bool

//字符串
//string

//字符
//没有单独的字符类型  可以用byte

//派生/复杂数据类型，之后在说

func test9() {
	test8()
}
func test8()  {
	//string 转其他基本数据类型   注意函数返回值
	var str string = "true"
	var b bool
	//b,= strconv.ParseBool(str)1/说明
	// 1. strconv. ParseBool(str)函数会返回两个值（value bool, err error) 
	// 2．因为我只想获取到 value bool ,不想获取 err所以我使用_忽略
	b , _ = strconv.ParseBool(str)
	fmt.Printf(" 数据类型： %T , 数据值: %t  \n" , b , b)

	str = "1234560"
	var x , _ =  strconv.ParseInt(str , 10 , 64)
	fmt.Printf(" 数据类型：%T , 数据值：%d\n" , x , x)

	str = "ss"
	var x1 , _ =  strconv.ParseInt(str , 10 , 64)
	//在将String类型转成基本数据类型时，要确保String类型能够转成有效的数据
	//比如我们可以把"123"，转成一个整数，但是不能把"hello”转成一个整数
	//如果这样做，Golang直接将其转成0 ，其它类型也是一样的道理. float => 0 bool => false
	fmt.Printf(" 数据类型：%T , 数据值：%d\n" , x1 , x1)

	str = "123.123"
	var x2 , _ = strconv.ParseFloat(str , 64)
	fmt.Printf(" 数据类型：%T , 数据值：%15.7f \n" , x2 , x2)
	/*
	数据类型： bool , 数据值: true
	数据类型：int64 , 数据值：1234560
	数据类型：int64 , 数据值：0
	数据类型：float64 , 数据值：    123.1230000
	*/
	var aa int  = 100
	fmt.Printf("%T , %d \n" , aa , unsafe.Sizeof(aa))//int , 8

}
func test7() {
	//基本数据类型转String
	var a int = 10
	var b float32 = 100.22
	var c byte = 'a'
	var d bool = true
	var e string
	e  = fmt.Sprintf("%d" , a)
	fmt.Printf("int     -> String : \n1、str type is : %T , str data is : %s \n", e , e)
	
	e  = fmt.Sprintf("%f" , b)
	fmt.Printf("float32 -> String : \n2、str type is : %T , str data is : %s \n", e , e)

	e  = fmt.Sprintf("%c" , c)
	fmt.Printf("byte    -> String : \n3、str type is : %T , str data is : %s \n", e , e)
	
	e  = fmt.Sprintf("%t" , d)
	fmt.Printf("bool    -> String : \n4、str type is : %T , str data is : %s \n", e , e)

	fmt.Printf("\n")

	//第二种方式strconv   
	var str string 
	str = strconv.FormatInt(int64(a) , 16)
	fmt.Printf("int     -> String : \n1、str type is : %T , str data is : %s \n", str , str)

	str = strconv.FormatFloat(float64(b) , 'f' ,  16 , 64)
	//说明:‘f’格式  16:表示小数位保留16位    64 :表示这个小数是float64
	fmt.Printf("float32 -> String : \n2、str type is : %T , str data is : %s \n", str , str)

	str = strconv.FormatBool(d)
	fmt.Printf("bool    -> String : \n3、str type is : %T , str data is : %s \n", str , str)

	fmt.Printf("\n")
	/*
	int     -> String :
	1、str type is : string , str data is : 10
	float32 -> String :
	2、str type is : string , str data is : 100.220001
	byte    -> String :
	3、str type is : string , str data is : a
	bool    -> String :
	4、str type is : string , str data is : true
	int     -> String :
	1、str type is : string , str data is : a
	float32 -> String :
	2、str type is : string , str data is : 100.2200012207031250
	bool    -> String :
	3、str type is : string , str data is : true
	*/

	//strconv.Itoa()   ： int -> string 
	str = strconv.Itoa(a)
	fmt.Printf("int     -> String : \n1、str type is : %T , str data is : %s \n", str , str)

	fmt.Printf("\n")
}
func test6()  {
	/*
	//数据类型转换
	int=8 int16 int32 -> float
	float32 -> float64
	//上述两种转换过程不会存在问题
	float -> int=8 int16 int32
	float64 -> float32 
	//会出现精度丢失的现象,或者溢出

	1) Go中，数据类型的转换可以是从表示范围小-->表示范围大，也可以范围大--->范围小
	2)被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化!
	3)在转换中，比如将 int64转成int8【-128---127】，编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样。因此在转换时，需要考虑范围.

	*/
	//在Go中不存在自动拆装箱  因此需要 显示转换才可以进行数据类型的转换
	a := 185
	b := float32(a)
	var c float32 = 125e+23
	d := float64(c)
	e := int(c)
	fmt.Printf("a = %d\n" , a)
	fmt.Printf("b = %5.5f\n" , b)
	fmt.Printf("c = %v\n" , c)
	fmt.Printf("d = %v\n" , d)
	fmt.Printf("e = %d\n" , e)

	//小练习
	var n1 int8  = 12
	var n2 int32 
	var n3 int64  
	fmt.Printf("数据类型：%T , 字节数：%d" , 20 , unsafe.Sizeof(20))
	fmt.Printf("\n")
	// n2  =  n1 + 20 错误
	// n3  =  n1 + 20 错误
	n2 = int32(n1) + 20 
	n3 = int64(n1) + 20 
	fmt.Printf("%d,%d",n2,n3)

	//常见小问题
	// var t1 int32 = 12
	// var t2 int8
	// var t3 int8
	//t2 = int8(t1) + 127   编译通过   但是溢出
	//t3 = int8(t1) + 128   编译不通过 编译器识别128已经超出了int8的范围  因此不通过编译
}
func test5()  {
	//数据类型，默认值
	var a int     //0
	var b float32 //0
	var c float64 //0
	var d string  //""
	var e bool	  //false
	var f byte	  //0
	fmt.Printf("int %d , float32 %v , float64 %v , string %v , bool %v , byte %d ",a,b,c,d,e,f)
}
func test4()  {
	//字符串
	var str  ="hello Go ！\n"
	var str1 =`hello \n Go !`
	fmt.Printf(str)
	fmt.Printf("原样输出，防止攻击和源代码的原样输出显示："+str1)
	fmt.Println("\n")
	//字符串的两种定义方式： ""  ``
	//字符串拼接
	var str2 = "hello " + "Go ! "
	fmt.Printf(str2)
}
func test3()  {
	//l逻辑型
	flag := true
	fmt.Println("flag = " , flag)
	fmt.Println("\n")
	fmt.Printf("所占用字节数：%d,数据类型为：%T\n" , unsafe.Sizeof(flag) , flag)
}
func test2() {
	//浮点数
	//float32  float64
	var f1 float32 = 98.5
	var f2 float32 = -0.4002321
	var f3 float64 = -15315441
	var f4 float32 = -0.0000014451
	fmt.Printf("f1 = ", f1 , "\n", "f2 = ", f2 , "\n", "f3 = ", f3 , "\n", "f4 = ", f4, "\n")
	fmt.Println("f1 = ", f1 , "\n", "f2 = ", f2 , "\n", "f3 = ", f3 , "\n", "f4 = ", f4, "\n")
	//fmt.Printf("\n")
	// var f5 float32 = 1e+100
	//fmt.Printf("溢出：f5 = ", f5 , '\n')

	num1 := 1.2
	fmt.Printf("默认浮点型类型为%T , 字节数%d \n",  num1 , unsafe.Sizeof(num1))
}

func test1() {
	//字符
	var c byte = 255
	var n1 int = 125
	fmt.Println("c=", c)
	fmt.Printf("data type c is %T \nint byte number is %d\n", c, unsafe.Sizeof(n1))

	var c1 byte = 65
	fmt.Printf("字符65是：%c\n", c1)
	var c2 byte = 97
	fmt.Printf("字符97是：%c\n", c2)

	var cc = '呗'
	fmt.Printf("默认浮点型类型为%T,字节数%d \n",  cc , unsafe.Sizeof(cc) ) 
	fmt.Println("输出字符的int值是: \n" , cc , )

	fmt.Printf("随便输出一个字符%c\n",12345)
}

/*
import "fmt"

	mt包实现了类似C语言printf和scanf的格式化I/O。格式化动作（'verb'）源自C语言但更简单。

	Printing
	verb：

	通用：

	%v	值的默认格式表示
	%+v	类似%v，但输出结构体时会添加字段名
	%#v	值的Go语法表示
	%T	值的类型的Go语法表示
	%%	百分号
	布尔值：

	%t	单词true或false
	整数：

	%b	表示为二进制
	%c	该值对应的unicode码值
	%d	表示为十进制
	%o	表示为八进制
	%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	%x	表示为十六进制，使用a-f
	%X	表示为十六进制，使用A-F
	%U	表示为Unicode格式：U+1234，等价于"U+%04X"
	浮点数与复数的两个组分：

	%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
	%e	科学计数法，如-1234.456e+78
	%E	科学计数法，如-1234.456E+78
	%f	有小数部分但无指数部分，如123.456
	%F	等价于%f
	%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	字符串和[]byte：

	%s	直接输出字符串或者[]byte
	%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	%x	每个字节用两字符十六进制数表示（使用a-f）
	%X	每个字节用两字符十六进制数表示（使用A-F）    
	指针：

	%p	表示为十六进制，并加上前导的0x    
	没有%u。整数如果是无符号类型自然输出也是无符号的。类似的，也没有必要指定操作数的尺寸（int8，int64）。

	宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。举例如下：

	%f:    默认宽度，默认精度
	%9f    宽度9，默认精度
	%.2f   默认宽度，精度2
	%9.2f  宽度9，精度2
	%9.f   宽度9，精度0    
	宽度和精度格式化控制的是Unicode码值的数量（不同于C的printf，它的这两个因数指的是字节的数量）。两者任一个或两个都可以使用'*'号取代，此时它们的值将被对应的参数（按'*'号和verb出现的顺序，即控制其值的参数会出现在要表示的值前面）控制，这个操作数必须是int类型。

	对于大多数类型的值，宽度是输出字符数目的最小数量，如果必要会用空格填充。对于字符串，精度是输出字符数目的最大数量，如果必要会截断字符串。

	对于整数，宽度和精度都设置输出总长度。采用精度时表示右对齐并用0填充，而宽度默认表示用空格填充。

	对于浮点数，宽度设置输出总长度；精度设置小数部分长度（如果有的话），除了%g和%G，此时精度设置总的数字个数。例如，对数字123.45，格式%6.2f 输出123.45；格式%.4g输出123.5。%e和%f的默认精度是6，%g的默认精度是可以将该值区分出来需要的最小数字个数。

	对复数，宽度和精度会分别用于实部和虚部，结果用小括号包裹。因此%f用于1.2+3.4i输出(1.200000+3.400000i)。

	其它flag：

	'+'	总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
	' '	对数值，正数前加空格而负数前加负号；
	'-'	在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
	'#'	切换格式：
		八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）；
		对%q（%#q），如果strconv.CanBackquote返回真会输出反引号括起来的未转义字符串；
		对%U（%#U），输出Unicode格式后，如字符可打印，还会输出空格和单引号括起来的go字面值；
		对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格；
	'0'	使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；
	verb会忽略不支持的flag。例如，因为没有十进制切换模式，所以%#d和%d的输出是相同的。

	对每一个类似Printf的函数，都有对应的Print型函数，该函数不接受格式字符串，就效果上等价于对每一个参数都是用verb %v。
	另一个变体Println型函数会在各个操作数的输出之间加空格并在最后换行。

	不管verb如何，如果操作数是一个接口值，那么会使用接口内部保管的值，而不是接口，因此：

	var i interface{} = 23
	fmt.Printf("%v\n", i)
	会输出23。

	除了verb %T和%p之外；对实现了特定接口的操作数会考虑采用特殊的格式化技巧。按应用优先级如下：

	1. 如果操作数实现了Formatter接口，会调用该接口的方法。Formatter提供了格式化的控制。

	2. 如果verb %v配合flag #使用（%#v），且操作数实现了GoStringer接口，会调用该接口。

	如果操作数满足如下两条任一条，对于%s、%q、%v、%x、%X五个verb，将考虑：

	3. 如果操作数实现了error接口，Error方法会用来生成字符串，随后将按给出的flag（如果有）和verb格式化。

	4. 如果操作数具有String方法，这个方法将被用来生成字符串，然后将按给出的flag（如果有）和verb格式化。

	复合类型的操作数，如切片和结构体，格式化动作verb递归地应用于其每一个成员，而不是作为整体一个操作数使用。
	因此%q会将[]string的每一个成员括起来，%6.2f会控制浮点数组的每一个元素的格式化。

	为了避免可能出现的无穷递归，如：

	type X string
	func (x X) String() string { return Sprintf("<%s>", x) }
	应在递归之前转换值的类型：

	func (x X) String() string { return Sprintf("<%s>", string(x)) }
	显式指定参数索引：

	在Printf、Sprintf、Fprintf三个函数中，默认的行为是对每一个格式化verb依次对应调用时成功传递进来的参数。
	但是，紧跟在verb之前的[n]符号表示应格式化第n个参数（索引从1开始）。同样的在'*'之前的[n]符号表示采用第n个参数的值作为宽度或精度。
	在处理完方括号表达式[n]后，除非另有指示，会接着处理参数n+1，n+2……（就是说移动了当前处理位置）。例如：

	fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	会生成"22 11"，而：

	fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6),
	等价于：

	fmt.Sprintf("%6.2f", 12.0),
	会生成" 12.00"。因为显式的索引会影响随后的verb，这种符号可以通过重设索引用于多次打印同一个值：

	fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
	会生成"16 17 0x10 0x11"

	格式化错误：

	如果给某个verb提供了非法的参数，如给%d提供了一个字符串，生成的字符串会包含该问题的描述，如下所例：

	错误的类型或未知的verb：%!verb(type=value)
		Printf("%d", hi):          %!d(string=hi)
	太多参数（采用索引时会失效）：%!(EXTRA type=value)
		Printf("hi", "guys"):      hi%!(EXTRA string=guys)
	太少参数: %!verb(MISSING)
		Printf("hi%d"):            hi %!d(MISSING)
	宽度/精度不是整数值：%!(BADWIDTH) or %!(BADPREC)
		Printf("%*s", 4.5, "hi"):  %!(BADWIDTH)hi
		Printf("%.*s", 4.5, "hi"): %!(BADPREC)hi
	没有索引指向的参数：%!(BADINDEX)
		Printf("%*[2]d", 7):       %!d(BADINDEX)
		Printf("%.[2]d", 7):       %!d(BADINDEX)
	所有的错误都以字符串"%!"开始，有时会后跟单个字符（verb标识符），并以加小括弧的描述结束。

	如果被print系列函数调用时，Error或String方法触发了panic，fmt包会根据panic重建错误信息，用一个字符串说明该panic经过了fmt包。
	例如，一个String方法调用了panic("bad")，生成的格式化信息差不多是这样的：

	%!s(PANIC=bad)
	%!s指示表示错误（panic）出现时的使用的verb。

	Scanning
	一系列类似的函数可以扫描格式化文本以生成值。

	Scan、Scanf和Scanln从标准输入os.Stdin读取文本；Fscan、Fscanf、Fscanln从指定的io.Reader接口读取文本；Sscan、Sscanf、Sscanln从一个参数字符串读取文本。

	Scanln、Fscanln、Sscanln会在读取到换行时停止，并要求一次提供一行所有条目；Scanf、Fscanf、Sscanf只有在格式化文本末端有换行时会读取到换行为止；其他函数会将换行视为空白。

	Scanf、Fscanf、Sscanf会根据格式字符串解析参数，类似Printf。例如%x会读取一个十六进制的整数，%v会按对应值的默认格式读取。格式规则类似Printf，有如下区别：

	%p 未实现
	%T 未实现
	%e %E %f %F %g %G 效果相同，用于读取浮点数或复数类型
	%s %v 用在字符串时会读取空白分隔的一个片段
	flag '#'和'+' 未实现   
	在无格式化verb或verb %v下扫描整数时会接受常用的进制设置前缀0（八进制）和0x（十六进制）。

	宽度会在输入文本中被使用（%5s表示最多读取5个rune来生成一个字符串），但没有使用精度的语法（没有%5.2f，只有%5f）。

	当使用格式字符串进行扫描时，多个连续的空白字符（除了换行符）在输出和输出中都被等价于一个空白符。在此前提下，格式字符串中的文本必须匹配输入的文本；如果不匹配扫描会中止，函数的整数返回值说明已经扫描并填写的参数个数。

	在所有的扫描函数里，\r\n都被视为\n。

	在所有的扫描函数里，如果一个操作数实现了Scan方法（或者说，它实现了Scanner接口），将会使用该接口为该操作数扫描文本。
	另外，如果如果扫描到（准备填写）的参数比提供的参数个数少，会返回一个错误。

	提供的所有参数必须为指针或者实现了Scanner接口。注意：Fscan等函数可能会在返回前多读取一个rune，这导致多次调用这些函数时可能会跳过部分输入。只有在输入里各值之间没有空白时，会出现问题。
	如果提供给Fscan等函数的io.Reader接口实现了ReadRune方法，将使用该方法读取字符。如果该io.Reader接口还实现了UnreadRune方法，将是使用该方法保存字符，这样可以使成功执行的Fscan等函数不会丢失数据。
	如果要给一个没有这两个方法的io.Reader接口提供这两个方法，使用bufio.NewReader。
	*/