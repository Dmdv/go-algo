package main

import "fmt"

func main() {
	// Rotate the array to right by k steps
	// [1, 2, 3, 4, 5] => [4, 5, 1, 2, 3]
	arr := []int{1, 2, 3, 4, 5}
	k := 6

	n := len(arr)
	k = k % n

	reverse(arr, 0, n-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, n-1)

	fmt.Printf("%v\n", arr)
}

func reverse(arr []int, i int, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
