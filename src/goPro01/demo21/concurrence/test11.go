package main

import(
	"fmt"
	"sync"
	//"strconv"
)


/**
type Map
    //删除指定key
    func (m *Map) Delete(key interface{})
    
    //查询指定key
    func (m *Map) Load(key interface{}) (value interface{}, ok bool)
 
    //查询，查不到则追加
    func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
 
    //遍历map
    func (m *Map) Range(f func(key, value interface{}) bool)
 
    //添加
    func (m *Map) Store(key, value interface{})
*/
type Class struct{
	//并发安全字典
	Student sync.Map
}
func handler(key , val interface{} ) bool{
	fmt.Printf(" key: %T val：%T  key: %s val：%s\n",key,val,key,val)
	return true
}
func main()  {
	class := &Class{}

	//存储值
	class.Student.Store("lisi","class 1")
	class.Student.Store("zhangsan","class 2")
	class.Student.Store("wangwu","class 3")

	//遍历，传入一个处理函数，遍历的时候函数返回false则停止遍历
	class.Student.Range(handler)

	//查询或者追加
	_,loaded :=class.Student.LoadOrStore("zzu","class4")
	if loaded{
		fmt.Println("success load!")
	}else{
		fmt.Println("success Store!")
	}
	//遍历
	class.Student.Range(handler)
	//删除
	class.Student.Delete("zzu")
	fmt.Println("success delete!")
	//遍历
	class.Student.Range(handler)
}

/**
PS D:\Go\GoProject\src\goPro01\demo21\concurrence> go run .\test11.go
key: lisi val：class 1
key: zhangsan val：class 2
key: wangwu val：class 3
success Store!
key: wangwu val：class 3
key: zzu val：class4
key: lisi val：class 1
key: zhangsan val：class 2
success delete!
key: lisi val：class 1
key: zhangsan val：class 2
key: wangwu val：class 3
*/