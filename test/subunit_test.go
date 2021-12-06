package test_test

import (
	"gopher/test"
	"testing"
)

/**
子测试
*/
func sub1(t *testing.T) {
	var a = 1
	var b = 2
	var c = 3

	actual := test.Add(a, b)

	if c != actual {
		t.Errorf("%d, %d", c, actual)
	}
}

func sub2(t *testing.T) {
	var a = 1
	var b = 2
	var c = 3

	actual := test.Add(a, b)

	if c != actual {
		t.Errorf("%d, %d", c, actual)
	}
}

func sub3(t *testing.T) {
	var a = 1
	var b = 2
	var c = 3

	actual := test.Add(a, b)

	if c != actual {
		t.Errorf("%d, %d", c, actual)
	}
}

// go test subunit_test.go -v   -v参数代表执行测试
// go test subunit_test.go -v -run /A   -run /A 代表只执行 匹配（包含/A）的用例
func TestSub(t *testing.T) {
	//setup code
	t.Run("A=1", sub1)
	t.Run("A=2", sub2)
	t.Run("B=1", sub3)
}
