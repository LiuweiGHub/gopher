package test_test

import (
	"gopher/test"
	"testing"
)

/**
  性能测试
*/
func BenchmarkMakeSliceWithAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test.MakeSliceWithAlloc()
	}
}

func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test.MakeSliceWithoutAlloc()
	}
}
