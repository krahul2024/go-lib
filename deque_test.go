package main

import (
	"testing"

	"github.com/krahul2024/go-lib/deque"
)

func BenchmarkDequeOperations(b *testing.B) {
	d := deque.NewDeque[int]()
	for i := 0; i <= b.N; i++ {
		d.PushFront(i)
		d.PopFront()
	}

	for i := 0; i <= b.N; i++ {
		d.PushBack(i)
	}
}
