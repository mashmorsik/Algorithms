package main

import "testing"

func BenchmarkIsPalindromeTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(3003)
	}
}

func BenchmarkIsPalindromeRuneTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindromeRune(3003)
	}
}
