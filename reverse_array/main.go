package main

func main() {
	// Reverse the array
	// [1, 2, 3, 4, 5] => [5, 4, 3, 2, 1]
	arr := []int{1, 2, 3, 4, 5}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
