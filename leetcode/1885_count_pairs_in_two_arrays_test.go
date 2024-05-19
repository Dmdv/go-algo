package leetcode

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test", args{[]int{2, 1, 3}, []int{7, 3, 2}}, 0},
		{"test", args{[]int{2, 1, 3}, []int{3, 7, 2}}, 0},
		{"test", args{[]int{1, 10, 6, 2}, []int{1, 4, 1, 5}}, 5},
		{"test", args{[]int{2, 1, 2, 1}, []int{1, 2, 1, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
