package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "01,01,1900"
	dayNames := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	text := strings.Split(input, ",")
	//fmt.Println(text)
	dateInt := []int{}
	for index := range text {
		//fmt.Println(index, element)
		temp, err := strconv.Atoi(text[index])
		if err != nil {
			fmt.Println("Error during conversion.")
			return
		}
		dateInt = append(dateInt, temp)
	}
	fmt.Println(dateInt)
	totalDay := 0
	for year := 1900; year < dateInt[2]; year++ {
		if year%4 == 0 {
			totalDay += 366
			fmt.Println("The year is a leap year 366!")
		} else {
			totalDay += 365
			fmt.Println("The year isn't a leap year 365!")
		}
	}
	fmt.Println("totalDay:", totalDay)
	fmt.Println("dateInt[2]:", dateInt[2])
	monthCount := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if dateInt[2]%4 == 0 {
		monthCount[1] = 29
	} else {
		monthCount[1] = 28
	}
	fmt.Println("monthCount:", monthCount)
	fmt.Println("dateInt[1]:", dateInt[1])
	for month := 0; month < dateInt[1]-1; month++ {
		totalDay += monthCount[month]
	}
	totalDay += dateInt[0] - 1
	fmt.Println("totalDay:", totalDay)
	fmt.Println("Today is", dayNames[totalDay%7])
}
