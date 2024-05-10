package main

import "testing"

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{2, 7, 11, 15}, 9)
	}
}

func BenchmarkTwoSumFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumFaster([]int{2, 7, 11, 15}, 9)
	}
}
