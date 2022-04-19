package main

<<<<<<< HEAD
import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// 版本A:
	for ix := range values { // ix是索引值
		func() {
			fmt.Print(ix, " ")
		}() // 调用闭包打印每个索引值
=======
func main() {
	a := "hello"
	var b *string
	var c *string
	for k, v := range a {
		if k < 2 {
			*b = v
		}
		if k >= 2 {
			&c += v
		}
>>>>>>> 60df1f4f46255541ed2e92dc8769ab2e51f67e16
	}
	fmt.Println()
	// 版本B: 和A版本类似，但是通过调用闭包作为一个协程
	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()
	time.Sleep(5e9)
	// 版本C: 正确的处理方式
	for ix := range values {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	fmt.Println()
	fmt.Println()
	time.Sleep(5e9)
	// 版本D: 输出值:
	for ix := range values {
		val := values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	time.Sleep(1e9)
}
