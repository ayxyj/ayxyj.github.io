package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main()  {
	test04()

}

func test09()  {
	type IT struct {
		Company  string   `json:"company"`
		Subjects []string `json:"subjects"`
		IsOk     bool     `json:"isok"`
		Price    float64  `json:"price"`
	}

		b := []byte(`{
						"company": "itcast",
						"subjects": [
							"Go",
							"C++",
							"Python",
							"Test"
						],
						"isok": true,
						"price": 666.666
					}`)

		var t IT
		err := json.Unmarshal(b, &t)
		if err != nil {
			fmt.Println("json err:", err)
		}
		fmt.Println(t)
		//运行结果：{itcast [Go C++ Python Test] true 666.666}

		//只想要Subjects字段
		type IT2 struct {
			Subjects []string `json:"subjects"`
		}

		var t2 IT2
		err = json.Unmarshal(b, &t2)
		if err != nil {
			fmt.Println("json err:", err)
		}
		fmt.Println(t2)
		//运行结果：{[Go C++ Python Test]}
}

func test08()  {
	type IT struct {
		//Company不会导出到JSON中
		Company string `json:"-"`

		// Subjects 的值会进行二次JSON编码
		Subjects []string `json:"subjects"`

		//转换为字符串，再输出
		IsOk bool `json:",string"`
		IsOk2 bool `json:"isOk2"`

		// 如果 Price 为空，则不输出到JSON串中
		Price float64 `json:"price, omitempty"`
	}
	t1 := IT{Company: "itcast", Subjects: []string{"Go", "C++", "Python", "Test"}, IsOk: true, IsOk2:false}

	b, err := json.Marshal(t1)
	//json.MarshalIndent(t1, "", "    ")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
	//{"subjects":["Go","C++","Python","Test"],"IsOk":"true","isOk2":false,"price":0}
}

func test04()  {

	type Author struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	type Comment struct {
		ID      int64  `json:"id"`
		Content string `json:"content"`
		Author  string `json:"author"`
	}
	type Post struct {
		ID        int64         `json:"id"`
		Content   string        `json:"content"`
		Author    Author        `json:"author"`
		Published bool          `json:"published"`
		Label     []string      `json:"label"`
		NextPost  *Post         `json:"nextPost"`
		Comments  []*Comment    `json:"comments"`
	}


	var files []string

	root := "D:/Go/GoProject/src/goPro01/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	//for _, file := range files {
	//	fmt.Println(file)
	//}


	fmt.Println("pwd:")
	str, _ := os.Getwd()
	fmt.Println(str)


	// 打开json文件
	fmt.Println()
	//fh, err := os.Open("D:/Go/GoProject/src/goPro01/demo23/json/a.json")
	getwd, err := os.Getwd()
	//fh, err := os.Open(getwd + "/src/goPro01/demo23/json/a.json")
	fh, err := os.Open(getwd + "/a.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()
	// 读取json文件，保存到jsonData中
	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T", jsonData)
	var post Post
	// 解析json数据到post中
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post)
}

func test03()  {

	//map
	mapss := map[string]string{
		"python":"nb",
		"Java":"nb1",
		"Go":"nb2",
	}
	fmt.Println("Mapss print: ")
	fmt.Println(mapss)
	fmt.Println("mapss json print: ")
	mapssJson, _ := json.Marshal(mapss)
	mapssJsonIndent , _ := json.MarshalIndent(mapss ,"" , "\t")
	fmt.Println(string(mapssJson))
	fmt.Println(string(mapssJsonIndent))

	//上面是已知json数据结构的解析方式，如果json结构是未知的或者结构可能会发生改变的情况，则解析到struct是不合理的。
	//这时可以解析到空接口interface{}或map[string]interface{}类型上，这两种类型的结果是完全一致的。
	fmt.Println("unMarshal json :")
	//var it map[string]interface{}
	//var it interface{}
	var it map[string]string
	json.Unmarshal(mapssJsonIndent, &it)
	fmt.Println(it)
}
func test02()  {
	/**
	struct能被转换的字段都是首字母大写的字段，但如果想要在json中使用小写字母开头的key，可以使用struct的tag来辅助反射。
	构建json数据
	    Marshal()和MarshalIndent()函数可以将数据封装成json数据。
	        struct、slice、array、map都可以转换成json
			struct转换成json的时候，只有字段首字母大写的才会被转换
			map转换的时候，key必须为string
			封装的时候，如果是指针，会追踪指针指向的对象进行封装
	*/
	//struct
	it1 := &IT{
		Company: "zzu",
		Class:   []string{"python", "go" , "Java" , "c"},      //slice
		Flag:    true,
		Price:   666.66,
	}

	maps := make(map[int]string)
	maps[1000]="hello Java"

	fmt.Printf("%s\n" , it1.Company)

	indent, err := json.MarshalIndent(it1 , "" , "    ")
	marshal, err := json.Marshal(it1)

	if err!=nil {
		fmt.Printf("json foramt  marshalIndent error :%s\n", err)
	}
	fmt.Println("marshalIndent Json print:",string(indent))
	fmt.Println("marshal Json print:",string(marshal))
}

