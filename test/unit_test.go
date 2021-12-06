package test_test

import (
	"gopher/test"
	"testing"
)

/**
单元测试
*/
func TestAdd(t *testing.T) {
	var a = 1
	var b = 2
	var c = 3

	actual := test.Add(a, b)

	if c != actual {
		t.Errorf("%d, %d", c, actual)
	}
}
