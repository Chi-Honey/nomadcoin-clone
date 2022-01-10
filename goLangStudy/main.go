package main

import "fmt"

func plus(a, b int, name string) (int, string) {
	return a + b, name
}

func plusAll(a ...int) int {
	total := 0
	for index, item := range a {
		total += item
		total += index
	}
	return total
}
func plusItem(a ...int) int {
	total := 0
	for _, item := range a { // _로 인덱스 무시
		total += item
	}
	return total
}

func main() {
	result, name := plus(2, 2, "nico")
	result2 := plusAll(2, 3, 4, 5, 6, 7, 8, 9, 10)
	result3 := plusItem(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result, name)
	fmt.Println(result2, result3)
}
