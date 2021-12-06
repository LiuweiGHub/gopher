package test_test

import "gopher/test"

/**
示例测试
*/
func ExampleSayHello() {
	test.SayHello()
	// output: Hello world
}

func ExampleSayGoodBye() {
	test.SayGoodBye()
	//output:
	//Good
	//Bye
}

func ExamplePrintNames() {
	test.PrintNames()
	//unordered output:
	//bob
	//emm
	//susan
	//sari
}
