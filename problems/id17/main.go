package main

import (
	"fmt"
	"github.com/divan/num2words"
)

func main() {
	total := 0
	for i := 1; i <= 1000; i++ {
		num := num2words.ConvertAnd(i)
		for _, v := range num {
			if v != 32 && v != 45 {
				total++
			}
		}
	}
	fmt.Println(total)
}
