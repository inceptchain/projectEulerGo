package main

import (
	"fmt"
)

func main() {
	s := []int{}
	exponent := 1000
	carry := 0
	total := 1
	value := 0
	s = append(s, total)
	for i := 0; i < exponent; i++ {
		for j := len(s) - 1; j >= 0; j-- {
			s[j] = s[j] * 2
			s[j] += carry
			carry = 0

			if s[j] > 9 && j == 0 {
				s[j] = s[j] % 10
				s = append([]int{1}, s...)
				//break
			}
			if s[j] > 9 && j != 0 {
				s[j] = s[j] % 10
				carry = 1
			}
		}
	}

	for _, v := range s {
		value += v
	}
	fmt.Println(s)
	fmt.Println(value)
}
