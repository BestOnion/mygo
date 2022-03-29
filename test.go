package main

import (
	"fmt"
	"math"
)

func main() {
	str := "http://www.baidu.com"
	fmt.Print(str)

}
func (st *Stack) Pop() int {
	v := 0
	for ix := len(st) - 1; ix >= 0; ix-- {
		if v = st[ix]; v != 0 {
			st[ix] = 0
			return v
		}
	}
}
