package main

import "fmt"

func main() {
	result := solution(1905)
	fmt.Println(result)
}

func solution(year int) int {
	var result int
	if year <= 0 {
		result = 0
	} else if year <= 100 {
		result = 1
	} else if year%100 == 0 {
		result = year / 100
	} else {
		result = (year / 100) + 1
	}

	return result
}
