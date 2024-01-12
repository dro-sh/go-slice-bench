package main

import (
	"fmt"
	"testing"
)

var table = []int{
	5,
	10,
	50,
	100,
	500,
	1000,
	10000,
}

func BenchmarkInsertAtSlice(b *testing.B) {
	for _, count := range table {
		b.Run(fmt.Sprintf("input_size_%d", count), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = appendSlice(count)
			}
		})
	}
}
