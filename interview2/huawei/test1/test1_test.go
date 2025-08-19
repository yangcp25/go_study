package test1

import "testing"

func TestSubStrLasPos(t *testing.T) {
	type args struct {
		str string
		sub string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 基础功能测试
		{"Single char match", args{"abcdc", "abc"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubStrLasPos(tt.args.str, tt.args.sub); got != tt.want {
				t.Errorf("SubStrLasPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
