package singleton

import (
	"reflect"
	"testing"
)

func TestGetInstance(t *testing.T) {
	var tests []struct {
		name string
		want *Singleton
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestGetLazyInstance(t *testing.T) {
	var tests []struct {
		name string
		want *Singleton
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLazyInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLazyInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
