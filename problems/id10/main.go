package main 

import (
	"fmt"
	"math"
)

func main() {
	max := 2000000
	value := 2
	sum := 0
	sp := []int{}
	sqrt := 0
	for value <= max {
		sqrt = int(math.Sqrt(float64(value)))
		for i:=1; i<value; i++ {
			if value%i == 0 && i != 1{
				break
			}

			if i > sqrt || value == 2{
				sp = append(sp, value)
				fmt.Println(value)
				break
			}
		}
		value++
	}

	for _, v := range sp {
		sum += v
	}

	fmt.Println(sum)
}