package main

import "fmt"

func main() {
	result := solution([]int{3, 6, -2, -5, 7, 3})
	fmt.Println(result)
}

func solution(inputArray []int) int {
	length := len(inputArray)
	maximum := inputArray[0] * inputArray[1]
	for i := 0; i < length-1; i++ {
		temporary := inputArray[i] * inputArray[i+1]
		if temporary > maximum {
			maximum = temporary
		}
	}
	return maximum
}
