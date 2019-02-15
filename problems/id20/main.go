package main 

import (
	"fmt"
	"strconv"
)

func main() {
	n := 100
	val := 0

	for i:=n; i>0; i-- {
		if n-1 >0 {
			val = i*n-1
		}
	}

	fmt.Println(val)
}
