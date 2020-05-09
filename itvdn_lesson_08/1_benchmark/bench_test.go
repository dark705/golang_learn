package bench

import "testing"

// go test -bench=. -benchmem
func BenchmarkFib11(b *testing.B) {
	b.SetBytes(int64(fib(11)))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		fib(11)
	}
}

func BenchmarkFib32(b *testing.B) {
	b.SetBytes(int64(fib(32)))
	b.ResetTimer() //reset timer we need real time work
	for n := 0; n < b.N; n++ {
		fib(32)
	}
}
