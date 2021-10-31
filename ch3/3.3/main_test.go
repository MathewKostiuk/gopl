package main

import "testing"

func BenchmarkSuperimage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Supersample(2048, 2048)
	}
}
