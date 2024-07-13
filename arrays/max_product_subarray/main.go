package main

func main() {
	// Find the contiguous subarray within an array (containing at least one number) which has the largest product.
	// [2, 3, -2, 4] => 6
	// [-2, 0, -1] => 0
	arr := []int{2, 3, -2, 4}

	maxProduct := arr[0]
	minProduct := arr[0]
	result := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < 0 {
			maxProduct, minProduct = minProduct, maxProduct
		}

		maxProduct = max(arr[i], maxProduct*arr[i])
		minProduct = min(arr[i], minProduct*arr[i])

		result = max(result, maxProduct)
	}

	println(result)
}
