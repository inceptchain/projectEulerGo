package main

import (
	"fmt"
)

func main() {
	s := []int{}
	exponent := 15
	total := 2
	s = append(s, total)
	for i := 0; i < exponent; i++ {
		for j, v := range s {
			s[j] = v * 2

			if s[j] > 9 {
				s[j] = s[j] % 10
				s = append(s, 1)
				break
			}
		}
	}

	fmt.Println(s)
}
