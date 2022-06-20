package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			result = append(result, arr[i])
		}
	}
	fmt.Println(result)

}
