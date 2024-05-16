package main

import "fmt"

func main() {
	// Find max sum of subarray

	// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

	// Use Kadane's algorithm

	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := arr[0]
	sum := arr[0]

	for i := 1; i < len(arr); i++ {
		sum = max(arr[i], sum+arr[i])
		maxSum = max(maxSum, sum)
	}

	fmt.Printf("Max sum of subarray: %d\n", maxSum)
	sum = maxSubArray(arr)
	fmt.Printf("Max sum of subarray: %d\n", maxSum)
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxCurrent, maxGlobal := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > maxCurrent+nums[i] {
			maxCurrent = nums[i]
		} else {
			maxCurrent += nums[i]
		}

		if maxCurrent > maxGlobal {
			maxGlobal = maxCurrent
		}
	}

	return maxGlobal
}
