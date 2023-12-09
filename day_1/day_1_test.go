package day1

import "testing"

func BenchmarkOldPuzzleTwo(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		OldPuzzleTwo()
	}
}

func BenchmarkPuzzleTwo(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		PuzzleTwo()
	}
}

func BenchmarkPuzzleTwoReverseString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		PuzzleTwoReverseString()
	}
}

func BenchmarkPuzzleTwoReverseBytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		PuzzleTwoReverseBytes()
	}
}
