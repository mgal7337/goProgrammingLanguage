package main

import "testing"

func BenchmarkPrintArgsUsingSeparator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintArgsUsingSeparator()
		//605384	      1952 ns/op	     400 B/op	       5 allocs/op
	}
}

func BenchmarkPrintArgsUsingJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintArgsUsingJoin()
		//703628	      1781 ns/op	     160 B/op	       2 allocs/op
	}
}
