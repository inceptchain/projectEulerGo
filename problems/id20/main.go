package main

import (
	"fmt"
)

func main() {
	num := 100
	s := []int{}
	remainder := 0
	if num >= 100 && num < 1000 {
		i2 := (num / 100) % 10
		i1 := (num / 10) % 10
		i0 := (num % 10)

		s = append(s, i2, i1, i0)

		fmt.Println(i2)
		fmt.Println(i1)
		fmt.Println(i0)
		fmt.Println(s)
	}
	for i := 20; i > 0; i-- {
		l := len(s)
		if i > 9 && i < 100 {
			i1 := (i / 10) % 10
			i0 := (i % 10)

			for j := l; j >= 0; j-- {
				n := s[j] * i0
				//n1 := i1*j[l]

				if n > 9 {
					j1 := (n / 10) % 10
					j0 := (n % 10)
					s[j] = j0 + remainder
					remainder = 0

					if j == 0 {
						s = append([]int{j1}, s...)
					} else {
						remainder = j1
					}
				} else {

				}
			}
			fmt.Println(i1)
			fmt.Println(i0)
			fmt.Println(s)

		}
	}
}
