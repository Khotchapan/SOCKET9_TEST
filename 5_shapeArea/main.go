package main

import "fmt"

func main() {
	result := solution(7000)
	fmt.Println(result)
}

func solution(n int) int {
	return n*n + (n-1)*(n-1)
}
