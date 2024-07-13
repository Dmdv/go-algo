package matrix

import "testing"

func Test_rotate90Clockwise(t *testing.T) {
	testsClock := []struct {
		name   string
		matrix [][]int
		result [][]int
	}{
		{"clockwise",
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		}}

	for _, tt := range testsClock {
		t.Run(tt.name, func(t *testing.T) {
			rotate90Clockwise(tt.matrix)
			if !compareMatrix(tt.matrix, tt.result) {
				t.Errorf("got %v, want %v", tt.matrix, tt.result)
			} else {
				t.Logf("got %v, want %v", tt.matrix, tt.result)
			}
		})
	}
}

func Test_rotate90AntiClockwise(t *testing.T) {
	testsClock := []struct {
		name   string
		matrix [][]int
		result [][]int
	}{
		{"anti-clockwise",
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}},
		}}

	for _, tt := range testsClock {
		t.Run(tt.name, func(t *testing.T) {
			rotate90AntiClockWise(tt.matrix)
			if !compareMatrix(tt.matrix, tt.result) {
				t.Errorf("got %v, want %v", tt.matrix, tt.result)
			} else {
				t.Logf("got %v, want %v", tt.matrix, tt.result)
			}
		})
	}
}

func compareMatrix(a [][]int, b [][]int) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
