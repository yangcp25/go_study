package main

import "testing"

func Test_slice(t *testing.T) {
	s := make([]int, 10)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_slice2(t *testing.T) {
	s := make([]int, 0, 10)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_slice3(t *testing.T) {
	s := make([]int, 10, 11)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}
