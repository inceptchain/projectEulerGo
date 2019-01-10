package main 

import "fmt"

func main() {
	max := 2000000
	value := 2
	sum := 0
	sp := []int{}
	for value <= max {
		for i:=2; i<value; i++ {
			if value%i == 0 && i != 1 && i != value {
				value++
				goto end
			} 
		}
		sp = append(sp, value)
		fmt.Println(value)
		value++
		end:
	}
	


		for _, v := range sp {
			sum += v
		}

	fmt.Println(sum)
}