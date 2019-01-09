package main

//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.
import (
	"fmt"
)

var mult = []int{}

func main() {
	t:=0
	findMultiples(1000)
	for _, m := range mult {
		t += m	
	}

	fmt.Println(t)
}

func findMultiples(to int) {
	m1 := 3
	m2 := 5
	for i := 1; i < to; i++ {
		if i%m1 == 0 || i%m2 == 0 {
			mult = append(mult, i)
		}
	}
}
