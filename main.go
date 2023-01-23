package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "23,01,2023"
	dayNames := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	text := strings.Split(input, ",")
	dateInt := []int{}
	for index := range text {
		temp, err := strconv.Atoi(text[index])
		if err != nil {
			fmt.Println("Error during conversion.")
			return 
		}
		dateInt = append(dateInt, temp)
	}
	totalDay := 0
	for year := 1900; year < dateInt[2]; year++ {
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			totalDay += 366
		} else {
			totalDay += 365
		}
	}
	monthCount := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if dateInt[2]%4 == 0 && dateInt[2]%100 != 0 || dateInt[2]%400 == 0 {
		monthCount[1] = 29
	} else {
		monthCount[1] = 28
	}
	for month := 0; month < dateInt[1]-1; month++ {
		totalDay += monthCount[month]
	}
	totalDay += dateInt[0] - 1
	fmt.Println("Today is", dayNames[totalDay%7])
}
