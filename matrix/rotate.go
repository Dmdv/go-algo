package matrix

func rotate90Clockwise(matrix [][]int) {
	transpose(matrix)

	n := len(matrix)
	// Reverse the elements of each row of the matrix.
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

// Rotate the matrix by 90 degrees in an anti-clockwise direction.
// Approach:
// 1. Find the transpose of the matrix.
// 2. Reverse the elements of each row of the matrix.
// Time Complexity: O(n^2), where n is the length of the matrix.
// Space Complexity: O(1).
func rotate90AntiClockWise(matrix [][]int) {
	transpose(matrix)

	n := len(matrix)
	// Reverse the elements of each row of the matrix.
	for j := 0; j < n; j++ {
		for i := 0; i < n/2; i++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}
}

func transpose(matrix [][]int) {
	n := len(matrix)
	// Find the transpose of the matrix.
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func rotate180(matrix [][]int) {
	transpose(matrix)

	n := len(matrix)
	// Reverse the elements of each row of the matrix.
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}
}
