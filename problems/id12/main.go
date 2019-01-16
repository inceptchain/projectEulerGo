package main 

import (
	"fmt"
)

func main() {
	nat := []int{}
	terms := []int{}
	value := 5
	divisors := 0
	for value > divisors{
		num := 1
		terms = append(terms, num)
		for i, v := range terms {
			for j:=1; j<=len(terms); j++ {
				if v%j == 0 {
					divisors++
					fmt.Println(terms[i])
				}
			}
			num++
			num = terms[i] + num
		}
		
	}
}