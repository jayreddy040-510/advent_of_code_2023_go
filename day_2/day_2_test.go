package day2

import "testing"

func BenchmarkPuzzleOne(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		PuzzleOne()
	}
}
