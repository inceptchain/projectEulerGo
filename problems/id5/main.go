package main 

import "fmt"

//2520 is the smallest number that can be divided by each of the numbers 
//from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all of the numbers 
//from 1 to 20?

func main() {
	r := 1
	max := 20
	run := true

	for run {
		for i:=1; i<=max; i++ {
			
			if r%i != 0 {
				r++
				break
			}

			if i == max {
				run = false
			}
		}
	}

	
	

	fmt.Println(r)
}