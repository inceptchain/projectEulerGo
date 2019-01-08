package main 

//The sum of the squares of the first ten natural numbers is,

//1^2 + 2^2 + ... + 10^2 = 385
//The square of the sum of the first ten natural numbers is,

//(1 + 2 + ... + 10)^2 = 552 = 3025
//Hence the difference between the sum of the squares of the first 
//ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//Find the difference between the sum of the squares of the first one hundred 
//natural numbers and the square of the sum.

import "fmt"

func main() {
	nat := 100
	sum1 := 0
	sum2 := 0
	dif := 0

	for i := 1; i <= nat; i++ {
		sum1 += i*i
		sum2 += i

		if i == nat {
			sum2*=sum2
		}
	}

	dif = sum2 - sum1
	fmt.Println(dif)
}