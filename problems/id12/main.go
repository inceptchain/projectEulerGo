package main 

import (
	"fmt"
	"math"
)

func main() {
	terms := []int{}
	max := 0
	createTerms(13000, &terms) //could automate this
		
	for i:=0; i<len(terms); i++ {
		count := []int{}
		sqrt := math.Sqrt(float64(terms[i]))
		for j:=1.0; j<=sqrt; j++ {
			if terms[i]%int(j) == 0 {
				a := terms[i]/int(j)
				if a == int(j) {
					count = append(count, a)
				} else {
					count = append(count, a, int(j))	
				}	
			}
		}
		if len(count) > max {
			max = len(count)
			fmt.Println(count[0])
		}
	}
	fmt.Println(max)
}

func createTerms(value int, slice *[]int) {
	for i:=1; i<=value; i++ {
		num := 0
		num = i*(i+1)/2
		*slice = append(*slice, num)
		//fmt.Println()
	}
}