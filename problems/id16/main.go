package main

import (
	"fmt"
)

func main() {
	s := []int{}
	exponent := 10

	total := 2
	s = append(s, total)
	for i := 0; i < exponent; i++ {
		for j, v := range s {
			s[j] = v * 2

			if s[j] > 9 && j == 0 {
				s[j] = s[j] % 10
				s = append([]int{1}, s...)
				break
			}
			if s[j] > 9 && j != 0 {
				s[j] = s[j] % 10
				s[j-1] = s[j-1] + 1
			}
		}
	}

	fmt.Println(s)
}
