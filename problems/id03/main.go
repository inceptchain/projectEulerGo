package main 

import "fmt"

//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?

func main() {
	v := 600851475143
	pf := 0

	for i:=3; i<v; i++ {
		if v%i == 0 {

			for j:=2; j<i; j++ {
				if i%j == 0{
					goto end
				}
			}

		pf = i
		fmt.Println(pf)		
	end:		
		}
	}
}