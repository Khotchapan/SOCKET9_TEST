package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//input := "31/01/2023"
	var input string
	fmt.Scanf("%s", &input)
	text := strings.Split(input, "/")
	dayNames := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	monthCount := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	dateInt := []int{}
	totalDay := 0

	for index := range text {
		temp, err := strconv.Atoi(text[index])
		if err != nil {
			fmt.Println("Error during conversion.")
			return
		}
		dateInt = append(dateInt, temp)
	}
	if err := Validate(dateInt); err != nil {
		fmt.Println(err)
		return
	}

	for year := 1900; year < dateInt[2]; year++ {
		if IsLeapYear(year) {
			totalDay += 366
		} else {
			totalDay += 365
		}
	}

	if IsLeapYear(dateInt[2]) {
		monthCount[1] = 29
	}

	for month := 0; month < dateInt[1]-1; month++ {
		totalDay += monthCount[month]
	}

	totalDay += dateInt[0] - 1
	fmt.Println("Today is", dayNames[totalDay%7])
}

func Validate(data []int) error {

	if data[0] < 0 || data[0] > 31 {
		return fmt.Errorf("Error Not within the specified date")
	}
	if !IsLeapYear(data[2]) && data[1] == 2 && data[0] > 28 {
		return fmt.Errorf("Error Not within the specified date")
	}
	if IsLeapYear(data[2]) && data[1] == 2 && data[0] > 29 {
		return fmt.Errorf("Error Not within the specified date")
	}

	if data[1] < 0 || data[1] > 12 {
		return fmt.Errorf("Error Not within the specified month")
	}

	if data[2] < 1900 {
		return fmt.Errorf("Error Not within the specified year")
	}

	return nil
}

func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
