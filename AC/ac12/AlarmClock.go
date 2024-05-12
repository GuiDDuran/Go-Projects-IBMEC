package main

import (
	"fmt"
)

func main(){
	var h1, m1, h2, m2 int
	for {
		fmt.Scan(&h1, &m1, &h2, &m2)

		if h1 == 0 && m1 == 0 && h2 == 0 && m2 == 0 {
			break
		}

		currentTime := h1 * 60 + m1
		alarmTime := h2 * 60 + m2
		var sleepTime int

		if currentTime < alarmTime {
			sleepTime = alarmTime - currentTime
		} else {
			sleepTime = 1440 - currentTime + alarmTime
		}

		fmt.Println(sleepTime)
	}
}