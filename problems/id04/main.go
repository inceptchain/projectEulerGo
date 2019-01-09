package main 

import "fmt"
//A palindromic number reads the same both ways. 
//The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//Find the largest palindrome made from the product of two 3-digit numbers.

func main() {
	pal := 0
	r := 0

	for a := 100; a < 1000; a++ {
		for b := 100; b < 1000; b++ {
			r = a*b
			r0 := (r/100000)%10
			r1 := (r/10000)%10
			r2 := (r/1000)%10
			r3 := (r/100)%10
			r4 := (r/10)%10
			r5 := r%10

			if r0 == 0 {
				if r1 == r5 && r2 == r4 {
					pal = updatePalAndPrint(r, pal)	
				} 
			} else if r0 == r5 && r1 == r4 && r2 == r3 {
					pal = updatePalAndPrint(r, pal)
			}
		}
	}	
}

func updatePalAndPrint(r int, pal int) int{
	if r > pal {
		fmt.Println(r)
		return r
	} else {
		return pal
	}
}