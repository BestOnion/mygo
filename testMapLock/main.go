package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	totalMap = make(map[int]uint64, 200)
)
var mapLock sync.RWMutex

func jiecheng(n int) {
	res := 1
	for j := n; j > 0; j-- {
		res += j
	}
	mapLock.Lock()
	totalMap[n] = uint64(res)
	mapLock.Unlock()
}

func main() {
	for i := 200; i > 0; i-- {
		go jiecheng(i)
	}
	time.Sleep(time.Second * 20) //fatal error: concurrent map writes
	mapLock.Lock()
	fmt.Println(totalMap)
	mapLock.Unlock()
}