func test01()  {

	/**
	使用struct tag的时候，几个注意点：
		tag中标识的名称将称为json数据中key的值
		tag可以设置为`json:"-"`来表示本字段不转换为json数据，即使这个字段名首字母大写
		如果想要json key的名称为字符"-"，则可以特殊处理`json:"-,"`，也就是加上一个逗号
		如果tag中带有,omitempty选项，那么如果这个字段的值为0值，即false、0、""、nil等，这个字段将不会转换到json中
		如果字段的类型为bool、string、int类、float类，而tag中又带有,string选项，那么这个字段的值将转换成json字符串 */
	type Person struct {
		Id int  		`json:"id,string"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	lisi := Person{
		Id:       1000,
		Username: "lisi",
		Password: "123456",
	}

	indent, err := json.MarshalIndent(lisi, "", "\t")
	if err!=nil {
		fmt.Println("err json marshal :" , err)
	}

	fmt.Println(string(indent))

}
type IT struct {
	Company string
	Class []string
	Flag bool
	Price float64
}



///--------------------------------------------------找的demo


func test07()  {
	//再例如，从标准输入读json数据，解码后删除名为Name的元素，最后重新编码后输出到标准输出。
		dec := json.NewDecoder(os.Stdin)
		enc := json.NewEncoder(os.Stdout)
		for {
			var v map[string]interface{}
			if err := dec.Decode(&v); err != nil {
				log.Println(err)
				return
			}
			for k := range v {
				if k != "Name" {
					delete(v, k)
				}
			}
			if err := enc.Encode(&v); err != nil {
				log.Println(err)
			}
		}
}

func test06()  {

	/**
	解析、创建json流
	除了可以直接解析、创建json数据，还可以处理流式数据。
		type Decoder解码json到Go数据结构
		type Encoder编码Go数据结构到json
	*/
	const jsonStream = `
    {"Name": "Ed", "Text": "Knock knock."}
    {"Name": "Sam", "Text": "Who's there?"}
    {"Name": "Ed", "Text": "Go fmt."}
    {"Name": "Sam", "Text": "Go fmt who?"}
    {"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

func test05()  {
	// 读取json数据
	fh, err := os.Open("D:/Go/GoProject/src/goPro01/demo23/json/a.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()
	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析json数据到interface{}
	var unknown interface{}
	err = json.Unmarshal(jsonData, &unknown)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 进行断言，并switch匹配
	m := unknown.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "type: string\nvalue: ", vv)
			fmt.Println("------------------")
		case float64:
			fmt.Println(k, "type: float64\nvalue: ", vv)
			fmt.Println("------------------")
		case bool:
			fmt.Println(k, "type: bool\nvalue: ", vv)
			fmt.Println("------------------")
		case map[string]interface{}:
			fmt.Println(k, "type: map[string]interface{}\nvalue: ", vv)
			for i, j := range vv {
				fmt.Println(i,": ",j)
			}
			fmt.Println("------------------")
		case []interface{}:
			fmt.Println(k, "type: []interface{}\nvalue: ", vv)
			for key, value := range vv {
				fmt.Println(key, ": ", value)
			}
			fmt.Println("------------------")
		default:
			fmt.Println(k, "type: nil\nvalue: ", vv)
			fmt.Println("------------------")
		}
	}
}

