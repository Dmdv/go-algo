package longest_common_prefix

func longestCommonPrefix(source, target string) int {
	// Given two strings source and target, return the length of the longest common prefix of source and target.
	// If there is no common prefix, return 0.
	// A prefix is a substring of a string that occurs at the beginning of the string.
	// Example 1:
	// Input: source = "abcdef", target = "abc"
	// Output: 3
	// Explanation: The longest common prefix is "abc".
	// Example 2:
	// Input: source = "abcdef", target = "def"
	// Output: 0
	// Explanation: There is no common prefix.
	// Constraints:
	// 1 <= source.length, target.length <= 10^5
	// source and target consist of lowercase English letters.
	// Approach:
	// 1. Create a variable to store the length of the longest common prefix.
	// 2. Iterate over the strings source and target.
	// 3. If the characters are equal, increment the length of the longest common prefix.
	// 4. Return the length of the longest common prefix.
	// Complexity Analysis:
	// Time Complexity: O(n), where n is the length of the string source.
	// Space Complexity: O(1).
	var result int
	for i := 0; i < len(source) && i < len(target); i++ {
		if source[i] == target[i] {
			result++
		} else {
			break
		}
	}
	return result
}
