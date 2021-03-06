package main 
//Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
import (
	"fmt"
)

func main() {
	
	t := 0
	a := 1
	b := 2
	num := 0


	for b <= 4000000 {
		
		num = b
		b = a+num
		a = num

		if a%2 == 0 {
			t += a
		}
		fmt.Println(b)
	}

	fmt.Println(t)

}