package main 

import (
	"fmt"
	"math"
)

func main() {
	run := true
	for run {
		for a := 1; a < 1000; a++ {
			for b := 1; b < 1000; b++ {
				c := math.Sqrt(float64(a*a) + float64(b*b))
				if float64(a)+float64(b)+c == 1000 {
					fmt.Printf(" a = %v, b = %v, c =%v\n", a,b,c)
					fmt.Printf("product: %v\n", float64(a)*float64(b)*c)
					run = false	
				}	
			}
		}
	}	
}