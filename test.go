package main

import (
	"fmt"
	"math"
)

func main() {
	str := "http://www.baidu.com"
	fmt.Print(str)

}

func Unit8FromInt(n int) (uint8, error) {

	if 0 <= n && n < math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is is out of the unit8 ranger", n)
}
