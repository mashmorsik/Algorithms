package main

import "testing"

func BenchmarkFindOutlier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindOutlier([]int{100, 3, 208, 666, 700, 60, 14})
	}
}

func BenchmarkFindOutlierFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindOutlierFast([]int{100, 3, 208, 666, 700, 60, 14})
	}
}

func BenchmarkFindOutlierFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindOutlierFaster([]int{100, 3, 208, 666, 700, 60, 14})
	}
}
