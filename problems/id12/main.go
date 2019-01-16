package main 

import (
	"fmt"
)

func main() {
	terms := []int{}
	length := 0
	divisors := 4

	createTerms(25, &terms)

	for i:= 1; length <= divisors; i++ {
		
		temp := 0
		//fmt.Println(result)

		for _, v := range terms {
			temp = getFactors(v)
		}

		if temp > length {
			length = temp
			fmt.Printf("Length: %v\n", length)
			fmt.Printf("Value: %v\n", terms[len(terms)-1])
		}

	}
	//fmt.Println(length)
}

func getFactors(value int) int{
	factors := []int{}
	for i := 1; i <= value; i++ {
		if value%i == 0 {
			factors = append(factors, i)
			fmt.Println(i)
		}
		//fmt.Println(factors)	
	}
	return len(factors)
}

func createTerms(value int, slice *[]int) {
	num := 0
	for i:=1; i<=value; i++ {
		num += i
		*slice = append(*slice, num)
	}
}