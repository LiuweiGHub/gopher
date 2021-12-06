package test_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	println("TestMain setup")

	retCode := m.Run() //执行测试  包括性能测试、单元测试、示例测试

	println("TestMain teardown")

	os.Exit(retCode)
}
