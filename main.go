// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	days := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	input := "11 07 1907"
	text := strings.Split(input, " ")
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
		//fmt.Println(year)
		if year%4 == 0 {
			totalDay += 366
			//fmt.Println("The year is a leap year!")
		} else {
			totalDay += 365
			//fmt.Println("The year isn't a leap year!")
		}
	}
	fmt.Println(totalDay, "days")
	fmt.Println("Today is Monday", days[totalDay%7])
}
