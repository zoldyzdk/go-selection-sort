package main

import "fmt"

func lowestValue(teste []int) int {
	arr := teste
	lowestIndex := 0
	lowest := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < lowest {
			lowest = arr[i]
			lowestIndex = i
		}
	}
	return lowestIndex
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	var newArr []int
	arr := []int{3, 2, 7, 5, 1}
	for range arr {
		minor := lowestValue(arr)
		newArr = append(newArr, arr[minor])
		arr = remove(arr, minor)
	}
	fmt.Println(newArr)
}
