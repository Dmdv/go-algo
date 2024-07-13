package longest_common_prefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		source string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{"abcdef", "abc"}, 3},
		{"case 2", args{"abcdef", "def"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
