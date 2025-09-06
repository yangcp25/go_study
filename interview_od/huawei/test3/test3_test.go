package test3

import (
	"testing"
)

func TestExchangePos(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		// 基础情况
		{"", ""},
		{"A", "A"},
		{"a", "a"},
		{"1", "1"},

		// 已经有序
		{"ABCabc123", "ABCabc123"},

		// 完全逆序
		{"123cbaCBA", "CBAcba123"},

		// 混合
		{"aA1", "Aa1"},
		{"1aA", "Aa1"},
		{"Aa1Bb2Cc3", "ABCabc123"},

		// 两类字符
		{"abc123", "abc123"},
		{"ABC123", "ABC123"},
		{"ABCabc", "ABCabc"},

		// 全部同类
		{"ABCDEF", "ABCDEF"},
		{"abcdef", "abcdef"},
		{"123456", "123456"},

		// 随机混乱
		{"Zy9Xx8Ww7", "ZXWyxw987"},
		{"mN1oO2pP3", "NOPomp213"},

		// 边界
		{"A1a", "Aa1"},
		{"1A1aA1", "AAa111"},
		{"zZ9zZ9", "ZZzz99"},

		// 重复模式 / 长度较大
		{"Aa1Aa1Aa1Aa1", "AAAAaaaa1111"},
		{"987ZYXcba654", "ZYXcba987654"},
	}

	for _, c := range cases {
		runes := []rune(c.input)
		exchangePos(runes)
		got := string(runes)
		if got != c.expected {
			t.Errorf("exchangePos(%q) = %q; expected %q", c.input, got, c.expected)
		}
	}
}
