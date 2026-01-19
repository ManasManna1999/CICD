package main

import "fmt"

func Add(a, b int) int {
	return a + b
}
//given a comments to check the pipelines
func main() {
	fmt.Println("CI/CD Demo: 1 + 2 =", Add(1, 2))
}
