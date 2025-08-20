package test1

import "testing"

func TestSubStrLasPos(t *testing.T) {
	tests := []struct {
		name string
		L    string
		S    string
		want int
	}{
		{"S is empty", "abc", "", -1},
		{"L is empty", "", "abc", -1},
		{"S equals L", "abc", "abc", 2},
		{"S longer than L", "abc", "abcd", -1},
		{"Single char match", "a", "a", 0},
		{"Single char no match", "a", "b", -1},
		{"Multiple same chars", "aa", "aa", 1},
		{"Multiple chars missing one", "ab", "aa", -1},
		{"Valid subsequence", "abcde", "ace", 4},
		{"Invalid subsequence", "abcde", "aec", -1},
		{"Another valid case", "abcdc", "abc", 4},
		{"Another invalid case", "aa", "ab", -1},
		{"Large L no match", "abcdefghijklmnopqrstuvwxyz", "zz", -1},
		{"Large L valid", "abcdefghijklmnopqrstuvwxyz", "az", 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubStrLasPos(tt.L, tt.S); got != tt.want {
				t.Errorf("SubStrLasPos(%q, %q) = %v, want %v", tt.L, tt.S, got, tt.want)
			}
		})
	}
}

func BenchmarkSubStrLasPos(b *testing.B) {
	// 创建大规模测试数据
	L := make([]byte, 500000)
	for i := range L {
		L[i] = 'a'
	}
	// 在末尾添加一些不同的字符
	L[499999] = 'b'
	L[499998] = 'c'

	S := "abc"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SubStrLasPos(string(L), S)
	}
}
