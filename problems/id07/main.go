package main 
//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, 
//we can see that the 6th prime is 13.
//What is the 10 001st prime number?

import "fmt"

func main() {
	index := 10001
	count := 0
	num := 2

	for count < index {
		for i:=2; i<num; i++ {
			if num%i == 0 {
				num++
				goto end
			}
		}
		fmt.Println(num)
		count++
		num++
		end:
	}
}