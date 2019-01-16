package main 

import (
	"fmt"
)

func main() {
	terms := []int{}

	result := createTerms(10, &terms)
	fmt.Println(result)

	for _, v := range terms {
		fmt.Println(v)
	}


}

func getFactors(value int) {
	
	for i := 1; i <= value; i++ {
		
	}
}

func createTerms(value int, slice *[]int) int{
	num := 0
	for i:=1; i<=value; i++ {
		num += i
		*slice = append(*slice, num)
	}
	return num
}