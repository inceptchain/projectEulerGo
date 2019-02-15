package main 

import (
	"fmt"
)

func main() {
	startWDay := 1
	startYear := 1901
	endYear := 2000 
	tSundays := 0
	wDays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday",}
	for y := startYear; y <= endYear; y++ {

		mDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,}

		if y%4 == 0 && y%100 != 0{
			mDays[1] = 29 
		}

		if y%400 == 0 {
			mDays[1] = 29
		}

		for _, v := range mDays{
			
			for i := 1; i<=v; i++ {

				if i == 1 && wDays[startWDay] == "Sunday" {
					tSundays ++
				}
				if startWDay >= 6{
					startWDay = 0
				}else{
					startWDay ++
				}
			}
		}
		fmt.Println(tSundays)
	}	
}