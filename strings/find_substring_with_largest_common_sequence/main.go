package main

func main() {
	// Given a string s and a string t, find the length of the longest substring of s that contains only characters from t.

	// Example 1:
	// Input: s = "abcde", t = "ace"
	// Output: 3
	// Explanation: The substring "ace" is the longest substring that consists only of characters from t.

	// Example 2:
	// Input: s = "abc", t = "abc"
	// Output: 3
	// Explanation: The substring "abc" is the longest substring that consists only of characters from t.

	// Example 3:
	// Input: s = "a", t = "a"
	// Output: 1
	// Explanation: The substring "a" is the longest substring that consists only of characters from t.

	// Constraints:
	// 1 <= s.length <= 10^5
	// 1 <= t.length <= 10^5
	// s and t consist of English letters.

	// Approach:
	// 1. Create a map to store the frequency of characters in t.
	// 2. Create a variable to store the start index of the substring.
	// 3. Create a variable to store the max length of the substring.
	// 4. Create a variable to store the count of characters in t.
	// 5. Iterate over the string s.
	// 6. If the character is in t, decrement the count of characters in t.
	// 7. If the count of characters in t is less than 0, increment the count of characters in t.
	// 8. If the count of characters in t is 0, calculate the length of the substring.
	// 9. If the length of the substring is greater than the max length of the substring, update the max length of the substring.
	// 10. Return the max length of the substring.

	// Complexity Analysis:
	// Time Complexity: O(n), where n is the length of the string s.
	// Space Complexity: O(1).

}
