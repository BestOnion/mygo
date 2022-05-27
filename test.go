package main

import (
	"fmt"
	"log"
	"unicode/utf8"
)

func main() {

	//性能调优代码片段

	//1.字符串处理

	str := "hello哈"
	bytestring := []byte(str)
	bytestring[0] = 'c'

	str = string(bytestring)
	fmt.Println(str)

	//2.获取字符串的子串

	sustr := str[:3]
	fmt.Println(sustr)

	for i, _ := range sustr {
		fmt.Println(i)
		fmt.Println(sustr[i])
	}
	fmt.Println(len(sustr))                  //字符串的字节数
	fmt.Println(utf8.RuneCountInString(str)) //字符串的字符数

	//初始化数组
	arr := new([2]string)
	arr[0] = "23"
	fmt.Println(arr)
	//初始化切片
	sl := make([]string, 5, 10)
	sl[0] = "哈哈哈"
	sl[1] = "132"
	sl[2] = "132"
	sl[3] = "132"
	sl[4] = "132"
	fmt.Println(sl)
	//初始化切片
	maps := make(map[string]int)
	maps["map"] = 1
	maps["map2"] = 1
	fmt.Println(maps)
	//检查map key是否存在
	if data, exits := maps["map"]; exits {
		fmt.Println("存在数据")
		fmt.Println(data)
	}
	//删除map 中的key
	delete(maps, "map2")

	//2.结构体
	//初始化结构体

	classifier(true, 0.2, "hah")
}

func classifier(items ...interface{}) {
	for _, x := range items {
		fmt.Print(x)
		switch x.(type) {
		case bool:
			fmt.Println("is bool")
		case float32:
			fmt.Println("is float 32")
		case string:
			fmt.Println("is string")
		default:
			fmt.Println("is default")
		}
	}
	protot(test)

	fmt.Print("继续啦1")
}
func test() {
	fmt.Println("qwer")
	panic("这里不能继续了")
	fmt.Print("继续啦")
}

// 如何用内建函数recover()终止panic过程
func protot(g func()) {
	defer func() {
		log.Println("done")
		if x := recover(); x != nil {
			log.Printf("runtime panic %v", x)
		}
	}()
	fmt.Println("1s")
	log.Println("start")
	g()
	fmt.Print("hahah")
}
