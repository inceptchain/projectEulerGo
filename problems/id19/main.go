package main 

import (
	"fmt"
)

func main() {
	startWDay := 1
	startYear := 1901
	endYear := 2000 
	tDays := 0
	mDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	wDays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for y := startYear; y <= endYear; y++ {
		if y % 4 == 0 && y%100 != 0{
			mDays[1] = 29
			startWDay++ 
		}

		if y%400 == 0 {
			mDays[1] = 29
			startWDay++
		}

		for _, v := range mdays{
			tDays += v
		}

		for 

	}	
}