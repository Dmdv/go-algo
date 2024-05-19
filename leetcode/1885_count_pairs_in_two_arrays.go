package leetcode

import "sort"

// Given two integer arrays nums1 and nums2 of length n, count the pairs of indices (i, j) such that i < j and nums1[i] + nums1[j] > nums2[i] + nums2[j].
// Hint: We can write it as nums1[i] - nums2[i] > nums2[j] - nums1[j] instead of nums1[i] + nums1[j] > nums2[i] + nums2[j].

// nums1[i] + nums1[j] > nums2[i] + nums2[j] is the equal to
// (nums1[i] - nums2[i]) + (nums1[j] - nums2[j]) > 0
// So, if we define nums = [nums1[0]-nums2[0], nums1[1]-nums2[1], ...],
// This problem can be rewritten as:
// How many (i, j) pairs with i < j that nums[i] + nums[j] > 0

// It's a classic problem like:
// Given a value x, count the number of y in the nums that satisfies (y + x > 0)
//
// So we can sort the nums, and pick element one by one. If diff(right) > diff(left), means right is larger than the left,
// which means we can use right - left pairs.
// Then use binary search to find out how many values which are larger than the negative of the picked element.
// Finally, remember to skip the counted pairs. (That means we just want to find the index j, which is larger than i.)

func countPairs(nums1 []int, nums2 []int) int64 {
	diff := make([]int, 0, len(nums1))

	for i := 0; i < len(nums1); i++ {
		diff = append(diff, nums1[i]-nums2[i])
	}

	sort.Ints(diff)

	result := 0
	left, right := 0, len(diff)-1

	for left < right {
		i := diff[right] + diff[left]
		if i > 0 {
			right--
		} else {
			result += len(diff) - 1 - right
			left++
		}
	}

	result += (len(diff) - right) * (len(diff) - right - 1) / 2

	return int64(result)
}
