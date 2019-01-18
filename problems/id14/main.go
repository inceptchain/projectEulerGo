package main 

import (
	"fmt"
)

func main() {
	maxLen := 0
	sp := 0
	
	
	for j:=13; j<1000000; j++ {
		start := j
		seq := []int{}
		seq = append(seq, start)
		for i:=0; i<len(seq); i++ {

			if seq[i]%2 == 0 {
				num := seq[i]/2
				seq = append(seq, num)

			} else {
				num := (3*seq[i]) + 1
				seq = append(seq, num)
			}

			//fmt.Println(seq[i])

			if seq[i] == 1 {
				seq = append(seq, seq[i])
				if maxLen < len(seq) {
					maxLen = len(seq)
					sp = start
					fmt.Println(maxLen)
					fmt.Println(sp)
				}
				break
			} 
		}
	}
}