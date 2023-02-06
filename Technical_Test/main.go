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

	if err := ValidateData(dateInt); err != nil {
		fmt.Println(err)
		return
	}

	for year := 1900; year < dateInt[2]; year++ {
		if IsValidLeapYear(year) {
			totalDay += 366
		} else {
			totalDay += 365
		}
	}

	if IsValidLeapYear(dateInt[2]) {
		monthCount[1] = 29
	}

	for month := 0; month < dateInt[1]-1; month++ {
		totalDay += monthCount[month]
	}

	totalDay += dateInt[0] - 1
	fmt.Println("Today is", dayNames[totalDay%7])
}

func ValidateData(data []int) error {

	if IsThirtyFirstDay(data[0]) && IsDurationThirtyFirstDay(data[0]) {
		return fmt.Errorf("Error Not within the specified date 1")
	}
	if IsThirtiethDay(data[0]) && IsDurationThirtiethDay(data[0]) {
		return fmt.Errorf("Error Not within the specified date 2")
	}
	if !IsValidLeapYear(data[2]) && IsDurationLeapYear(data[0], data[1]) {
		return fmt.Errorf("Error Not within the specified date 3")
	}
	if IsValidLeapYear(data[2]) && IsNotDurationLeapYear(data[0], data[1]) {
		return fmt.Errorf("Error Not within the specified date 4")
	}
	if IsDurationMonth(data[1]) {
		return fmt.Errorf("Error Not within the specified month")
	}
	if IsDurationYear(data[2]) {
		return fmt.Errorf("Error Not within the specified year")
	}

	return nil
}

func IsValidLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}

func IsThirtiethDay(day int) bool {
	thirtieth := []int{1, 3, 5, 7, 8, 10, 12}
	length := len(thirtieth)
	for i := 0; i < length; i++ {
		if thirtieth[i] == day {
			return true
		}
	}
	return false
}

func IsThirtyFirstDay(day int) bool {
	thirtyFirst := []int{4, 6, 9, 11}
	length := len(thirtyFirst)
	for i := 0; i < length; i++ {
		if thirtyFirst[i] == day {
			return true
		}
	}
	return false
}

func IsDurationThirtiethDay(day int) bool {
	if day < 1 || day > 30 {
		return true
	}
	return false
}

func IsDurationThirtyFirstDay(day int) bool {
	if day < 1 || day > 31 {
		return true
	}
	return false
}

func IsDurationLeapYear(day, month int) bool {
	if month == 2 && day < 1 || day > 28 {
		return true
	}
	return false
}
func IsNotDurationLeapYear(day, month int) bool {
	if month == 2 && day < 1 || day > 29 {
		return true
	}
	return false
}
func IsDurationMonth(month int) bool {
	if month < 1 || month > 12 {
		return true
	}
	return false
}
func IsDurationYear(year int) bool {
	if year < 1900 {
		return true
	}
	return false
}
