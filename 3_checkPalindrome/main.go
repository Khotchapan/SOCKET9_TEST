package main

import "fmt"

func main() {
	result := solution("aabaa")
	fmt.Println(result)
}

func solution(inputString string) bool {
	length := len(inputString)
	for i := 0; i < length/2; i++ {
		if inputString[i] != inputString[length-1-i] {
			return false
		}
	}
	return true
}
